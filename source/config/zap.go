package config

type Zap struct {
	Level string `mapstructure:"level" json:"level" yaml:"level"`
	Filename string `mapstructure:"filename" json:"filename" yaml:"filename"`
	MaxSize int `mapstructure:"maxsize" json:"maxsize" yaml:"maxsize"`
	MaxAge int `mapstructure:"maxage" json:"maxage" yaml:"maxage"`
	MaxBackups int `mapstructure:"maxbackups" json:"maxbackups" yaml:"maxbackups"`
}