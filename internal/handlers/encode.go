package handlers

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"net/http"
	db "urlShorter/db/sqlc"
	"urlShorter/infra/repository"
	"urlShorter/internal/domain/encode"
	"urlShorter/internal/domain/entity"
	"urlShorter/internal/handlers/errors"
)

type EncodeURL struct {
	urlRepository repository.URLRepository
	encoder       encode.URLEncoder
}

func NewEncodeURL(urlRepository repository.URLRepository, encoder encode.URLEncoder) EncodeURL {
	return EncodeURL{
		urlRepository: urlRepository,
		encoder:       encoder,
	}
}

func (e *EncodeURL) Create(ctx *fiber.Ctx) error {
	body := ctx.Body()
	var input entity.URL
	if err := json.Unmarshal(body, &input); err != nil {
		return err
	}

	encodedLink := e.encoder.Encode([]byte(input.FullURL))

	url, err := e.urlRepository.CreateURL(ctx.Context(), db.CreateURLParams{
		UrlEncoded:     encodedLink,
		UrlOriginal:    input.FullURL,
		ExpirationDate: input.ExpirationDate,
	})
	if err != nil {
		customError := err.(*errors.DatabaseError)
		ctx.Status(customError.StatusCode)
		return customError
	}
	ctx.Status(http.StatusOK)
	return ctx.JSON(url)
}

func (e *EncodeURL) Delete(ctx *fiber.Ctx) error {
	urlEncoded := ctx.Params("id")
	err := e.urlRepository.DeleteURL(ctx.Context(), urlEncoded)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		return ctx.JSON(err)
	}
	ctx.Status(http.StatusOK)
	return err
}

func (e *EncodeURL) Get(ctx *fiber.Ctx) error {
	urlEncoded := ctx.Params("id")
	encodedURL, err := e.urlRepository.GetURL(ctx.Context(), urlEncoded)

	if err != nil {
		customError := err.(*errors.DatabaseError)
		ctx.Status(customError.StatusCode)
		return ctx.JSON(err)
	}

	err = ctx.Redirect(encodedURL, http.StatusMovedPermanently)

	return err
}
