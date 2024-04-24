package config

type Config struct {
	Default DefaultConfig `yaml:"default"`
	Mysql   MysqlConfig   `yaml:"mysql"`
}

type DefaultConfig struct {
	// debug mode
	Mode   string `yaml:"mode"`
	Listen int    `yaml:"listen"`

	// 自动创建表结构
	AutoMigrate bool `yaml:"autoMigrate"`
}

type MysqlConfig struct {
	Host     string `yaml:"host"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Port     int    `yaml:"port"`
	Name     string `yaml:"name"`
}
