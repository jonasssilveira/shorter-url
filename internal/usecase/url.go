package usecase

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	db "urlShorter/db/sqlc"
	"urlShorter/infra/repository"
	customerror "urlShorter/internal/handlers/errors"
)

type URL struct {
	query repository.URLRepository
}

func (u *URL) DeleteDeprecatedURL(ctx context.Context) error {
	return u.query.DeleteDeprecatedURL(ctx)
}

func (u *URL) UpdateURL(ctx context.Context, arg db.UpdateURLParams) (db.UpdateURLRow, error) {
	//TODO implement me
	panic("implement me")
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
				return db.CreateURLRow{}, customerror.NewDatabaseError(
					fmt.Sprintf("error to retrieve url, error %v", err.Error()),
					err.Error(),
					http.StatusInternalServerError,
				)
			}
			return createURL, err
		}
	}

	updateURL, err := u.query.UpdateURL(ctx, db.UpdateURLParams{
		UrlEncoded:     params.UrlEncoded,
		ExpirationDate: params.ExpirationDate,
	})

	if err != nil {
		return db.CreateURLRow{}, customerror.NewDatabaseError(
			fmt.Sprintf("error to retrieve url, error %v", err.Error()),
			err.Error(),
			http.StatusInternalServerError)
	}

	return db.CreateURLRow{
		UrlEncoded:     updateURL.UrlEncoded,
		ExpirationDate: updateURL.ExpirationDate,
	}, err
}

func (u *URL) DeleteURL(ctx context.Context, urlEncoded string) error {
	return u.query.DeleteURL(ctx, urlEncoded)
}

func (u *URL) GetURL(ctx context.Context, urlEncoded string) (string, error) {
	url, err := u.query.GetURL(ctx, urlEncoded)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", customerror.NewDatabaseError(
				fmt.Sprintf("not found %v", err.Error()),
				err.Error(),
				http.StatusNotFound)
		}
		return "", customerror.NewDatabaseError(
			fmt.Sprintf("error to retrieve data, error %v", err.Error()),
			err.Error(),
			http.StatusInternalServerError,
		)
	}
	return url, nil
}
