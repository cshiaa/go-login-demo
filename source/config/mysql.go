package config


type Mysql struct {
	Dbname string `mapstructure:"dbname" json:"dbname" yaml:"dbname"`
	Dirver string `mapstructure:"dirver" json:"dirver" yaml:"dirver"`
	Host string `mapstructure:"host" json:"host" yaml:"host"`
	Username string `mapstructure:"username" json:"username" yaml:"username"`
	Password string `mapstructure:"password" json:"password" yaml:"password"`
	Port string `mapstructure:"port" json:"port" yaml:"port"`
}
	// DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", db_user, db_password, db_host, db_port, db_name)

func (m *Mysql) Dsn() string {
	return m.Username + ":" + m.Password + "@tcp(" + m.Host + ":" + m.Port + ")/" + m.Dbname + "?charset=utf8&parseTime=True&loc=Local"
}
