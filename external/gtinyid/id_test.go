package gtinyid

import (
	"log"
	"testing"
)

func TestNewIdGenerator(t *testing.T) {
	log.Print(NewIdGenerator("test", "test", "http://127.0.0.1:8080").Next())
}
