package logx

import (
	"log"
	"os"
	"testing"
)

func TestLogger_JSONFormat(t *testing.T) {
	l := New(os.Stdout, "logx: ", log.LstdFlags)
	l.Println("print...")
	l.WithCallers().Json("with", "callers...")
	l.SetCallersLevel(2).WithCallers().Jsonf("%s %s", "with", "callers...")
	l.WithFields(Fields{
		"v": "v0.0.0",
	}).Json("with fields")
}
