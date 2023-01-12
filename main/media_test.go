package biz

import (
	"code.byted.org/videoarch/go-onvif/gosoap"
	"code.byted.org/videoarch/go-onvif/media"
	"code.byted.org/videoarch/go-onvif/xsd/onvif"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/fs"
	"io/ioutil"
	"testing"
)

func TestGetProfiles(t *testing.T) {
	req := media.GetProfiles{}
	resp, err := dev.CallMethod(req)
	assert.Nil(t, err)
	bResp := readResponse(resp)
	soap := gosoap.SoapMessage(bResp)
	data := media.GetProfilesResponse{}
	err = xml.Unmarshal([]byte(soap.Body()), &data)
	assert.Nil(t, err)
	b, _ := json.Marshal(data)
	err = ioutil.WriteFile("profile.xml", bResp, fs.ModePerm)
	assert.Nil(t, err)
	fmt.Printf(string(b))
}

func TestGetProfile(t *testing.T) {
	req := media.DeleteProfile{
		ProfileToken: "MediaProfile00002",
	}
	resp, err := dev.CallMethod(req)
	assert.Nil(t, err)
	bResp := readResponse(resp)
	soap := gosoap.SoapMessage(bResp)
	data := media.GetProfileResponse{}
	err = xml.Unmarshal([]byte(soap.Body()), &data)
	assert.Nil(t, err)
	b, _ := json.Marshal(data)
	assert.Nil(t, err)
	fmt.Printf(string(b))
}

func TestCreateProfiles(t *testing.T) {
	req := media.CreateProfile{
		Name:  "Profile_Test",
		Token: "Profile_Test",
	}
	resp, err := dev.CallMethod(req)
	assert.Nil(t, err)
	bResp := readResponse(resp)
	soap := gosoap.SoapMessage(bResp)
	data := media.CreateProfileResponse{}
	err = xml.Unmarshal([]byte(soap.Body()), &data)
	assert.Nil(t, err)
	fmt.Printf("%s", string(bResp))
}

func TestDeleteProfile(t *testing.T) {
	req := media.DeleteProfile{
		ProfileToken: "Profile_Test",
	}
	resp, err := dev.CallMethod(req)
	assert.Nil(t, err)
	bResp := readResponse(resp)
	soap := gosoap.SoapMessage(bResp)
	data := media.DeleteProfileResponse{}
	err = xml.Unmarshal([]byte(soap.Body()), &data)
	assert.Nil(t, err)
	fmt.Printf("%s", string(bResp))
}

func TestAddVideoSourceToProfile(t *testing.T) {
	req := media.AddVideoSourceConfiguration{
		ProfileToken:       "Profile_Test",
		ConfigurationToken: "VideoSourceToken",
	}
	resp, err := dev.CallMethod(req)
	assert.Nil(t, err)
	bResp := readResponse(resp)
	soap := gosoap.SoapMessage(bResp)
	data := media.AddVideoSourceConfigurationResponse{}
	err = xml.Unmarshal([]byte(soap.Body()), &data)
	assert.Nil(t, err)
	fmt.Printf("%s", string(bResp))
}

func TestAddAudioSourceToProfile(t *testing.T) {
	req := media.AddAudioSourceConfiguration{
		ProfileToken:       "Profile_Test",
		ConfigurationToken: "AudioSourceConfigToken",
	}
	resp, err := dev.CallMethod(req)
	assert.Nil(t, err)
	bResp := readResponse(resp)
	soap := gosoap.SoapMessage(bResp)
	data := media.AddVideoSourceConfigurationResponse{}
	err = xml.Unmarshal([]byte(soap.Body()), &data)
	assert.Nil(t, err)
	fmt.Printf("%s", string(bResp))
}

func TestAddVideoEncodeToProfile(t *testing.T) {
	req := media.AddVideoEncoderConfiguration{
		ProfileToken:       "Profile_Test",
		ConfigurationToken: "VideoEncoderToken_1",
	}
	resp, err := dev.CallMethod(req)
	assert.Nil(t, err)
	bResp := readResponse(resp)
	soap := gosoap.SoapMessage(bResp)
	data := media.AddVideoSourceConfigurationResponse{}
	err = xml.Unmarshal([]byte(soap.Body()), &data)
	assert.Nil(t, err)
	fmt.Printf("%s", string(bResp))
}

