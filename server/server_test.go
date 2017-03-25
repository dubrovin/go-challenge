package server

import (
	"testing"
	"time"
)

func TestNewServer(t *testing.T) {
	expectedTimeout := time.Millisecond * 500
	newServer := NewServer(expectedTimeout)

	if newServer.Client == nil {
		t.Error("Client equals to nil")
	}
}

func TestServerRun(t *testing.T) {
	expectedTimeout := time.Millisecond * 500
	newServer := NewServer(expectedTimeout)
	newServer.Run()
}