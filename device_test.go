package onvif

import (
	"code.byted.org/videoarch/go-onvif/onvif/v10/device"
	"github.com/stretchr/testify/require"
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

	capabilities, err := eptDevice.OptGetCapabilities(device.GetCapabilities{})
	require.NoError(t, err)
	require.NotNil(t, capabilities)
}