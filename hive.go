package gohive

import (
	"fmt"
	"time"

	"git.apache.org/thrift.git/lib/go/thrift"
	"github.com/weber09/gohive/tcliservice"
)

const (
	//HiveNone status when no query/command executed
	HiveNone = "None"
	//HiveRunning status when a query/command was executed
	HiveRunning = "Running"
	//HiveFinished status when a query/command has finished its execution
	HiveFinished = "Finished"
)

//ConnParams holds informations for connection to Hive instance
type ConnParams struct {
	Host     string
	Port     string
	Auth     string
	Database string
	Username string
	Password string
}

//Connect calls the newConnection function
func Connect(params *ConnParams) (*Connection, error) {
	return newConnection(params)
}

//Connection is the struct holding the thrift generated structs to connect and handle the session with Hive
type Connection struct {
	client          *tcliservice.TCLIServiceClient
	sessionHandle   *tcliservice.TSessionHandle
	operationHandle *tcliservice.TOperationHandle
	data            [][]interface{}
	status          string
	transport       *thrift.TTransport
}

func (p *Connection) getClient() *tcliservice.TCLIServiceClient {
	return p.client
}

func (p *Connection) getSessionHandle() *tcliservice.TSessionHandle {
	return p.sessionHandle
}

func (p *Connection) getOperationHandle() *tcliservice.TOperationHandle {
	return p.operationHandle
}

func (p *Connection) getTransport() *thrift.TTransport {
	return p.transport
}

func newConnection(params *ConnParams) (*Connection, error) {

	host := "localhost"
	port := "10000"
	database := "default"
	username := ""
	password := ""
	auth := "NONE"
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

		if len(params.Auth) > 0 {
			auth = params.Auth
		}
	}

	socket, err := thrift.NewTSocket(fmt.Sprintf("%s:%s", host, port))

	if err != nil {
		return nil, err
	}

	var transport thrift.TTransport

	switch auth {
	case "NOSASL":
		transport, err = thrift.NewTBufferedTransportFactory(4096).GetTransport(socket)
		if err != nil {
			return nil, err
		}
		break
	case "LDAP":
	case "KERBEROS":
	case "NONE":
	case "CUSTOM":
		fmt.Println("SASL!!")
		break
	default:
		return nil, fmt.Errorf("Only NONE, NOSASL, LDAP, KERBEROS, CUSTOM authentication are supported. Got [%s]", auth)
	}

	//transport = socket

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

	conn := &Connection{client: client, sessionHandle: response.SessionHandle}

	conn.Execute(fmt.Sprintf("USE %s", database))

	conn.resetStatus()

	conn.transport = &transport

	return conn, nil
}

func (p *Connection) resetStatus() {
	p.status = HiveNone
}

func (p *Connection) close() error {
	req := &tcliservice.TCloseSessionReq{SessionHandle: p.getSessionHandle()}
	response, err := p.getClient().CloseSession(req)
	if err != nil {
		return err
	}

	if response.Status.StatusCode != tcliservice.TStatusCode_SUCCESS_STATUS {
		return fmt.Errorf("Error status closing session with hive: [%s]", *response.Status.ErrorMessage)
	}

	err = (*p.getTransport()).Close()
	if err != nil {
		return err
	}

	return nil
}

//Execute executes a query or command to the Hive connected instance
func (p *Connection) Execute(query string) (string, error) {

	req := &tcliservice.TExecuteStatementReq{SessionHandle: p.sessionHandle, Statement: query}

	response, err := p.client.ExecuteStatement(req)

	status := ""

	if response.IsSetStatus() {

		statusResp := response.GetStatus()

		errorCode := ""
		errorMessage := ""
		infoMessages := []string{}
		sqlState := ""

		if statusResp.IsSetErrorCode() {
			errorCode = fmt.Sprint(statusResp.GetErrorCode())
		}

		if statusResp.IsSetErrorMessage() {
			errorMessage = statusResp.GetErrorMessage()
		}

		if statusResp.IsSetInfoMessages() {
			infoMessages = statusResp.GetInfoMessages()
		}

		if statusResp.IsSetSqlState() {
			sqlState = statusResp.GetSqlState()
		}

		statusCode := statusResp.GetStatusCode()

		status = fmt.Sprintf("Status code: %s\nSql state:%s\nInfo messages: %v\nError code: %s\nError Message: %s\n", statusCode, sqlState, infoMessages, errorCode, errorMessage)
	}

	if err != nil {
		return status, err
	}

	p.operationHandle = response.OperationHandle

	p.status = HiveRunning

	return status, nil
}

