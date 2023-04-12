package config

//存储类型 local本地存储  etcd存储 consul存储
type Storage struct {
	Local LocalStorage `mapstructure:"local" json:"local" yml:"local" yaml:"local"`
	Type string `mapstructure:"type" json:"type" yml:"type" yaml:"type"`
}

