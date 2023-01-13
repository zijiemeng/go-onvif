# Onvif Golang Client SDK

# 使用方法
```shell
go get code.byted.org/videoarch/go-onvif
```

# 支持协议
```shell
https://github.com/onvif/specs.git
branch: 22.12
```

# 使用方法
```go
device, _ := NewDevice(DeviceParams{
    Xaddr:      address,    // IP:Port
    Username:   username,
    Password:   password,
    HttpClient: &http.Client{Timeout: 6 * time.Second},
})

endpoints := dev.EndpointManager
eptDevice := endpoints.Device
if eptDevice != nil {
    capabilities, _ := eptDevice.OptGetCapabilities(device.GetCapabilities{})
	fmt.Printf("%#v\n", capabilities)
}

eptMedia := endpoints.Media // endpoints.Media.Media -> onvif v10, endpoints.Media.Media2 -> onvif v20
if eptMedia != nil && eptMedia.Media != nil {
    capabilities, _ := eptMedia.Media.OptGetServiceCapabilities(media.GetServiceCapabilities{})
    fmt.Printf("%#v\n", capabilities)
}
```