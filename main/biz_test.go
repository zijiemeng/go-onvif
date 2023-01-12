package biz

import (
	"code.byted.org/videoarch/go-onvif"
	"code.byted.org/videoarch/go-onvif/gosoap"
	"code.byted.org/videoarch/go-onvif/ptz"
	o "code.byted.org/videoarch/go-onvif/xsd/onvif"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
	"time"
)

func setUp() error {
	now := time.Now()
	dev, err = onvif.NewDevice(onvif.DeviceParams{
		Xaddr:      address,
		Username:   username,
		Password:   password,
		HttpClient: &http.Client{Timeout: 6 * time.Second},
	})
	fmt.Printf("%s\n", time.Since(now))
	if err != nil {
		return err
	}
	if dev == nil {
		return fmt.Errorf("device is nil")
	}
	return nil
}

func tearDown() error {
	return nil
}

func TestMain(m *testing.M) {
	err = setUp()
	if err != nil {
		panic(err)
	}
	m.Run()
	err = tearDown()
	if err != nil {
		panic(err)
	}
}

func TestPTZ(t *testing.T) {
	req := ptz.ContinuousMove{
		ProfileToken: "Profile_1",
		Velocity: o.PTZSpeed{
			PanTilt: o.Vector2D{
				X: 0.2,
				Y: 0.2,
			},
			Zoom: o.Vector1D{
				X: 0.1,
			},
		},
		Timeout: "",
	}
	resp, err := dev.CallMethod(req)
	assert.Nil(t, err)
	bResp := readResponse(resp)
	soap := gosoap.SoapMessage(bResp)
	data := ptz.ContinuousMoveResponse{}
	err = xml.Unmarshal([]byte(soap.Body()), &data)
	b, _ := json.Marshal(data)
	fmt.Printf(string(b))
}
