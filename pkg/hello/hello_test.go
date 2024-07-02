package hello

import "testing"

func TestHello(t *testing.T) {
	msg := GetMessage()
	if msg != "Hello, World!" {
		t.Errorf("expected \"Hello, World!\", but got %v", msg)
	}
}
