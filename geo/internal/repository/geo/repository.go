package geo

import (
	"context"
	"os"

	"geo/internal/model"

	"github.com/ekomobile/dadata/v2"
	"github.com/ekomobile/dadata/v2/api/suggest"
	"github.com/ekomobile/dadata/v2/client"
	"github.com/jackc/pgx/v4/pgxpool"
)

var _ Repository = (*repository)(nil)

type Repository interface {
	GetAddress(ctx context.Context, query string) ([]*model.Address, error)
	SetExtension(string) error
	FindQuery(ctx context.Context, query string) (int, error)
	GetAddressByHistory(ctx context.Context, historyId int) ([]*model.Address, error)
	CallExternalApi(ctx context.Context, endpoint string, params ...string) ([]*model.Address, error)
	SaveQuery(ctx context.Context, query string) (int, error)
	SaveAddress(ctx context.Context, historyId int, address *model.Address) (int, error)
	SaveHistoryAddress(ctx context.Context, history_id int) (int, error)
}

type repository struct {
	pool *pgxpool.Pool
	Api  *suggest.Api
}

func NewRepository(pool *pgxpool.Pool) *repository {
	creds := client.Credentials{
		ApiKeyValue:    os.Getenv("DADATA_API_KEY"),
		SecretKeyValue: os.Getenv("DADATA_SECRET_KEY"),
	}
	return &repository{
		Api:  dadata.NewSuggestApi(client.WithCredentialProvider(&creds)),
		pool: pool,
	}
}
