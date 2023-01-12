package biz

import (
	imaging "code.byted.org/videoarch/go-onvif/Imaging"
	"code.byted.org/videoarch/go-onvif/gosoap"
	"encoding/xml"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestImageInfo(t *testing.T) {
	req := imaging.GetImagingSettings{VideoSourceToken: "VideoSource_1"}
	resp, err := dev.CallMethod(req)
	assert.Nil(t, err)
	bResp := readResponse(resp)
	soap := gosoap.SoapMessage(bResp)
	fmt.Println(soap)
	data := imaging.GetImagingSettingsResponse{}
	err = xml.Unmarshal([]byte(soap.Body()), &data)
	assert.Nil(t, err)
}

func TestGetOptions(t *testing.T) {
	req := imaging.GetOptions{VideoSourceToken: "VideoSource_1"}
	resp, err := dev.CallMethod(req)
	assert.Nil(t, err)
	bResp := readResponse(resp)
	soap := gosoap.SoapMessage(bResp)
	data := imaging.GetOptionsResponse{}
	source := soap.Body()
	fmt.Println(source)
	err = xml.Unmarshal([]byte(source), &data)
	assert.Nil(t, err)
}

func TestGetMoveOptions(t *testing.T) {
	req := imaging.GetMoveOptions{VideoSourceToken: "VideoSource_1"}
	resp, err := dev.CallMethod(req)
	assert.Nil(t, err)
	bResp := readResponse(resp)
	soap := gosoap.SoapMessage(bResp)
	data := imaging.GetMoveOptionsResponse{}
	source := soap.Body()
	fmt.Println(source)
	err = xml.Unmarshal([]byte(source), &data)
	assert.Nil(t, err)
}

func TestGetCap(t *testing.T) {
	req := imaging.GetServiceCapabilities{}
	resp, err := dev.CallMethod(req)
	assert.Nil(t, err)
	bResp := readResponse(resp)
	soap := gosoap.SoapMessage(bResp)
	data := imaging.GetServiceCapabilitiesResponse{}
	source := soap.Body()
	fmt.Println(source)
	err = xml.Unmarshal([]byte(source), &data)
	assert.Nil(t, err)
}

func TestGetStatus(t *testing.T) {
	req := imaging.GetStatus{
		VideoSourceToken: "00000",
	}
	resp, err := dev.CallMethod(req)
	assert.Nil(t, err)
	bResp := readResponse(resp)
	soap := gosoap.SoapMessage(bResp)
	data := imaging.GetStatusResponse{}
	source := soap.Body()
	fmt.Println(source)
	err = xml.Unmarshal([]byte(source), &data)
	assert.Nil(t, err)
}
