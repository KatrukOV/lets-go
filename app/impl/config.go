package impl

import (
	"flag"
	"github.com/andrii-minchekov/lets-go/domain"
)

type flagConfig struct {
	addr      *string
	dsn       *string
	secret    *string
	htmlDir   *string
	staticDir *string
}

func NewFlagConfig() cfg.Config {
	config := flagConfig{
		addr:      flag.String("addr", ":4000", "HTTP network address"),
		dsn:       flag.String("dsn", "user=postgres dbname=snippetbox password=postgres sslmode=disable", "Postgres Datasource Name"),
		secret:    flag.String("secret", "s6Nd%+pPbnzHbS*+9Pk8qGWhTzbpa@ge", "Secret key"),
		htmlDir:   flag.String("html-dir", "./ui/html", "Path to HTML templates"),
		staticDir: flag.String("static-dir", "./ui/static", "Path to static assets"),
	}
	flag.Parse()
	return config
}

func (receiver flagConfig) DSN() string {
	return *receiver.dsn
}

func (receiver flagConfig) HTMLDir() string {
	return *receiver.htmlDir
}

func (receiver flagConfig) StaticDir() string {
	return *receiver.staticDir
}

func (receiver flagConfig) Secret() string {
	return *receiver.secret
}

func (receiver flagConfig) Addr() string {
	return *receiver.addr
}
