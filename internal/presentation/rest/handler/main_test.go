package handler_test

import (
	"flag"
	"os"
	"testing"
)

var updateGolden = flag.Bool("update-golden", false, "update .golden files")

func TestMain(m *testing.M) {
	flag.Parse()
	code := m.Run()
	os.Exit(code)
}
