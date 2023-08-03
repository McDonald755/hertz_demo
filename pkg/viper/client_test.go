package viper

import (
	"github.com/cloudwego/hertz/pkg/common/test/assert"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	_ = os.Setenv("RUN_ENV", "TEST")
	m.Run()
}

func TestInitViper(t *testing.T) {
	InitViper()
	assert.NotNil(t, Conf)
}
