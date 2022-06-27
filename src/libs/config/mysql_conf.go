package config

const (
	DefaultMysqlDSN = "admin:admin@tcp(127.0.0.1:3306)/gmangos"
)

func (c *MySQLConf) ParseConfig() (err error) {
	if c.DSN == "" {
		c.DSN = DefaultMysqlDSN
	}

	return
}
