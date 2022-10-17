package main

import (
	"code.byted.org/videoarch/go-onvif"
	"code.byted.org/videoarch/go-onvif/gosoap"
	"code.byted.org/videoarch/go-onvif/media"
	t "code.byted.org/videoarch/go-onvif/xsd/onvif"
	"encoding/xml"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	ip := flag.String("ip", "", "device ip")
	port := flag.Int("port", -1, "device port")
	user := flag.String("user", "", "device account")
	password := flag.String("pwd", "", "device password")
	flag.Parse()
	if *ip == "" {
		fmt.Println("error: ip is empty")
		os.Exit(1)
	}
	if *user == "" {
		fmt.Println("error: user is empty")
		os.Exit(1)
	}
	addr := *ip
	if *port > 0 {
		addr = fmt.Sprintf("%s:%d", *ip, *port)
	}
	device, e := onvif.NewDevice(onvif.DeviceParams{
		Xaddr:      addr,
		Username:   *user,
		Password:   *password,
		HttpClient: nil,
	})
	if e != nil {
		fmt.Println("onvif device error:", e.Error())
		os.Exit(1)
	}
	fmt.Println("manufacturer :", device.GetDeviceInfo().Manufacturer)
	profiles := media.GetProfilesResponse{}
	{
		req := media.GetProfiles{}
		resp, err := device.CallMethod(req)
		if err != nil {
			fmt.Println("get profiles error:", e.Error())
			os.Exit(1)
		}
		bResp, _ := ioutil.ReadAll(resp.Body)
		soap := gosoap.SoapMessage(bResp)
		err = xml.Unmarshal([]byte(soap.Body()), &profiles)
		if err != nil {
			fmt.Println("get stream url xml error:", e.Error(), soap.Body())
			os.Exit(1)
		}
	}
	for _, profile := range profiles.Profiles {
		req := media.GetStreamUri{
			ProfileToken: profile.Token,
			StreamSetup: t.StreamSetup{
				Stream: "RTP-Unicast",
				Transport: &t.Transport{
					Protocol: "TCP",
				},
			},
		}
		resp, err := device.CallMethod(req)
		if err != nil {
			fmt.Println("get stream url error:", e.Error())
			os.Exit(1)
		}
		bResp, _ := ioutil.ReadAll(resp.Body)
		soap := gosoap.SoapMessage(bResp)
		data := media.GetStreamUriResponse{}
		err = xml.Unmarshal([]byte(soap.Body()), &data)
		if err != nil {
			fmt.Println("get stream url xml error:", e.Error(), soap.Body())
			os.Exit(1)
		}
		fmt.Println(profile.Token, ":", data.MediaUri.Uri)
	}
}