func TestAddAudioEncodeToProfile(t *testing.T) {
	req := media.AddAudioEncoderConfiguration{
		ProfileToken:       "Profile_Test",
		ConfigurationToken: "MainAudioEncoderToken",
	}
	resp, err := dev.CallMethod(req)
	assert.Nil(t, err)
	bResp := readResponse(resp)
	soap := gosoap.SoapMessage(bResp)
	data := media.AddVideoSourceConfigurationResponse{}
	err = xml.Unmarshal([]byte(soap.Body()), &data)
	assert.Nil(t, err)
	fmt.Printf("%s", string(bResp))
}

func TestGetVideoSources(t *testing.T) {
	req := media.GetVideoSources{}
	resp, err := dev.CallMethod(req)
	assert.Nil(t, err)
	bResp := readResponse(resp)
	soap := gosoap.SoapMessage(bResp)
	data := media.GetVideoSourcesResponse{}
	err = xml.Unmarshal([]byte(soap.Body()), &data)
	assert.Nil(t, err)
	fmt.Printf("%s", string(bResp))
}

func TestGetVideoSourceConfigurations(t *testing.T) {
	req := media.GetVideoSourceConfigurations{}
	resp, err := dev.CallMethod(req)
	assert.Nil(t, err)
	bResp := readResponse(resp)
	soap := gosoap.SoapMessage(bResp)
	data := media.GetVideoSourceConfigurationsResponse{}
	err = xml.Unmarshal([]byte(soap.Body()), &data)
	assert.Nil(t, err)
	fmt.Printf("%s", string(bResp))
}

func TestGetAudioSourceConfigurations(t *testing.T) {
	req := media.GetAudioSourceConfigurations{}
	resp, err := dev.CallMethod(req)
	assert.Nil(t, err)
	bResp := readResponse(resp)
	soap := gosoap.SoapMessage(bResp)
	data := media.GetAudioSourceConfigurationsResponse{}
	err = xml.Unmarshal([]byte(soap.Body()), &data)
	assert.Nil(t, err)
	fmt.Printf("%s", string(bResp))
}

func TestGetVideoSourceConfiguration(t *testing.T) {
	req := media.GetVideoSourceConfiguration{
		ConfigurationToken: "VideoSourceToken",
	}
	resp, err := dev.CallMethod(req)
	assert.Nil(t, err)
	bResp := readResponse(resp)
	soap := gosoap.SoapMessage(bResp)
	data := media.GetVideoSourceConfigurationResponse{}
	err = xml.Unmarshal([]byte(soap.Body()), &data)
	assert.Nil(t, err)
	fmt.Printf("%s", string(bResp))
}

func TestGetVideoEncodeConfiguration(t *testing.T) {
	var conf onvif.VideoEncoderConfiguration
	{
		req := media.GetVideoEncoderConfiguration{
			ConfigurationToken: "VideoEncoderToken_1",
		}
		resp, err := dev.CallMethod(req)
		assert.Nil(t, err)
		bResp := readResponse(resp)
		soap := gosoap.SoapMessage(bResp)
		data := media.GetVideoEncoderConfigurationResponse{}
		err = xml.Unmarshal([]byte(soap.Body()), &data)
		assert.Nil(t, err)
		conf = data.Configuration
		fmt.Printf("%s", string(bResp))
	}
	{
		conf.Token = "VideoEncoderToken_3"
		conf.Name = "VideoEncoder_3"
		req := media.SetVideoEncoderConfiguration{
			Configuration:    conf,
			ForcePersistence: false,
		}
		resp, err := dev.CallMethod(req)
		bResp := readResponse(resp)
		soap := gosoap.SoapMessage(bResp)
		data := media.SetVideoEncoderConfigurationResponse{}
		err = xml.Unmarshal([]byte(soap.Body()), &data)
		assert.Nil(t, err)
		fmt.Printf("%s", string(bResp))
	}
}

