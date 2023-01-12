package biz

import (
	"code.byted.org/gopkg/logs"
	"code.byted.org/videoarch/go-onvif"
	"code.byted.org/videoarch/go-onvif/gosoap"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
)

const (
	//address  = "192.168.3.3"
	//username = "admin"
	//password = "bytedance123"
	//address  = "10.93.205.29"
	//username = "admin"
	//password = "a12345678"
	//address  = "10.193.205.178"
	//username = "admin"
	//password = "a12345678"
	//address  = "10.93.205.29"
	address  = "169.254.80.77"
	username = "admin"
	password = "a12345678"
	//address  = "10.76.12.223"
	//username = "admin"
	//password = "bytedance123"
	//address  = "10.94.50.207"
	//username = "admin"
	//password = "a12345678"
	//password = "a123456781"
)

var (
	dev        *onvif.Device = nil
	err        error         = nil
	reflectMap map[string]*FuncNode
)

type FuncNode struct {
	Name string
	Req  reflect.Type
	Resp reflect.Type
}

func Call(req interface{}, resp interface{}) error {
	err := CheckParam(req, resp)
	if err != nil {
		return err
	}
	r, err := dev.CallMethod(req)
	if err != nil {
		return nil
	}
	bResp := readResponse(r)
	soap := gosoap.SoapMessage(bResp)
	b := soap.BodyBytes()
	if b == nil {
		return fmt.Errorf("body is nil")
	}
	err = xml.Unmarshal(b, resp)
	if err != nil {
		return err
	}
	return nil
}

func IsPtr(data interface{}) bool {
	return reflect.ValueOf(data).Kind() == reflect.Ptr
}

func CheckParam(req interface{}, resp interface{}) error {
	if IsPtr(req) || !IsPtr(resp) {
		return fmt.Errorf("param err, expect that req is not a ptr, resp is a ptr")
	}
	tResp := reflect.TypeOf(resp)
	if tResp.Elem().Kind() == reflect.Ptr {
		return fmt.Errorf("param err, expect that *resp is not a ptr")
	}
	tReqName := reflect.TypeOf(req).Name()
	logs.Infof(tReqName)
	node, ok := reflectMap[tReqName]
	if !ok {
		return fmt.Errorf("param type: %s is not in reflect map", tReqName)
	}
	if tReqName != node.Req.Name() {
		return fmt.Errorf("param type: %s is not in reflect map", tResp.Elem())
	}
	if tResp.Elem().Kind() != node.Resp.Kind() {
		return fmt.Errorf("param type: %s is not in reflect map", tResp.Elem())
	}
	return nil
}

func Init() error {
	reflectTypes := ReflectTypes{}
	reflectMap = make(map[string]*FuncNode)
	tt := reflect.TypeOf(reflectTypes)
	tv := reflect.ValueOf(reflectTypes)
	for i := 0; i < tt.NumField(); i++ {
		field := tt.Field(i)
		value := tv.Field(i)
		tag, ok := field.Tag.Lookup("method")
		if !ok {
			return fmt.Errorf("there's no tag function in fields: %v", field)
		}
		ttc := value.Type()
		r := FuncNode{
			Name: tag,
			Req:  nil,
			Resp: nil,
		}
		for j := 0; j < ttc.NumField(); j++ {
			f := ttc.Field(j)
			v := value.Field(j)
			_, ok = f.Tag.Lookup("req")
			if ok {
				r.Req = v.Type()
				continue
			}
			_, ok = f.Tag.Lookup("resp")
			if ok {
				r.Resp = v.Type()
				continue
			}
		}
		if r.Req == nil || r.Resp == nil {
			return fmt.Errorf("struct field error")
		}
		reflectMap[tag] = &r
	}
	return nil
}

func readResponse(resp *http.Response) []byte {
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	return b
}
