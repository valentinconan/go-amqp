package utils

import (
	"testing"
	"os"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
    os.Setenv("VALUE","value")
    os.Exit(m.Run())
}

func TestGetenvDefault(t *testing.T){
    res := Getenv("","default")
    assert.Equal(t, "default", res)
}
func TestGetenv(t *testing.T){
    res := Getenv("VALUE","default")
    assert.Equal(t, "value", res)
}