package mysql

import (
	"github.com/cloudwego/hertz/pkg/common/test/assert"
	"hertz_demo/pkg/viper"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	_ = os.Setenv("RUN_ENV", "TEST")
	viper.InitViper()
	m.Run()
}

func TestInitMysql(t *testing.T) {
	InitMysql()
	assert.NotNil(t, DB)
}
