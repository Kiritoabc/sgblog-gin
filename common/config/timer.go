package config

type Timer struct {
	Spec        string   `mapstructure:"spec" json:"spec" yaml:"spec"` // CRON表达式
	Detail      []Detail `mapstructure:"detail" json:"detail" yaml:"detail"`
	Start       bool     `mapstructure:"start" json:"start" yaml:"start"`                      // 是否启用
	WithSeconds bool     `mapstructure:"with_seconds" json:"with_seconds" yaml:"with_seconds"` // 是否精确到秒
}

type Detail struct{}
