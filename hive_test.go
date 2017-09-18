package gohive

import
(
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

func TestConnectWithParam(t *testing.T){
	conn, err := Connect(&ConnParams{host: "localhost", database: "default", auth: "NOSASL"})

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