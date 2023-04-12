package config

// KubernetesConfig Path, Default ${HOME}/.kube/config
type KubernetesConfig struct {
	Path string `mapstructure:"path" json:"path" yaml:"path"`
}