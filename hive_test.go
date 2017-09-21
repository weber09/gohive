package gohive

import (
	"testing"
)

func TestConnect(t *testing.T) {
	conn, err := Connect(nil)

	if err != nil {
		t.Error(err)
	}

	if conn == nil {
		t.Error("conn not created")
	}

}

func TestConnectWithParam(t *testing.T) {
	conn, err := Connect(&ConnParams{Host: "localhost", Database: "default", Auth: "NOSASL"})

	if err != nil {
		t.Error(err)
	}

	if conn == nil {
		t.Error("conn not created with params")
	}

	if conn.client() == nil {
		t.Error("client not created by conn")
	}
}

func TestConnectGetSessionHandle(t *testing.T) {
	conn, err := Connect(nil)

	if err != nil {
		t.Error(err)
	}

	if conn.sessionHandle() == nil {
		t.Error("Session handle not created after connection")
	}
}

func TestExecuteQuery(t *testing.T) {
	conn, err := Connect(nil)

	if err != nil {
		t.Error(err)
	}

	status, err := conn.Execute("insert into test (id, name) values (4, 'gotest')")

	if err != nil {
		t.Errorf("error executing query %s", err)
	}

	if conn._operationHandle == nil {
		t.Error("Error receiveing operation handle")
	}

	t.Logf("status exec %s", status)
}
