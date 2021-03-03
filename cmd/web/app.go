package main

import (
	"github.com/alexedwards/scs"
	"github.com/andrii-minchekov/lets-go/app/usecases"
	"github.com/andrii-minchekov/lets-go/domain"
)

type App struct {
	Config   cfg.Config
	Cases    uc.UseCases
	Sessions *scs.Manager
	//TLSCert   string
	//TLSKey    string
}
