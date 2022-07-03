package factory

import (
	"github.com/nakoding-community/test-practical-devsecops/internal/abstraction"
	"github.com/nakoding-community/test-practical-devsecops/internal/factory/repository"
	"github.com/nakoding-community/test-practical-devsecops/internal/factory/usecase"
)

type Factory struct {
	Repository repository.Factory
	Usecase    usecase.Factory
	WsHub      *abstraction.WsHub
}

func Init() Factory {
	f := Factory{}
	// f.Repository = repository.Init()
	// f.Usecase = usecase.Init(f.Repository)
	f.WsHub = abstraction.NewWsHub()

	return f
}
