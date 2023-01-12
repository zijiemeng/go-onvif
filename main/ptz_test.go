package biz

import (
	"code.byted.org/videoarch/go-onvif/gosoap"
	"code.byted.org/videoarch/go-onvif/media"
	"code.byted.org/videoarch/go-onvif/ptz"
	o "code.byted.org/videoarch/go-onvif/xsd/onvif"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"log"
	"testing"
)

var profile o.ReferenceToken = "Profile_1"

func TestGetPtzNodes(t *testing.T) {
	req := ptz.GetNodes{}
	resp, err := dev.CallMethod(req)
	assert.Nil(t, err)
	bResp := readResponse(resp)
	soap := gosoap.SoapMessage(bResp)
	data := ptz.GetNodesResponse{}
	err = xml.Unmarshal([]byte(soap.Body()), &data)
	assert.Nil(t, err)
	fmt.Printf("%s", string(bResp))
}
func TestGetPtzNode(t *testing.T) {
	req := ptz.GetNode{
		NodeToken: "PTZNODETOKEN",
	}
	resp, err := dev.CallMethod(req)
	assert.Nil(t, err)
	bResp := readResponse(resp)
	soap := gosoap.SoapMessage(bResp)
	data := ptz.GetNodeResponse{}
	err = xml.Unmarshal([]byte(soap.Body()), &data)
	assert.Nil(t, err)
	fmt.Printf("%s", string(bResp))
}

type GetPresetToursResponse struct {
	XMLName    xml.Name `xml:"GetPresetToursResponse"`
	PresetTour []struct {
		Token  string `xml:"token,attr"`
		Status struct {
			State string `xml:"State"`
		} `xml:"Status"`
		AutoStart         bool   `xml:"AutoStart"`
		StartingCondition string `xml:"StartingCondition"`
		TourSpot          []struct {
			PresetDetail struct {
				PresetToken string `xml:"PresetToken"`
				Zoom        struct {
					X     string `xml:"x,attr"`
					Space string `xml:"space,attr"`
				} `xml:"Zoom"`
				Home          bool   `xml:"Home"`
				PTZPosition   string `xml:"PTZPosition"`
				TypeExtension string `xml:"TypeExtension"`
			} `xml:"PresetDetail"`
			Speed struct {
				PanTilt struct {
					X     string `xml:"x,attr"`
					Y     string `xml:"y,attr"`
					Space string `xml:"space,attr"`
				} `xml:"PanTilt"`
				Zoom struct {
					X     string `xml:"x,attr"`
					Space string `xml:"space,attr"`
				} `xml:"Zoom"`
			} `xml:"Speed"`
			StayTime string `xml:"StayTime"`
		} `xml:"TourSpot"`
	} `xml:"PresetTour"`
}

func TestGetPresetTours(t *testing.T) {
	req := ptz.GetPresetTours{
		ProfileToken: "Profile_1",
	}
	resp, err := dev.CallMethod(req)
	assert.Nil(t, err)
	bResp := readResponse(resp)
	soap := gosoap.SoapMessage(bResp)
	err = ioutil.WriteFile("presets.xml", []byte(soap), 0666)
	assert.Nil(t, err)
	data := ptz.GetPresetToursResponse{}
	data0 := GetPresetToursResponse{}
	fmt.Printf(soap.Body())

	err = xml.Unmarshal([]byte(soap.Body()), &data)
	assert.Nil(t, err)
	err = xml.Unmarshal([]byte(soap.Body()), &data0)
	assert.Nil(t, err)

	//data1 := ptz.GetPresetToursResponse{PresetTour: o.PresetTour{
	//	Token: "1",
	//	Name:  "",
	//	Status: o.PTZPresetTourStatus{
	//		State: "Idle",
	//	},
	//	AutoStart: true,
	//	TourSpot: o.PTZPresetTourSpot{
	//		PresetDetail: o.PTZPresetTourPresetDetail{
	//			PresetToken: "6",
	//			Home:        false,
	//		},
	//		Speed: o.PTZSpeed{
	//			PanTilt: o.Vector2D{
	//				X:     0.75,
	//				Y:     0.75,
	//				Space: "http://www.onvif.org/ver10/tptz/PanTiltSpaces/GenericSpeedSpace",
	//			},
	//			Zoom: o.Vector1D{
	//				X:     0.75,
	//				Space: "http://www.onvif.org/ver10/tptz/ZoomSpaces/ZoomGenericSpeedSpace",
	//			},
	//		},
	//		StayTime: "PT120S",
	//	},
	//}}
	//b, _ := xml.Marshal(data1)
	//e := ioutil.WriteFile("presets1.xml", b, 0666)
	//assert.Nil(t, e)
}

