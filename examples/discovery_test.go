package example

import (
	"code.byted.org/videoarch/go-onvif"
	"encoding/json"
	"io/ioutil"
	"log"
	"path"
	"regexp"
	"strings"
	"testing"
	"time"

	"code.byted.org/videoarch/go-onvif/device"
	discover "code.byted.org/videoarch/go-onvif/ws-discovery"
	"github.com/beevik/etree"
)

func TestGetAvailableDevicesAtSpecificEthernetInterface(t *testing.T) {

	// client()
	// runDiscovery("en0")

	for i := 0; i < 10; i++ {
		s := onvif.GetAvailableDevicesAtSpecificEthernetInterface("en5")
		log.Printf("%v", s)
		time.Sleep(2 * time.Second)
	}
}

func client() {
	dev, err := onvif.NewDevice(onvif.DeviceParams{Xaddr: "192.168.3.10", Username: "admin", Password: "zsyy12345"})
	if err != nil {
		panic(err)
	}

	log.Printf("output %+v", dev.GetServices())

	res, err := dev.CallMethod(device.GetUsers{})
	bs, _ := ioutil.ReadAll(res.Body)
	log.Printf("output %+v %s", res.StatusCode, bs)
}

// Host host
type Host struct {
	URL  string `json:"url"`
	Name string `json:"name"`
}

func runDiscovery(interfaceName string) {
	var hosts []*Host
	devices := discover.SendProbe(interfaceName, nil, []string{"dn:NetworkVideoTransmitter"}, map[string]string{"dn": "http://www.onvif.org/ver10/network/wsdl"})
	for _, j := range devices {
		doc := etree.NewDocument()
		if err := doc.ReadFromString(j); err != nil {
			log.Printf("error %s", err)
		} else {

			endpoints := doc.Root().FindElements("./Body/ProbeMatches/ProbeMatch/XAddrs")
			scopes := doc.Root().FindElements("./Body/ProbeMatches/ProbeMatch/Scopes")

			flag := false

			host := &Host{}

			for _, xaddr := range endpoints {
				xaddr := strings.Split(strings.Split(xaddr.Text(), " ")[0], "/")[2]
				host.URL = xaddr
			}
			if flag {
				break
			}
			for _, scope := range scopes {
				re := regexp.MustCompile(`onvif:\/\/www\.onvif\.org\/name\/[A-Za-z0-9-]+`)
				match := re.FindStringSubmatch(scope.Text())
				host.Name = path.Base(match[0])
			}

			hosts = append(hosts, host)

		}

	}

	bys, _ := json.Marshal(hosts)
	log.Printf("done %s", bys)
}
