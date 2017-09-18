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
	_sessionHandle *tcliservice.TSessionHandle
}

func (p *Connection) client() *tcliservice.TCLIServiceClient{
	return p._client
}

func (p *Connection) sessionHandle() *tcliservice.TSessionHandle{
	return p._sessionHandle
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

	err = transport.Open()
	if err != nil {
		return nil, err
	}

	protocolVersion := tcliservice.TProtocolVersion_HIVE_CLI_SERVICE_PROTOCOL_V10

	openSessionReq := &tcliservice.TOpenSessionReq{
		ClientProtocol: protocolVersion,
	}

	response, err := client.OpenSession(openSessionReq)

	if err != nil {
		return nil, err
	}

	if response.SessionHandle == nil {
		return nil, fmt.Errorf("No session handle created after connection")
	}

	if response.ServerProtocolVersion != protocolVersion {
		return nil, fmt.Errorf("Unable to handle protocol version %s", response.ServerProtocolVersion)
	}

	return &Connection{_client: client, _sessionHandle: response.SessionHandle}, nil
}