//func TestMarshal(t *testing.T) {
//	resp := ptz.GetPresetToursResponse{
//		XMLName: "",
//		PresetTour: []o.PresetTour{
//			o.PresetTour{
//				Token:             "0",
//				Name:              "",
//				Status:            o.PTZPresetTourStatus{},
//				AutoStart:         false,
//				StartingCondition: o.PTZPresetTourStartingCondition{},
//				TourSpot:          o.PTZPresetTourSpot{},
//				Extension:         "",
//			},
//			o.PresetTour{
//				Token:             "1",
//				Name:              "",
//				Status:            o.PTZPresetTourStatus{},
//				AutoStart:         false,
//				StartingCondition: o.PTZPresetTourStartingCondition{},
//				TourSpot:          o.PTZPresetTourSpot{},
//				Extension:         "",
//			},
//		},
//	}
//	b, e := xml.Marshal(resp)
//	assert.Nil(t, e)
//	fmt.Printf(string(b))
//}

func TestGetPresets(t *testing.T) {
	req := ptz.GetPresets{
		ProfileToken: "Profile_1",
	}
	resp, err := dev.CallMethod(req)
	assert.Nil(t, err)
	bResp := readResponse(resp)
	soap := gosoap.SoapMessage(bResp)
	fmt.Println(string(soap))
	data := ptz.GetPresetsResponse{}
	err = xml.Unmarshal([]byte(soap.Body()), &data)
	b, _ := json.Marshal(data)
	fmt.Printf(string(b))
}

func TestGetServiceCapabilities(t *testing.T) {
	req := ptz.GetServiceCapabilities{}
	resp, err := dev.CallMethod(req)
	assert.Nil(t, err)
	bResp := readResponse(resp)
	soap := gosoap.SoapMessage(bResp)
	data := ptz.GetServiceCapabilitiesResponse{}
	err = xml.Unmarshal([]byte(soap.Body()), &data)
	b, _ := json.Marshal(data)
	fmt.Printf(string(b))
}

func TestGetNode(t *testing.T) {
	req := ptz.GetNode{
		NodeToken: "PTZNODETOKEN",
	}
	resp, err := dev.CallMethod(req)
	assert.Nil(t, err)
	bResp := readResponse(resp)
	soap := gosoap.SoapMessage(bResp)
	data := ptz.GetNodesResponse{}
	err = xml.Unmarshal([]byte(soap.Body()), &data)
	assert.Nil(t, err)
	b, _ := json.Marshal(data)
	fmt.Printf(string(b))
}

func TestConfigurations(t *testing.T) {
	req := ptz.GetConfigurations{}
	resp, err := dev.CallMethod(req)
	assert.Nil(t, err)
	bResp := readResponse(resp)
	soap := gosoap.SoapMessage(bResp)
	data := ptz.GetConfigurationsResponse{}
	err = xml.Unmarshal([]byte(soap.Body()), &data)
	assert.Nil(t, err)
	fmt.Printf("%s", string(bResp))
}

func TestConfiguration(t *testing.T) {
	req := ptz.GetConfiguration{
		//XMLName:      "DefaultAbsolutePantTiltPositionSpace",
		ProfileToken: "Profile_1",
	}
	resp, err := dev.CallMethod(req)
	assert.Nil(t, err)
	bResp := readResponse(resp)
	soap := gosoap.SoapMessage(bResp)
	data := ptz.GetConfigurationResponse{}
	err = xml.Unmarshal([]byte(soap.Body()), &data)
	assert.Nil(t, err)
	fmt.Printf("%s", string(bResp))
}

func TestGetConfigurationOptions(t *testing.T) {
	req := ptz.GetConfigurationOptions{
		//XMLName:      "DefaultAbsolutePantTiltPositionSpace",
		ProfileToken: "Profile_2",
	}
	resp, err := dev.CallMethod(req)
	assert.Nil(t, err)
	bResp := readResponse(resp)
	soap := gosoap.SoapMessage(bResp)
	data := ptz.GetConfigurationOptionsResponse{}
	err = xml.Unmarshal([]byte(soap.Body()), &data)
	assert.Nil(t, err)
	fmt.Printf("%s", string(bResp))
}

func TestSetConfiguration(t *testing.T) {
	req := ptz.SetConfiguration{
		PTZConfiguration: o.PTZConfiguration{
			NodeToken: "Profile_1",
		},
		ForcePersistence: false,
	}
	resp, err := dev.CallMethod(req)
	assert.Nil(t, err)
	bResp := readResponse(resp)
	soap := gosoap.SoapMessage(bResp)
	data := ptz.SetConfigurationResponse{}
	err = xml.Unmarshal([]byte(soap.Body()), &data)
	assert.Nil(t, err)
	fmt.Printf("%s", string(bResp))
}

