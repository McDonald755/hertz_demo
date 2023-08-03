package mysql

import (
	"context"
	"demo2/pkg/viper"
	"fmt"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"net/url"
)

var (
	DB     *gorm.DB
	config *viper.Mysql
)

func InitMysql() {
	// 配置初始化
	config = viper.Conf.Mysql
	initMysql(context.Background())
}

func initMysql(ctx context.Context) {
	//拼接数据库链接信息
	sqlStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true&loc=%s",
		config.Username,
		config.Password,
		config.Host,
		config.Port,
		config.DatabaseName,
		config.Charset,
		url.QueryEscape(config.Loc),
	)

	db, err := gorm.Open(mysql.Open(sqlStr), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	if err != nil {
		hlog.CtxErrorf(ctx, "[Viper] ReadInConfig failed, err: %v", err)
	}
	//监听增删改的sql日志
	db.Callback().Update().Register("updateSql", callback)
	db.Callback().Delete().Register("deleteSql", callback)
	db.Callback().Create().Register("createSql", callback)

	//暴露数据库连接
	DB = db
}

// gorm 打印sql回调方法
func callback(option *gorm.DB) {
	hlog.Debug(option.Statement.SQL.String())
}