func TestVideoEncoderConfiguration(t *testing.T) {
	var opt *onvif.VideoEncoderConfigurationOptions
	{
		req := media.GetVideoEncoderConfigurationOptions{
			ConfigurationToken: "VideoEncoderToken_2",
		}
		resp, err := dev.CallMethod(req)
		assert.Nil(t, err)
		bResp := readResponse(resp)
		soap := gosoap.SoapMessage(bResp)
		data := media.GetVideoEncoderConfigurationOptionsResponse{}
		err = xml.Unmarshal([]byte(soap.Body()), &data)
		assert.Nil(t, err)
		opt = &data.Options
		fmt.Println(string(bResp))
	}
	fmt.Println(opt)
	var conf *onvif.VideoEncoderConfiguration
	{
		req := media.GetVideoEncoderConfiguration{
			ConfigurationToken: "VideoEncoderToken_2",
		}
		resp, err := dev.CallMethod(req)
		assert.Nil(t, err)
		bResp := readResponse(resp)
		soap := gosoap.SoapMessage(bResp)
		data := media.GetVideoEncoderConfigurationResponse{}
		err = xml.Unmarshal([]byte(soap.Body()), &data)
		assert.Nil(t, err)
		conf = &data.Configuration
		fmt.Println(string(bResp))
	}
	println(conf)
	{
		req := media.SetVideoEncoderConfiguration{
			Configuration: onvif.VideoEncoderConfiguration{
				ConfigurationEntity: onvif.ConfigurationEntity{
					Token:    conf.Token,
					Name:     conf.Name,
					UseCount: conf.UseCount,
				},
				Quality: 3,
			},
		}
		resp, err := dev.CallMethod(req)
		assert.Nil(t, err)
		bResp := readResponse(resp)
		soap := gosoap.SoapMessage(bResp)
		data := media.SetVideoEncoderConfigurationResponse{}
		err = xml.Unmarshal([]byte(soap.Body()), &data)
		assert.Nil(t, err)
		fmt.Println(string(bResp))
	}
	{
		req := media.GetVideoEncoderConfiguration{
			ConfigurationToken: "VideoEncoderToken_2",
		}
		resp, err := dev.CallMethod(req)
		assert.Nil(t, err)
		bResp := readResponse(resp)
		soap := gosoap.SoapMessage(bResp)
		data := media.GetVideoEncoderConfigurationResponse{}
		err = xml.Unmarshal([]byte(soap.Body()), &data)
		assert.Nil(t, err)
		conf = &data.Configuration
		fmt.Println(string(bResp))
	}
	println(conf)
}

func TestGetVideoEncodeConfigurations(t *testing.T) {
	req := media.GetVideoEncoderConfigurations{}
	resp, err := dev.CallMethod(req)
	assert.Nil(t, err)
	bResp := readResponse(resp)
	soap := gosoap.SoapMessage(bResp)
	data := media.GetVideoEncoderConfigurationsResponse{}
	err = xml.Unmarshal([]byte(soap.Body()), &data)
	assert.Nil(t, err)
	fmt.Printf("%s", string(bResp))
}

func TestGetAudioEncodeConfiguration(t *testing.T) {
	req := media.GetAudioEncoderConfigurations{}
	resp, err := dev.CallMethod(req)
	assert.Nil(t, err)
	bResp := readResponse(resp)
	soap := gosoap.SoapMessage(bResp)
	data := media.GetAudioEncoderConfigurationsResponse{}
	err = xml.Unmarshal([]byte(soap.Body()), &data)
	assert.Nil(t, err)
	fmt.Printf("%s", string(bResp))
}

func TestGetCompatibleVideoEncodeConfiguration(t *testing.T) {
	req := media.GetCompatibleVideoEncoderConfigurations{
		ProfileToken: "Profile_1",
	}
	resp, err := dev.CallMethod(req)
	assert.Nil(t, err)
	bResp := readResponse(resp)
	soap := gosoap.SoapMessage(bResp)
	data := media.GetCompatibleVideoEncoderConfigurationsResponse{}
	err = xml.Unmarshal([]byte(soap.Body()), &data)
	assert.Nil(t, err)
	fmt.Printf("%s", string(bResp))
}

func TestGetCompatibleVideoSourceConfigurations(t *testing.T) {
	req := media.GetCompatibleVideoSourceConfigurations{
		ProfileToken: "Profile_1",
	}
	resp, err := dev.CallMethod(req)
	assert.Nil(t, err)
	bResp := readResponse(resp)
	soap := gosoap.SoapMessage(bResp)
	data := media.GetCompatibleVideoSourceConfigurationsResponse{}
	err = xml.Unmarshal([]byte(soap.Body()), &data)
	assert.Nil(t, err)
	fmt.Printf("%s", string(bResp))
}

