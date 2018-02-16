package gohive

import (
	"fmt"
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

	if conn.getClient() == nil {
		t.Error("client not created by conn")
	}
}

func TestConnectGetSessionHandle(t *testing.T) {
	conn, err := Connect(nil)

	if err != nil {
		t.Error(err)
	}

	if conn.getSessionHandle() == nil {
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

	if conn.operationHandle == nil {
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

	if conn.operationHandle == nil {
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

	if conn.operationHandle == nil {
		t.Error("Error receiveing operation handle")
	}

	fetch, err := conn.FetchOne()

	if err != nil {
		t.Error(err)
	}

	t.Logf("fetchOne = [%s]", fetch)
}

func TestExecuteComplete_FetchOne(t *testing.T) {
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

func TestExecuteComplete_FetchMany(t *testing.T) {
	conn, err := Connect(nil)

	if err != nil {
		t.Errorf("Error connecting [%s]", err)
		t.Error(err)
	}

	_, err = conn.Execute("create table test (id int, name string)")

	if err != nil {
		t.Errorf("Error crating table: [%s]", err)
	}

	_, err = conn.Execute("insert into test (id, name) values (1, \"Mark\"), (2, \"John\"), (3, \"Paul\"), (4, \"Ruecker\"), (5, \"Danny\")")

	if err != nil {
		t.Errorf("Error inserting data: [%s]", err)
	}

	_, err = conn.Execute("select * from test")

	if err != nil {
		t.Errorf("Error querying the table: [%s]", err)
	}

	fetch, err := conn.FetchMany(3)

	if err != nil {
		t.Error(err)
	}

	if len(fetch) != 3 {
		t.Fatal("Should have returned 3 rows from Hive table")
	} else {
		t.Logf("fetchMany(10) = [%v]", fetch)
	}
	_, err = conn.Execute("drop table test")

	if err != nil {
		t.Errorf("Error dropping table: [%s]", err)
	}
}

func TestExecuteComplete_FetchMany_MoreThanExistingRows(t *testing.T) {
	conn, err := Connect(nil)

	if err != nil {
		t.Errorf("Error connecting [%s]", err)
		t.Error(err)
	}

	_, err = conn.Execute("create table test (id int, name string)")

	if err != nil {
		t.Errorf("Error crating table: [%s]", err)
	}

	_, err = conn.Execute("insert into test (id, name) values (1, \"Mark\"), (2, \"John\"), (3, \"Paul\"), (4, \"Ruecker\"), (5, \"Danny\")")

	if err != nil {
		t.Errorf("Error inserting data: [%s]", err)
	}

	_, err = conn.Execute("select * from test")

	if err != nil {
		t.Errorf("Error querying the table: [%s]", err)
	}

	fetch, err := conn.FetchMany(15)

	if err != nil {
		t.Error(err)
	}

	if len(fetch) != 5 {
		t.Fatal("Should have returned 5 rows from Hive table")
	} else {
		t.Logf("fetchMany(10) = [%v]", fetch)
	}
	_, err = conn.Execute("drop table test")

	if err != nil {
		t.Errorf("Error dropping table: [%s]", err)
	}
}

func TestExecuteComplete_FetchMany_Input0Size(t *testing.T) {
	conn, err := Connect(nil)

	if err != nil {
		t.Errorf("Error connecting [%s]", err)
		t.Error(err)
	}

	_, err = conn.Execute("create table test (id int, name string)")

	if err != nil {
		t.Errorf("Error crating table: [%s]", err)
	}

	_, err = conn.Execute("insert into test (id, name) values (1, \"Mark\"), (2, \"John\"), (3, \"Paul\"), (4, \"Ruecker\"), (5, \"Danny\")")

	if err != nil {
		t.Errorf("Error inserting data: [%s]", err)
	}

	_, err = conn.Execute("select * from test")

	if err != nil {
		t.Errorf("Error querying the table: [%s]", err)
	}

	fetch, err := conn.FetchMany(0)

	if err != nil {
		t.Error(err)
	}

	if len(fetch) != 5 {
		t.Fatal("Should have returned 5 rows from Hive table")
	} else {
		t.Logf("fetchMany(10) = [%v]", fetch)
	}
	_, err = conn.Execute("drop table test")

	if err != nil {
		t.Errorf("Error dropping table: [%s]", err)
	}
}

func TestExecuteComplete_FetchAll(t *testing.T) {
	conn, err := Connect(nil)

	if err != nil {
		t.Errorf("Error connecting [%s]", err)
		t.Error(err)
	}

	_, err = conn.Execute("create table test (id int, name string)")

	if err != nil {
		t.Errorf("Error crating table: [%s]", err)
	}

	_, err = conn.Execute("insert into test (id, name) values (1, \"Mark\"), (2, \"John\"), (3, \"Paul\"), (4, \"Ruecker\"), (5, \"Danny\")")

	if err != nil {
		t.Errorf("Error inserting data: [%s]", err)
	}

	_, err = conn.Execute("select * from test")

	if err != nil {
		t.Errorf("Error querying the table: [%s]", err)
	}

	fetch, err := conn.FetchAll()

	if err != nil {
		t.Error(err)
	}

	if len(fetch) != 5 {
		t.Fatal("Should have returned 5 rows from Hive table")
	} else {
		t.Logf("fetchMany(10) = [%v]", fetch)
	}
	_, err = conn.Execute("drop table test")

	if err != nil {
		t.Errorf("Error dropping table: [%s]", err)
	}
}

func TestExecuteComplete_FetchWithoutQuery(t *testing.T) {
	conn, err := Connect(nil)

	if err != nil {
		t.Errorf("Error connecting [%s]", err)
		t.Error(err)
	}

	_, err = conn.Execute("create table test (id int, name string)")

	if err != nil {
		t.Errorf("Error crating table: [%s]", err)
	}

	_, err = conn.Execute("insert into test (id, name) values (1, \"Mark\"), (2, \"John\"), (3, \"Paul\"), (4, \"Ruecker\"), (5, \"Danny\")")

	if err != nil {
		t.Errorf("Error inserting data: [%s]", err)
	}

	fetch, err := conn.FetchOne()

	if err != nil {
		t.Errorf("Error fetching data: [%s]", err)
	}

	if fetch != nil {
		t.Errorf("No data should be returned! No query were run to return data.\n[%s]", fetch)
	}

	_, err = conn.Execute("drop table test")

	if err != nil {
		t.Errorf("Error dropping table: [%s]", err)
	}
}

func TestExecuteComplete_FetchAll_5kData(t *testing.T) {
	conn, err := Connect(nil)

	if err != nil {
		t.Errorf("Error connecting [%s]", err)
		t.Error(err)
	}

	_, err = conn.Execute("drop table test")
	_, err = conn.Execute("create table test (id int, name string)")

	if err != nil {
		t.Errorf("Error crating table: [%s]", err)
	}

	queryCmd := "insert into test (id, name) values "
	for i := 1; i <= 5000; i++ {
		queryCmd += fmt.Sprintf("(%d, \"%d\"), ", i, i)
		if i%1000 == 0 {
			queryCmd = queryCmd[:len(queryCmd)-2]
			_, err := conn.Execute(queryCmd)
			if err != nil {
				t.Errorf("Error inserting data: [%s]", err)
			}
			queryCmd = "insert into test (id, name) values "
		}
	}

	_, err = conn.Execute("select * from test")

	if err != nil {
		t.Errorf("Error querying the table: [%s]", err)
	}

	fetch, err := conn.FetchAll()

	if err != nil {
		t.Error(err)
	}

	if len(fetch) != 5000 {
		t.Fatalf("Should have returned 5000 rows from Hive table\nReturned rows count: [%d]", len(fetch))
	} else {
		t.Logf("fetchAll() = [%v]", fetch)
	}
	_, err = conn.Execute("drop table test")

	if err != nil {
		t.Errorf("Error dropping table: [%s]", err)
	}
}
