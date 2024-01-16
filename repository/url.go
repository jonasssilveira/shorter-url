package repository

import (
	"context"
	db "urlShorter/db/sqlc"
)

type URLRepository interface {
	CreateURL(ctx context.Context, arg db.CreateURLParams) (db.CreateURLRow, error)
	DeleteDeprecatedURL(ctx context.Context) error
	DeleteURL(ctx context.Context, urlEncoded string) error
	GetURL(ctx context.Context, urlEncoded string) (db.Url, error)
	ListURL(ctx context.Context, arg db.ListURLParams) ([]db.Url, error)
	UpdateURL(ctx context.Context, arg db.UpdateURLParams) (db.UpdateURLRow, error)
}