func TestGetCompatibleConfigurations(t *testing.T) {
	req := ptz.GetCompatibleConfigurations{
		ProfileToken: "Profile_1",
	}
	resp, err := dev.CallMethod(req)
	assert.Nil(t, err)
	bResp := readResponse(resp)
	soap := gosoap.SoapMessage(bResp)
	data := ptz.GetCompatibleConfigurationsResponse{}
	err = xml.Unmarshal([]byte(soap.Body()), &data)
	assert.Nil(t, err)
	fmt.Printf("%s", string(bResp))
}

func TestGetStreamURL(t *testing.T) {
	var resp media.GetProfilesResponse
	mediaProfile := media.GetProfiles{}
	mediaProfileResp, err := dev.CallMethod(mediaProfile)
	if err != nil {
		log.Println(err)
	} else {
		content := readResponse(mediaProfileResp)
		soap := gosoap.SoapMessage(content)
		fmt.Printf(string(soap))
		xml.Unmarshal([]byte(soap.Body()), &resp)
	}
	streamUri := media.GetStreamUri{
		ProfileToken: resp.Profiles[0].Token,
	}
	streamResp, err := dev.CallMethod(streamUri)
	if err != nil {
		log.Println(err)
	} else {
		data := gosoap.SoapMessage(readResponse(streamResp))
		var streamResp media.GetStreamUriResponse
		xml.Unmarshal([]byte(data.Body()), &streamResp) //反序列化得到对应的结构体
		fmt.Printf("rstp uri:%s", streamResp.MediaUri.Uri)
	}
}

func TestPresets(t *testing.T) {
	{
		resp, err := dev.CallMethod(ptz.GetPresets{
			ProfileToken: profile,
		})
		assert.Nil(t, err)
		bResp := readResponse(resp)
		soap := gosoap.SoapMessage(bResp)
		fmt.Println(soap)
		data := ptz.GetPresetsResponse{}
		err = xml.Unmarshal([]byte(soap.Body()), &data)
		assert.Nil(t, err)
	}
	//{
	//	for i := 0; i < 300; i++ {
	//		{
	//			x := float64(2*i)/300 - 1
	//			y := float64(2*i)/300 - 1
	//			z := float64(i) / 300
	//			fmt.Printf("x: %f, y: %f, z: %f", x, y, z)
	//			fmt.Println()
	//			resp, err := dev.CallMethod(ptz.AbsoluteMove{
	//				ProfileToken: profile,
	//				Position: o.PTZVector{
	//					PanTilt: o.Vector2D{
	//						X:     x,
	//						Y:     y,
	//						Space: "http://www.onvif.org/ver10/tptz/PanTiltSpaces/PositionGenericSpace",
	//					},
	//					Zoom: o.Vector1D{
	//						X:     z,
	//						Space: "http://www.onvif.org/ver10/tptz/ZoomSpaces/PositionGenericSpace",
	//					},
	//				},
	//				Speed: o.PTZSpeed{
	//					PanTilt: o.Vector2D{
	//						X:     1,
	//						Y:     1,
	//						Space: "http://www.onvif.org/ver10/tptz/PanTiltSpaces/GenericSpeedSpace",
	//					},
	//					Zoom: o.Vector1D{
	//						X:     1,
	//						Space: "http://www.onvif.org/ver10/tptz/ZoomSpaces/ZoomGenericSpeedSpace",
	//					},
	//				},
	//			})
	//			assert.Nil(t, err)
	//			bResp := readResponse(resp)
	//			soap := gosoap.SoapMessage(bResp)
	//			fmt.Println(soap)
	//		}
	//		time.Sleep(time.Second)
	//		{
	//			resp, err := dev.CallMethod(ptz.SetPreset{
	//				ProfileToken: profile,
	//				PresetToken:  o.ReferenceToken(fmt.Sprintf("%d", i+1)),
	//				PresetName:   xsd.String(fmt.Sprintf("test_preset_%d", i)),
	//			})
	//			assert.Nil(t, err)
	//			bResp := readResponse(resp)
	//			soap := gosoap.SoapMessage(bResp)
	//			fmt.Println(soap)
	//		}
	//	}
	//}
	//{
	//	resp, err := dev.CallMethod(ptz.AbsoluteMove{
	//		ProfileToken: profile,
	//		Position: o.PTZVector{
	//			PanTilt: o.Vector2D{
	//				X:     0.25,
	//				Y:     0.25,
	//				Space: "http://www.onvif.org/ver10/tptz/PanTiltSpaces/PositionGenericSpace",
	//			},
	//			Zoom: o.Vector1D{
	//				X:     0.25,
	//				Space: "http://www.onvif.org/ver10/tptz/ZoomSpaces/PositionGenericSpace",
	//			},
	//		},
	//		Speed: o.PTZSpeed{
	//			PanTilt: o.Vector2D{
	//				X:     0.255,
	//				Y:     0.255,
	//				Space: "http://www.onvif.org/ver10/tptz/PanTiltSpaces/GenericSpeedSpace",
	//			},
	//			Zoom: o.Vector1D{
	//				X:     0.255,
	//				Space: "http://www.onvif.org/ver10/tptz/ZoomSpaces/ZoomGenericSpeedSpace",
	//			},
	//		},
	//	})
	//	assert.Nil(t, err)
	//	bResp := readResponse(resp)
	//	soap := gosoap.SoapMessage(bResp)
	//	fmt.Println(soap)
	//}
	//{
	//	resp, err := dev.CallMethod(ptz.SetPreset{
	//		ProfileToken: profile,
	//		PresetName:   "test_preset_55",
	//		PresetToken:  "55",
	//	})
	//	assert.Nil(t, err)
	//	bResp := readResponse(resp)
	//	soap := gosoap.SoapMessage(bResp)
	//	fmt.Println(soap)
	//}
}

