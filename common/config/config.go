package config

// todo：配置文件暂时没有

type Server struct {
	JWT    JWT             `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Zap    Zap             `mapstructure:"zap" json:"zap" yaml:"zap"`
	System System          `mapstructure:"system" json:"system" yaml:"system"`
	Mysql  Mysql           `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	DBList []SpecializedDB `mapstructure:"db-list" json:"db-list" yaml:"db-list"`
	Redis  Redis           `mapstructure:"redis" json:"redis" yaml:"redis"`
	Timer  Timer           `mapstructure:"timer" json:"timer" yaml:"timer"`
}
