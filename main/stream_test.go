package biz

import (
	"code.byted.org/videoarch/go-onvif/gosoap"
	"code.byted.org/videoarch/go-onvif/media"
	"code.byted.org/videoarch/go-onvif/xsd/onvif"
	"encoding/xml"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestXor(t *testing.T) {
	fmt.Printf("%d", 1^2^3^4^5)
}

func TestGetStreamUri(t *testing.T) {
	req := media.GetStreamUri{
		ProfileToken: "MediaProfile00000",
		StreamSetup: onvif.StreamSetup{
			Stream: "RTP-Unicast",
			Transport: &onvif.Transport{
				Protocol: "TCP",
			},
		},
	}
	resp, err := dev.CallMethod(req)
	assert.Nil(t, err)
	bResp := readResponse(resp)
	soap := gosoap.SoapMessage(bResp)
	data := media.GetStreamUriResponse{}
	err = xml.Unmarshal([]byte(soap.Body()), &data)
	assert.Nil(t, err)
	fmt.Printf("%s", string(bResp))
}

func TestSnapShot(t *testing.T) {
	req := media.GetSnapshotUri{
		ProfileToken: "Profile_1",
	}
	resp, err := dev.CallMethod(req)
	assert.Nil(t, err)
	bResp := readResponse(resp)
	soap := gosoap.SoapMessage(bResp)
	data := media.GetSnapshotUriResponse{}
	err = xml.Unmarshal([]byte(soap.Body()), &data)
	assert.Nil(t, err)
	fmt.Printf("%s", string(bResp))
}

func TestMulticast(t *testing.T) {
	reqStart := media.StartMulticastStreaming{
		ProfileToken: "Profile_1",
	}
	resp, err := dev.CallMethod(reqStart)
	assert.Nil(t, err)
	bResp := readResponse(resp)
	soap := gosoap.SoapMessage(bResp)
	dataStart := media.StartMulticastStreamingResponse{}
	err = xml.Unmarshal([]byte(soap.Body()), &dataStart)
	assert.Nil(t, err)
	fmt.Printf("%s", string(bResp))

	reqStop := media.StopMulticastStreaming{
		ProfileToken: "Profile_1",
	}
	resp, err = dev.CallMethod(reqStop)
	assert.Nil(t, err)
	bResp = readResponse(resp)
	soap = gosoap.SoapMessage(bResp)
	dataStop := media.StopMulticastStreamingResponse{}
	err = xml.Unmarshal([]byte(soap.Body()), &dataStop)
	assert.Nil(t, err)
	fmt.Printf("%s", string(bResp))
}
