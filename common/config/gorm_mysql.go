package config

type Mysql struct {
	GeneralDB `yaml:",inline" mapstructure:",squash"`
}

// jdbc:mysql://localhost:3306/Chinese_Learning_DB
func (m *Mysql) Dsn() string {
	return m.Username + ":" + m.Password + "@tcp(" + m.Path + ":" + m.Port + ")/" + m.Dbname + "?" + m.Config
}

func (m *Mysql) GetLogMode() string {
	return m.LogMode
}
