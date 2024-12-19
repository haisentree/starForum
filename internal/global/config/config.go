package config

type Config struct {
	Mysql Mysql
	Cache Cache
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