func TestGetVideoEncodeConfigurationOptions(t *testing.T) {
	req := media.GetVideoEncoderConfigurationOptions{
		ProfileToken: "Profile_1",
	}
	resp, err := dev.CallMethod(req)
	assert.Nil(t, err)
	bResp := readResponse(resp)
	soap := gosoap.SoapMessage(bResp)
	data := media.GetVideoEncoderConfigurationOptionsResponse{}
	err = xml.Unmarshal([]byte(soap.Body()), &data)
	assert.Nil(t, err)
	fmt.Printf("%s", string(bResp))
}

func TestGetVideoSourceConfigurationOptions(t *testing.T) {
	req := media.GetVideoSourceConfigurationOptions{
		ProfileToken: "Profile_1",
	}
	resp, err := dev.CallMethod(req)
	assert.Nil(t, err)
	bResp := readResponse(resp)
	soap := gosoap.SoapMessage(bResp)
	data := media.GetVideoSourceConfigurationOptionsResponse{}
	err = xml.Unmarshal([]byte(soap.Body()), &data)
	assert.Nil(t, err)
	fmt.Printf("%s", string(bResp))
}

func TestSetSynchronizationPoint(t *testing.T) {
	req := media.SetSynchronizationPoint{
		ProfileToken: "Profile_1",
	}
	resp, err := dev.CallMethod(req)
	assert.Nil(t, err)
	bResp := readResponse(resp)
	soap := gosoap.SoapMessage(bResp)
	data := media.SetSynchronizationPointResponse{}
	err = xml.Unmarshal([]byte(soap.Body()), &data)
	assert.Nil(t, err)
	fmt.Printf("%s", string(bResp))
}

func TestGetVideoSourceModes(t *testing.T) {
	req := media.GetVideoSourceModes{
		VideoSourceToken: "VideoSource_1", // todo: token问题
	}
	resp, err := dev.CallMethod(req)
	assert.Nil(t, err)
	bResp := readResponse(resp)
	soap := gosoap.SoapMessage(bResp)
	data := media.GetVideoSourceModesResponse{}
	err = xml.Unmarshal([]byte(soap.Body()), &data)
	assert.Nil(t, err)
	fmt.Printf("%s", string(bResp))
}

func TestGetOSDs(t *testing.T) {
	req := media.GetOSDs{
		ConfigurationToken: "VideoSourceToken",
	}
	resp, err := dev.CallMethod(req)
	assert.Nil(t, err)
	bResp := readResponse(resp)
	soap := gosoap.SoapMessage(bResp)
	data := media.GetOSDsResponse{}
	err = xml.Unmarshal([]byte(soap.Body()), &data)
	assert.Nil(t, err)
	fmt.Printf("%s", string(bResp))
}

func TestGetOSDOptions(t *testing.T) {
	req := media.GetOSDOptions{
		ConfigurationToken: "VideoSourceToken",
	}
	resp, err := dev.CallMethod(req)
	assert.Nil(t, err)
	bResp := readResponse(resp)
	soap := gosoap.SoapMessage(bResp)
	data := media.GetOSDOptionsResponse{}
	err = xml.Unmarshal([]byte(soap.Body()), &data)
	assert.Nil(t, err)
	fmt.Printf("%s", string(bResp))
}

func TestCreateOSD(t *testing.T) {
	req := media.CreateOSD{
		OSD: onvif.OSDConfiguration{
			DeviceEntity: onvif.DeviceEntity{
				Token: "TestOsdToken1",
			},
			VideoSourceConfigurationToken: "VideoSourceToken", // todo: invalid value
			Type:                          "Text",
		},
	}
	resp, err := dev.CallMethod(req)
	assert.Nil(t, err)
	bResp := readResponse(resp)
	soap := gosoap.SoapMessage(bResp)
	data := media.CreateOSDResponse{}
	err = xml.Unmarshal([]byte(soap.Body()), &data)
	assert.Nil(t, err)
	fmt.Printf("%s", string(bResp))
}
