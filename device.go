package onvif

import (
	"code.byted.org/videoarch/go-onvif/onvif"
	"code.byted.org/videoarch/go-onvif/onvif/common"
	"code.byted.org/videoarch/go-onvif/onvif/v10/device"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"time"

	"code.byted.org/videoarch/go-onvif/gosoap"
	"code.byted.org/videoarch/go-onvif/networking"
	wsdiscovery "code.byted.org/videoarch/go-onvif/ws-discovery"
	"github.com/beevik/etree"
)

//Xlmns XML Scheam
var Xlmns = map[string]string{
	"onvif":   "http://www.onvif.org/ver10/schema",
	"tds":     "http://www.onvif.org/ver10/device/wsdl",
	"trt":     "http://www.onvif.org/ver10/media/wsdl",
	"tev":     "http://www.onvif.org/ver10/events/wsdl",
	"tptz":    "http://www.onvif.org/ver20/ptz/wsdl",
	"timg":    "http://www.onvif.org/ver20/imaging/wsdl",
	"tan":     "http://www.onvif.org/ver20/analytics/wsdl",
	"xmime":   "http://www.w3.org/2005/05/xmlmime",
	"wsnt":    "http://docs.oasis-open.org/wsn/b-2",
	"xop":     "http://www.w3.org/2004/08/xop/include",
	"wsa":     "http://www.w3.org/2005/08/addressing",
	"wstop":   "http://docs.oasis-open.org/wsn/t-1",
	"wsntw":   "http://docs.oasis-open.org/wsn/bw-2",
	"wsrf-rw": "http://docs.oasis-open.org/wsrf/rw-2",
	"wsaw":    "http://www.w3.org/2006/05/addressing/wsdl",
}

//DeviceType alias for int
type DeviceType int

// Onvif Device Tyoe
const (
	NVD DeviceType = iota
	NVS
	NVA
	NVT
)

var (
	ErrorOffline         = fmt.Errorf("device is offline")
	ErrorUnauthorized    = fmt.Errorf("device is unauthorized")
	ErrorServiceNotFound = fmt.Errorf("service not found")
)

func (devType DeviceType) String() string {
	stringRepresentation := []string{
		"NetworkVideoDisplay",
		"NetworkVideoStorage",
		"NetworkVideoAnalytics",
		"NetworkVideoTransmitter",
	}
	i := uint8(devType)
	switch {
	case i <= uint8(NVT):
		return stringRepresentation[i]
	default:
		return strconv.Itoa(int(i))
	}
}

//Device for a new device of onvif and DeviceInfo
//struct represents an abstract ONVIF device.
//It contains methods, which helps to communicate with ONVIF device
type Device struct {
	params    DeviceParams
	Endpoints map[string]string
	info      common.DeviceInfo
	*onvif.EndpointManager
}

type DeviceParams struct {
	Xaddr      string
	Username   string
	Password   string
	HttpClient *http.Client
}

//GetServices return available endpoints
func (dev *Device) GetServices() map[string]string {
	return dev.Endpoints
}

//GetServices return available endpoints
func (dev *Device) GetDeviceInfo() common.DeviceInfo {
	return dev.info
}

func readResponse(resp *http.Response) string {
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	return string(b)
}

//GetAvailableDevicesAtSpecificEthernetInterface ...
func GetAvailableDevicesAtSpecificEthernetInterface(interfaceName string) []Device {
	/*
		Call an ws-discovery Probe Message to Discover NVT type Devices
	*/
	devices := wsdiscovery.SendProbe(interfaceName, nil, []string{"dn:" + NVT.String()}, map[string]string{"dn": "http://www.onvif.org/ver10/network/wsdl"})
	nvtDevices := make([]Device, 0)

	for _, j := range devices {
		doc := etree.NewDocument()
		if err := doc.ReadFromString(j); err != nil {
			fmt.Errorf("%s", err.Error())
			return nil
		}

		endpoints := doc.Root().FindElements("./Body/ProbeMatches/ProbeMatch/XAddrs")
		for _, xaddr := range endpoints {
			xaddr := strings.Split(strings.Split(xaddr.Text(), " ")[0], "/")[2]
			fmt.Println(xaddr)
			c := 0

			for c = 0; c < len(nvtDevices); c++ {
				if nvtDevices[c].params.Xaddr == xaddr {
					fmt.Println(nvtDevices[c].params.Xaddr, "==", xaddr)
					break
				}
			}

			if c < len(nvtDevices) {
				continue
			}

			dev, err := NewDevice(DeviceParams{Xaddr: strings.Split(xaddr, " ")[0]})

			if err != nil {
				fmt.Println("Error", xaddr)
				fmt.Println(err)
				continue
			} else {
				nvtDevices = append(nvtDevices, *dev)
			}
		}
	}

	return nvtDevices
}

