package gohive

import (
	"fmt"

	"git.apache.org/thrift.git/lib/go/thrift"
	"github.com/weber09/gohive/tcliservice"
)

type ConnParams struct {
	Host     string
	Port     string
	Auth     string
	Database string
	Username string
	Password string
}

func Connect(params *ConnParams) (*Connection, error) {
	return newConnection(params)
}

type Connection struct {
	_client          *tcliservice.TCLIServiceClient
	_sessionHandle   *tcliservice.TSessionHandle
	_operationHandle *tcliservice.TOperationHandle
}

func (p *Connection) client() *tcliservice.TCLIServiceClient {
	return p._client
}

func (p *Connection) sessionHandle() *tcliservice.TSessionHandle {
	return p._sessionHandle
}

func (p *Connection) operationHandle() *tcliservice.TOperationHandle {
	return p._operationHandle
}

func newConnection(params *ConnParams) (*Connection, error) {

	host := "localhost"
	port := "10000"
	database := "default"
	username := ""
	password := ""
	if params != nil {
		if len(params.Host) > 0 {
			host = params.Host
		}

		if len(params.Port) > 0 {
			port = params.Port
		}

		if len(params.Database) > 0 {
			database = params.Database
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
		Username:       &username,
		Password:       &password,
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

	conn := &Connection{_client: client, _sessionHandle: response.SessionHandle}

	conn.Execute(fmt.Sprintf("USE %s", database))

	return conn, nil
}

func (p *Connection) Execute(query string) (string, error) {

	req := &tcliservice.TExecuteStatementReq{SessionHandle: p._sessionHandle, Statement: query}

	response, err := p._client.ExecuteStatement(req)

	code := "error"
	if response != nil {
		code = response.Status.GetStatusCode().String()
	}

	if err != nil {
		return code, err
	}

	p._operationHandle = response.OperationHandle

	return code, nil
}

func (p *Connection) FetchOne() (string, error) {
	req := &tcliservice.TFetchResultsReq{
		OperationHandle: p.operationHandle(),
		Orientation:     tcliservice.TFetchOrientation_FETCH_NEXT,
		MaxRows:         1,
	}

	response, err := p.client().FetchResults(req)
	if err != nil {
		return "", err
	}

	result := ""

	for i, c := range response.Results.Columns {
		if c.IsSetBinaryVal() {
			result += string(c.GetBinaryVal().GetValues()[0])
		} else if c.IsSetBoolVal() {
			result += fmt.Sprintf("%v", c.GetBoolVal().GetValues()[0])
		} else if c.IsSetByteVal() {
			result += fmt.Sprintf("%v", c.GetByteVal().GetValues()[0])
		} else if c.IsSetDoubleVal() {
			result += fmt.Sprintf("%v", c.GetDoubleVal().GetValues()[0])
		} else if c.IsSetI16Val() {
			result += fmt.Sprintf("%v", c.GetI16Val().GetValues()[0])
		} else if c.IsSetI32Val() {
			result += fmt.Sprintf("%v", c.GetI32Val().GetValues()[0])
		} else if c.IsSetI64Val() {
			result += fmt.Sprintf("%v", c.GetI64Val().GetValues()[0])
		} else if c.IsSetStringVal() {
			result += fmt.Sprintf("\"%s\"", c.GetStringVal().GetValues()[0])
		}
		if i < len(response.Results.Columns)-1 {
			result += ", "
		}
	}

	return result, nil
}

func (p *Connection) fetchWhile() {

}
