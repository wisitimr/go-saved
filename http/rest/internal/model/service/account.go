package service

import (
	"context"
	mAccount "saved/http/rest/internal/model/account"
	mRes "saved/http/rest/internal/model/response"
)

type AccountService interface {
	Count(ctx context.Context) (mRes.CountDto, error)
	FindAll(ctx context.Context, query map[string][]string) ([]mAccount.Account, error)
	FindById(ctx context.Context, id string) (mAccount.Account, error)
	Create(ctx context.Context, payload mAccount.Account) (mAccount.Account, error)
	Update(ctx context.Context, id string, payload mAccount.Account) (mAccount.Account, error)
}
