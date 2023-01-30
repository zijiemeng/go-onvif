package common

type Client interface {
	CallMethodUnmarshal(endpoint string, method interface{}, result interface{}) *Fault
}
