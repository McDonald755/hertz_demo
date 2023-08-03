package viper

type Config struct {
	App   *App   `yaml:"App"`
	Mysql *Mysql `yaml:"Mysql"`
	Redis *Redis `yaml:"Redis"`
}

type App struct {
	HostPorts          string `yaml:"hostPorts"`          // 服务监听的地址和端口
	MaxRequestBodySize int    `yaml:"maxRequestBodySize"` // 最大的请求体大小
}

type Mysql struct {
	DriverName   string `yaml:"driverName"`   // 数据库类型
	Host         string `yaml:"host"`         //host
	Port         string `yaml:"port"`         //端口
	DatabaseName string `yaml:"databaseName"` //数据库名称
	Username     string `yaml:"username"`     //账户
	Password     string `yaml:"password"`     //密码
	Charset      string `yaml:"charset"`      //字符集
	Loc          string `yaml:"loc"`          //时区
}

type Redis struct {
	Addr               string `yaml:"Addr"`               // 服务所在地址和端口
	Password           string `yaml:"Password"`           // 密码
	Db                 int    `yaml:"Db"`                 // 数据库编号
	EncodeLockSecond   int    `yaml:"EncodeLockSecond"`   // 加密锁限流间隔
	DecodeLockSecond   int    `yaml:"DecodeLockSecond"`   // 解密锁限流间隔
	DriftLockSecond    int    `yaml:"DriftLockSecond"`    // Drift锁限流间隔
	CompressLockSecond int    `yaml:"CompressLockSecond"` // Compress锁限流间隔
	DriftLimit         int    `yaml:"DriftLimit"`         // 漂流信缓存数量限制
}
