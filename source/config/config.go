package config

type Server struct {

	Zap     Zap     `mapstructure:"zap" json:"zap" yaml:"zap"`
	// gorm
	Mysql  Mysql           `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	//Jenkins
	Jenkins Jenkins	`mapstructure:"jenkins" json:"jenkins" yaml:"jenkins"`
}


