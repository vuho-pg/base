package config

type MySQL struct {
	Host     string `mapstructure:"MYSQL_HOST"`
	Port     int    `mapstructure:"MYSQL_PORT"`
	Username string `mapstructure:"MYSQL_USERNAME"`
	Password string `mapstructure:"MYSQL_PASSWORD"`
	Database string `mapstructure:"MYSQL_DATABASE"`
}

func ParseMySQL() (MySQL, error) {
	data := MySQL{}
	if err := Parse(&data); err != nil {
		return data, err
	}
	return data, nil
}
