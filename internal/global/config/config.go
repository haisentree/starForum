package config

type Config struct {
	Mysql  Mysql
	Cache  Cache
	Email  Email
	Server Server
}

// debug模式下，验证码随机都成功
type Server struct {
	Debug bool
}

type Mysql struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	DBName   string `yaml:"dbname"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Charset  string `yaml:"charset"`
}

type Cache struct {
	CacheTime  int   `yaml:"cacheTime"`
	CleanTime  int   `yaml:"cleanTime"`
	ExpireTime int64 `yaml:"expireTime"`
}

type Email struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Sender   string `yaml:"sender"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	TLS      bool   `yaml:"tls"`
}
