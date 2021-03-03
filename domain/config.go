package cfg

type Config interface {
	Addr() string
	DSN() string
	HTMLDir() string
	StaticDir() string
	Secret() string
}
