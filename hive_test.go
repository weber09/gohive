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

func TestCreateTable(t *testing.T) {
	conn, err := Connect(nil)

	if err != nil {
		t.Error(err)
	}

	status, err := conn.Execute("create table test(id string, name string)")

	if err != nil {
		t.Errorf("error executing command %s", err)
	}

	if conn._operationHandle == nil {
		t.Error("Error receiveing operation handle")
	}

	t.Logf("status exec %s", status)
}

func TestExecuteInsertionCommand(t *testing.T) {
	conn, err := Connect(nil)

	if err != nil {
		t.Errorf("Error connecting [%s]", err)
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

func TestExecuteQuery(t *testing.T) {
	conn, err := Connect(nil)

	if err != nil {
		t.Errorf("Error connecting [%s]", err)
		t.Error(err)
	}

	_, err = conn.Execute("select * from test")

	if err != nil {
		t.Errorf("error executing query %s", err)
	}

	if conn._operationHandle == nil {
		t.Error("Error receiveing operation handle")
	}

	fetch, err := conn.FetchOne()

	if err != nil {
		t.Error(err)
	}

	t.Logf("fetchOne = [%s]", fetch)
}

func TestExecuteComplete(t *testing.T) {
	conn, err := Connect(nil)

	if err != nil {
		t.Errorf("Error connecting [%s]", err)
		t.Error(err)
	}

	_, err = conn.Execute("create table test (id int, name string)")

	if err != nil {
		t.Errorf("Error crating table: [%s]", err)
	}

	_, err = conn.Execute("insert into test (id, name) values (1, \"Mark\"), (2, \"John\")")

	if err != nil {
		t.Errorf("Error inserting data: [%s]", err)
	}

	_, err = conn.Execute("select * from test")

	if err != nil {
		t.Errorf("Error querying the table: [%s]", err)
	}

	fetch, err := conn.FetchOne()

	if err != nil {
		t.Error(err)
	}

	t.Logf("fetchOne = [%v]", fetch)

	_, err = conn.Execute("drop table test")

	if err != nil {
		t.Errorf("Error dropping table: [%s]", err)
	}
}