func TestPresetTours(t *testing.T) {
	{
		resp, err := dev.CallMethod(ptz.GetPresetTours{
			ProfileToken: profile,
		})
		assert.Nil(t, err)
		bResp := readResponse(resp)
		soap := gosoap.SoapMessage(bResp)
		fmt.Println(soap)
	}
	//{
	//	resp, err := dev.CallMethod(ptz.CreatePresetTour{
	//		ProfileToken: profile,
	//	})
	//	assert.Nil(t, err)
	//	bResp := readResponse(resp)
	//	soap := gosoap.SoapMessage(bResp)
	//	fmt.Println(soap)
	//}
	tour := o.PresetTour{}
	{
		resp, err := dev.CallMethod(ptz.GetPresetTour{
			ProfileToken:    profile,
			PresetTourToken: "1",
		})
		assert.Nil(t, err)
		bResp := readResponse(resp)
		soap := gosoap.SoapMessage(bResp)
		fmt.Println(soap)
		data := ptz.GetPresetTourResponse{}
		err = xml.Unmarshal(soap.BodyBytes(), &data)
		assert.Nil(t, err)
		tour = data.PresetTour
	}
	var presets []o.PTZPreset
	{
		resp, err := dev.CallMethod(ptz.GetPresets{
			ProfileToken: profile,
		})
		assert.Nil(t, err)
		bResp := readResponse(resp)
		soap := gosoap.SoapMessage(bResp)
		fmt.Println(soap)
		data := ptz.GetPresetsResponse{}
		err = xml.Unmarshal([]byte(soap.Body()), &data)
		presets = data.Preset
		assert.Nil(t, err)
	}
	{
		tour.TourSpot = []o.PTZPresetTourSpot{
			{
				PresetDetail: o.PTZPresetTourPresetDetail{
					PresetToken:   presets[0].Token,
					Home:          true,
					PTZPosition:   presets[0].PTZPosition,
					TypeExtension: "",
				},
				Speed: o.PTZSpeed{
					PanTilt: o.Vector2D{
						X:     0,
						Y:     0,
						Space: "http://www.onvif.org/ver10/tptz/PanTiltSpaces/GenericSpeedSpace",
					},
					Zoom: o.Vector1D{
						X:     0,
						Space: "http://www.onvif.org/ver10/tptz/ZoomSpaces/ZoomGenericSpeedSpace",
					},
				},
				StayTime:  "PT2S",
				Extension: "",
			},
			{
				PresetDetail: o.PTZPresetTourPresetDetail{
					PresetToken:   presets[1].Token,
					Home:          false,
					PTZPosition:   presets[1].PTZPosition,
					TypeExtension: "",
				},
				Speed: o.PTZSpeed{
					PanTilt: o.Vector2D{
						X:     0,
						Y:     0,
						Space: "http://www.onvif.org/ver10/tptz/PanTiltSpaces/GenericSpeedSpace",
					},
					Zoom: o.Vector1D{
						X:     0,
						Space: "http://www.onvif.org/ver10/tptz/ZoomSpaces/ZoomGenericSpeedSpace",
					},
				},
				StayTime:  "PT2S",
				Extension: "",
			},
		}
		resp, err := dev.CallMethod(ptz.ModifyPresetTour{
			ProfileToken: profile,
			PresetTour:   tour,
		})
		assert.Nil(t, err)
		bResp := readResponse(resp)
		soap := gosoap.SoapMessage(bResp)
		fmt.Printf(string(soap))
		fmt.Println()
	}
}
