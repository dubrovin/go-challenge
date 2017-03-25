package server

import (
	"testing"
	"time"
	"github.com/stretchr/testify/require"
	"github.com/dubrovin/go-challnge/testserver"
	"net/http"
	"io/ioutil"
	"encoding/json"
)

func TestNewServer(t *testing.T) {
	expectedTimeout := time.Millisecond * 500
	addr := ":8080"
	newServer := NewServer(addr, expectedTimeout)

	require.NotNil(t, newServer.Client)
	require.Equal(t, addr, newServer.ListenAddr)
}

func init() {
	go testserver.Run()
	expectedTimeout := time.Millisecond * 500
	addr := ":8080"
	newServer := NewServer(addr, expectedTimeout)
	go newServer.Run()

}

func TestServerRun(t *testing.T) {
	resp, err := http.Get("http://127.0.0.1:8080/numbers?u=http://127.0.0.1:8090/primes&u=http://127.0.0.1:8090/fibo")
	require.Nil(t, err)
	require.NotNil(t, resp)
	var num Number
	body, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &num)
	require.Nil(t, err)
	require.Equal(t, []int{1, 2, 3, 5, 7, 8, 11, 13, 21}, num.Numbers)
}