func (dev *Device) getSupportedServices(resp *http.Response) {
	doc := etree.NewDocument()

	data, _ := ioutil.ReadAll(resp.Body)

	if err := doc.ReadFromBytes(data); err != nil {
		//log.Println(err.Error())
		return
	}
	services := doc.FindElements("./Envelope/Body/GetCapabilitiesResponse/Capabilities/*/XAddr")
	for _, j := range services {
		dev.addEndpoint(j.Parent().Tag, j.Text())
	}
	services = doc.FindElements("./Envelope/Body/GetCapabilitiesResponse/Capabilities/Extension/*/XAddr")
	for _, j := range services {
		dev.addEndpoint(j.Parent().Tag, j.Text())
	}
}

//NewDevice function construct a ONVIF Device entity
func NewDevice(params DeviceParams) (*Device, error) {
	dev := new(Device)
	dev.params = params
	dev.Endpoints = make(map[string]string)
	url := "http://" + dev.params.Xaddr + "/onvif/device_service"
	dev.addEndpoint("Device", url)

	client := loginClient
	if client == nil {
		client = defaultLoginClient
	}

	getCapabilities := device.GetCapabilities{Category: []common.CapabilityCategory{"All"}}
	resp, err := dev.CallMethod(getCapabilities)
	if err != nil {
		return nil, ErrorOffline
	}
	if resp.StatusCode == http.StatusUnauthorized {
		return nil, ErrorUnauthorized
	}
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("camera is not available at " + dev.params.Xaddr + " or it does not support ONVIF services")
	}
	dev.getSupportedServices(resp)

	getDeviceInfo := device.GetDeviceInformation{}
	resp, err = dev.CallMethod(getDeviceInfo)
	if err != nil {
		return nil, ErrorOffline
	}
	if resp.StatusCode == http.StatusUnauthorized {
		return nil, ErrorUnauthorized
	}
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("camera is not available at " + dev.params.Xaddr + " or it does not support ONVIF services")
	}
	bResp := readResponse(resp)
	soap := gosoap.SoapMessage(bResp)
	data := device.GetDeviceInformationResponse{}
	err = xml.Unmarshal([]byte(soap.Body()), &data)
	if err != nil {
		return nil, fmt.Errorf("unmarshal device: %s info: %s err: %v", params.Xaddr, soap.Body(), err)
	}
	dev.info = common.DeviceInfo{
		Manufacturer:    data.Manufacturer,
		Model:           data.Model,
		FirmwareVersion: data.FirmwareVersion,
		SerialNumber:    data.SerialNumber,
		HardwareId:      data.HardwareId,
	}
	dev.EndpointManager = onvif.NewEndpointManager(dev.Endpoints, dev)
	return dev, nil
}

func (dev *Device) addEndpoint(Key, Value string) {
	//use lowCaseKey
	//make key having ability to handle Mixed Case for Different vendor devcie (e.g. Events EVENTS, events)
	lowCaseKey := strings.ToLower(Key)

	// Replace host with host from device params.
	if u, err := url.Parse(Value); err == nil {
		u.Host = dev.params.Xaddr
		Value = u.String()
	}

	dev.Endpoints[lowCaseKey] = Value
}

//GetEndpoint returns specific ONVIF service endpoint address
func (dev *Device) GetEndpoint(name string) string {
	return dev.Endpoints[name]
}

func (dev Device) buildMethodSOAP(msg string) (gosoap.SoapMessage, error) {
	doc := etree.NewDocument()
	if err := doc.ReadFromString(msg); err != nil {
		//log.Println("Got error")

		return "", err
	}
	element := doc.Root()

	soap := gosoap.NewEmptySOAP()
	soap.AddBodyContent(element)

	return soap, nil
}

