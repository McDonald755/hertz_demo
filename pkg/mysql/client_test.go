package mysql

import (
	"demo2/pkg/viper"
	"github.com/cloudwego/hertz/pkg/common/test/assert"
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
