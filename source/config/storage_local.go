package config

type LocalStorage struct {
	FilePath string `mapstructure:"filepath" json:"filepath" yml:"filepath" yaml:"filepath"`
}
