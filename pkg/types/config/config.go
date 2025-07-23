package config

type MySQLConfig struct {
	DataSource      string `yaml:"DataSource"`
	MaxIdleConns    int    `yaml:"MaxIdleConns"`
	MaxOpenConns    int    `yaml:"MaxOpenConns"`
	ConnMaxLifetime int    `yaml:"ConnMaxLifetime"`
}

type RedisConfig struct {
	Host string `yaml:"Host"`
	Type string `yaml:"Type"`
	Pass string `yaml:"Pass"`
	DB   int    `yaml:"DB"`
}

type MonitorConfig struct {
	Enabled bool   `yaml:"Enabled"`
	Port    int    `yaml:"Port"`
	Path    string `yaml:"Path"`
	Name    string `yaml:"Name"`
	Pass    string `yaml:"Pass"`
}

type ThirdPlatform struct {
	Lazada    map[string]string `yaml:"Lazada"`
	Shopee    map[string]string `yaml:"Shopee"`
	ShopeeErp map[string]string `yaml:"ShopeeErp"`
	TikTok    map[string]string `yaml:"TikTok"`
	TikTokUs  map[string]string `yaml:"TikTokUs"`
}

type AppConfig struct {
	Name                string         `yaml:"Name"`
	Mode                string         `yaml:"Mode"`
	Log                 map[string]any `yaml:"Log"`
	Redis               RedisConfig    `yaml:"Redis"`
	MySQL               MySQLConfig    `yaml:"MySQL"`
	MySQLThird          MySQLConfig    `yaml:"MySQLThird"`
	MySQLThirdV2        MySQLConfig    `yaml:"MySQLThirdV2"`
	Monitor             MonitorConfig  `yaml:"Monitor"`
	ThirdPlatformConfig ThirdPlatform  `yaml:"ThirdPlatformConfig"`
}
