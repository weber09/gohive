package gohive

import (
	"github.com/weber09/gohive/tcliservice"
	"git.apache.org/thrift.git/lib/go/thrift"
	"fmt"
)

type ConnParams struct {
	host string
	port string
	auth string
	database string
	username string
	password string
}


func Connect(params *ConnParams) (*Connection, error) {
	return newConnection(params)
}

type Connection struct {
	_client *tcliservice.TCLIServiceClient
}

func (p *Connection) client() *tcliservice.TCLIServiceClient{
	return p._client
}

func newConnection(params *ConnParams)  (*Connection, error) {

	host := "localhost"
	port := "10000"
	if params != nil {
		if len(params.host) > 0 {
			host = params.host
		}

		if len(params.port) > 0 {
			port = params.port
		}
	}

	transport, err := thrift.NewTSocket(fmt.Sprintf("%s:%s", host, port))

	if err != nil {
		return nil, err
	}

	protocol := thrift.NewTBinaryProtocolFactoryDefault()

	client := tcliservice.NewTCLIServiceClientFactory(transport, protocol)

	return &Connection{_client: client}, nil
}