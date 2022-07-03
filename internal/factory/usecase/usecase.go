package usecase

import (
	"github.com/nakoding-community/test-practical-devsecops/internal/factory/repository"
	"github.com/nakoding-community/test-practical-devsecops/internal/usecase/auth"
	"github.com/nakoding-community/test-practical-devsecops/internal/usecase/user"
)

type Factory struct {
	Auth auth.Usecase
	User user.Usecase
}

func Init(r repository.Factory) Factory {
	f := Factory{}
	f.Auth = auth.NewUsecase(r)
	f.User = user.NewUsecase(r)

	return f
}