//getEndpoint functions get the target service endpoint in a better way
func (dev Device) getEndpoint(endpoint string) (string, error) {

	// common condition, endpointMark in map we use this.
	if endpointURL, bFound := dev.Endpoints[endpoint]; bFound {
		return endpointURL, nil
	}

	//but ,if we have endpoint like event、analytic
	//and sametime the Targetkey like : events、analytics
	//we use fuzzy way to find the best match url
	var endpointURL string
	for targetKey := range dev.Endpoints {
		if strings.Contains(targetKey, endpoint) {
			endpointURL = dev.Endpoints[targetKey]
			return endpointURL, nil
		}
	}
	return endpointURL, errors.New("target endpoint service not found")
}

func (dev Device) GetParam() DeviceParams {
	return DeviceParams{
		Xaddr:      dev.params.Xaddr,
		Username:   dev.params.Username,
		Password:   dev.params.Password,
		HttpClient: nil,
	}
}

//CallMethod functions call an method, defined <method> struct.
//You should use Authenticate method to call authorized requests.
func (dev Device) CallMethod(method interface{}) (*http.Response, error) {
	pkgPath := strings.Split(reflect.TypeOf(method).PkgPath(), "/")
	pkg := strings.ToLower(pkgPath[len(pkgPath)-1])

	endpoint, err := dev.getEndpoint(pkg)
	if err != nil {
		return nil, ErrorServiceNotFound
	}
	return dev.callMethodDo(endpoint, method)
}

func (dev Device) CallMethodUnmarshal(endpoint string, method interface{}, result interface{}) *common.Fault {
	methodKind := reflect.ValueOf(method).Type().Kind()
	resultKind := reflect.ValueOf(result).Type().Kind()
	if (methodKind != reflect.Struct && (methodKind != reflect.Ptr || method == nil)) ||
		resultKind != reflect.Ptr || result == nil {
		return common.NewFault(endpoint, fmt.Errorf("param error, method: %#v, response: %#v", method, result))
	}
	response, err := dev.callMethodDo(dev.Endpoints[endpoint], method)
	if err != nil {
		return common.NewFault(endpoint, err)
	}
	defer func() {
		_ = response.Body.Close()
	}()
	if err != nil {
		return common.NewFault(endpoint, err)
	}
	bResp, _ := io.ReadAll(response.Body)
	soap := gosoap.SoapMessage(bResp)
	b := soap.BodyBytes()
	if response.StatusCode != http.StatusOK {
		fault := common.NewFault(endpoint, fmt.Errorf("http status code is not ok: %d", response.StatusCode))
		_ = xml.Unmarshal(b, fault)
		return fault
	}
	err = xml.Unmarshal(b, result)
	if err != nil {
		return common.NewFault(endpoint, err)
	}
	return nil
}

//CallMethod functions call an method, defined <method> struct with authentication data
func (dev Device) callMethodDo(endpoint string, method interface{}) (*http.Response, error) {
	output, err := xml.MarshalIndent(method, "  ", "    ")
	if err != nil {
		return nil, err
	}

	soap, err := dev.buildMethodSOAP(string(output))
	if err != nil {
		return nil, err
	}

	soap.AddRootNamespaces(Xlmns)
	soap.AddAction()

	//Auth Handling
	if dev.params.Username != "" && dev.params.Password != "" {
		soap.AddWSSecurity(dev.params.Username, dev.params.Password)
	}

	var resp *http.Response
	resp, err = networking.SendSoap(dev.params.HttpClient, endpoint, soap.String())
	if err != nil {
		return resp, ErrorOffline
	}
	return resp, err
}

var (
	loginClient        *http.Client
	defaultLoginClient *http.Client
)

func init() {
	defaultLoginClient = &http.Client{Timeout: time.Second}
	trans := http.DefaultTransport.(*http.Transport).Clone()
	trans.MaxIdleConns = 100
	trans.MaxConnsPerHost = 100
	trans.MaxIdleConnsPerHost = 100
	defaultLoginClient.Transport = trans
}

func SetLoginHttpClient(client *http.Client) {
	loginClient = client
}