//FetchOne fetches one row of data from Hive
func (p *Connection) FetchOne() ([]interface{}, error) {
	if p.status == HiveNone {
		return nil, fmt.Errorf("No query were run to fetch data")
	}

	for {
		if p.status == HiveFinished {
			break
		}

		p.fetchData()

		time.Sleep(1)
	}

	if len(p.data) == 0 {
		return nil, nil
	}

	row := p.data[0]
	p.data = p.data[1:]

	return row, nil
}

//FetchMany fetches a number of rows defined in the size parameter (Use 0 to return the default 1000 rows)
func (p *Connection) FetchMany(size int) ([][]interface{}, error) {

	if size == 0 {
		size = 1000
	}

	rows := make([][]interface{}, 0, size)
	for i := 0; i < size; i++ {
		row, err := p.FetchOne()
		if err != nil {
			return nil, err
		}
		if row == nil {
			break
		}
		rows = append(rows, row)
	}

	return rows, nil
}

//FetchAll fetches all rows of a Hive table
func (p *Connection) FetchAll() ([][]interface{}, error) {
	rows := make([][]interface{}, 0)
	for {
		row, err := p.FetchOne()
		if err != nil {
			return nil, err
		}
		if row == nil {
			break
		}
		rows = append(rows, row)
	}
	return rows, nil
}

func (p *Connection) fetchData() error {
	req := &tcliservice.TFetchResultsReq{
		OperationHandle: p.getOperationHandle(),
		Orientation:     tcliservice.TFetchOrientation_FETCH_NEXT,
		MaxRows:         1000,
	}

	response, err := p.getClient().FetchResults(req)
	if err != nil {
		return err
	}

	if response.Results == nil {
		p.status = HiveFinished
		return nil
	}

	rows := mountResults(response.Results.GetColumns())

	p.data = append(p.data, rows...)

	if len(rows) == 0 {
		p.status = HiveFinished
	}

	return nil
}

func mountResults(columns []*tcliservice.TColumn) [][]interface{} {
	rows := make([][]interface{}, 0)
	nrRows := 0
	column := make([]interface{}, 0)
	for _, c := range columns {
		if c.IsSetBinaryVal() {
			column = c.GetBinaryVal().GetInterfaceArray()
			rows, nrRows = joinColumnsInRows(column, &rows, nrRows)
		} else if c.IsSetBoolVal() {
			column = c.GetBoolVal().GetInterfaceArray()
			rows, nrRows = joinColumnsInRows(column, &rows, nrRows)
		} else if c.IsSetByteVal() {
			column = c.GetByteVal().GetInterfaceArray()
			rows, nrRows = joinColumnsInRows(column, &rows, nrRows)
		} else if c.IsSetDoubleVal() {
			column = c.GetDoubleVal().GetInterfaceArray()
			rows, nrRows = joinColumnsInRows(column, &rows, nrRows)
		} else if c.IsSetI16Val() {
			column = c.GetI16Val().GetInterfaceArray()
			rows, nrRows = joinColumnsInRows(column, &rows, nrRows)
		} else if c.IsSetI32Val() {
			column = c.GetI32Val().GetInterfaceArray()
			rows, nrRows = joinColumnsInRows(column, &rows, nrRows)
		} else if c.IsSetI64Val() {
			column = c.GetI64Val().GetInterfaceArray()
			rows, nrRows = joinColumnsInRows(column, &rows, nrRows)
		} else if c.IsSetStringVal() {
			column = c.GetStringVal().GetInterfaceArray()
			rows, nrRows = joinColumnsInRows(column, &rows, nrRows)
		}
	}

	return rows
}

func joinColumnsInRows(column []interface{}, rows *[][]interface{}, nrRows int) ([][]interface{}, int) {
	wRows := *rows

	if nrRows == 0 {
		nrRows = len(column)
	}

	for j := 0; j < nrRows; j++ {
		row := make([]interface{}, 0)
		if len(wRows) <= j {
			row = append(row, column[j])
			wRows = append(wRows, row)
		} else {
			wRows[j] = append(wRows[j], column[j])
		}
	}

	return wRows, nrRows
}
