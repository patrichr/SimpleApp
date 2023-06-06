package repo

import (
	"context"

	"github.com/patrichr/SimpleApp/server/model"
)

func (repo *Repo) Register(ctx context.Context, req model.Register) (num int, err error) {
	resOrm := repo.DB.Create(req)
	return 0, resOrm.Error
}
