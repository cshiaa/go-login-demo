package config

type Jenkins struct {
	Url string `mapstructure:"url" json:"url" yaml:"url"`
	User string `mapstructure:"user" json:"user" yaml:"user"`
	Password string `mapstructure:"password" json:"password" yaml:"password"`
}

