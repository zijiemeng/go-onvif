package onvif

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/zijiemeng/go-onvif/onvif/v10/device"
	"github.com/zijiemeng/go-onvif/onvif/v10/media"
	"net/http"
	"testing"
	"time"
)

const (
	address  = "169.254.80.77"
	username = "test"
	password = "a12345678"
)

func TestOnvif(t *testing.T) {
	dev, err := NewDevice(DeviceParams{
		Xaddr:      address,
		Username:   username,
		Password:   password,
		HttpClient: &http.Client{Timeout: 6 * time.Second},
	})
	require.NoError(t, err)
	require.NotNil(t, dev)
	mgr := dev.EndpointManager
	eptDevice := mgr.Device
	require.NotNil(t, eptDevice)

	capabilities, fault := eptDevice.OptGetCapabilities(device.GetCapabilities{})
	require.Nil(t, fault)
	require.NotNil(t, capabilities)
	printJson(capabilities)

	ntp, fault := eptDevice.OptGetNTP(device.GetNTP{})
	require.Nil(t, fault)
	require.NotNil(t, ntp)
	printJson(ntp)

	eptMedia := mgr.Media
	if eptMedia != nil {
		capabilities, fault := eptMedia.Media.OptGetServiceCapabilities(media.GetServiceCapabilities{})
		require.Nil(t, fault)
		require.NotNil(t, capabilities)
		printJson(capabilities)
	}
}

func printJson(source interface{}) {
	b, _ := json.Marshal(source)
	fmt.Println(string(b))
}

type AItf interface {
}

type A struct {
	AItf
}

func AssertAItf(t *testing.T, a AItf, b AItf) {
	assert.True(t, a == b)
}

func TestEqual(t *testing.T) {
	var aItf AItf
	var a *A

	var b = AItf(a)
	var c AItf
	c = AItf(a)

	//aItf = c
	t.Logf("a %#v %#v", a, &a)
	t.Logf("b %#v %#v", b, &b)
	t.Logf("c %#v %#v", c, &c)
	t.Logf("aitf %#v %#v", aItf, &aItf)
	t.Log("=======")

	AssertAItf(t, aItf, a)
	t.Logf("aitf b")
	AssertAItf(t, aItf, b)
	t.Logf("aitf c")
	AssertAItf(t, c, aItf)
	t.Logf("b c")
	AssertAItf(t, c, b)
}
