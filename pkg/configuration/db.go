package configuration

type DbConfig struct {
	User, Password, Host, Port, Database string
	PoolMin, PoolMax                     int
}
