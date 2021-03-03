package impl

import (
	"github.com/andrii-minchekov/lets-go/app/usecases"
	cfg "github.com/andrii-minchekov/lets-go/domain"
)

type composedUseCases struct {
	uc.SnippetUseCase
	uc.UserUseCase
}

func NewComposedUseCases(config cfg.Config) uc.UseCases {
	db := Database{connect(config.DSN())}
	return composedUseCases{
		uc.NewSnippetUseCase(dbSnippetRepository{db}),
		uc.NewUserUseCase(dbUserRepository{db}),
	}
}
