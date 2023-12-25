package config

type Monorepo struct {
	TagPrefix string `yaml:"tag_prefix,omitempty" json:"tag_prefix,omitempty"`
	Dir       string `yaml:"dir,omitempty" json:"dir,omitempty"`
}
