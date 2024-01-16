package usecase

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	db "urlShorter/db/sqlc"
	"urlShorter/repository"
)

type URL struct {
	query repository.URLRepository
}

func NewURL(query repository.URLRepository) URL {
	return URL{query: query}
}

func (u *URL) CreateURL(
	ctx context.Context,
	params db.CreateURLParams,
) (db.CreateURLRow, error) {
	_, err := u.query.GetURL(ctx, params.UrlEncoded)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			createURL, err := u.query.CreateURL(ctx, params)
			if err != nil {
				fmt.Printf("error to get url, error %v", err.Error())
				return db.CreateURLRow{}, err
			}
			return createURL, err
		}
	}

	updateURL, err := u.query.UpdateURL(ctx, db.UpdateURLParams{
		UrlEncoded:     params.UrlEncoded,
		ExpirationDate: params.ExpirationDate,
	})

	if err != nil {
		fmt.Printf("error to get url, error %v", err.Error())
		return db.CreateURLRow{}, err
	}

	return db.CreateURLRow{
		UrlEncoded:     updateURL.UrlEncoded,
		ExpirationDate: updateURL.ExpirationDate,
	}, err
}

func (u *URL) DeleteURL(ctx context.Context, urlEncoded string) error {
	return u.query.DeleteURL(ctx, urlEncoded)
}
