package infrastructure

import (
	"github.com/imroc/req"
)

//Request .
type Request interface {
	Get(string, []byte) (*req.Resp, error)
	Post(string, []byte) (*req.Resp, error)
	Put(string, []byte) (*req.Resp, error)
	Delete(string) (*req.Resp, error)
}
