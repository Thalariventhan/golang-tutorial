package helloworld

import "testing"

func TestHelloWorld(t *testing.T) {
	Hello()

	t.Log("HelloWorldTest passed")
}