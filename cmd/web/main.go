package main

import (
	"github.com/alexedwards/scs"
	"github.com/andrii-minchekov/lets-go/app/impl"
	_ "github.com/lib/pq"
	"log"
	"time"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	config := impl.NewFlagConfig()

	// Initialize the Session Store
	sessionManager := scs.NewCookieManager(config.Secret())
	sessionManager.Lifetime(12 * time.Hour)
	sessionManager.Persist(true)

	app := &App{
		Config:   config,
		Cases:    impl.NewComposedUseCases(config),
		Sessions: sessionManager,
	}
	app.RunServer()
}
