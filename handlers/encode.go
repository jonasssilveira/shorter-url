package handlers

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"time"
	db "urlShorter/db/sqlc"
	"urlShorter/domain/encode"
	"urlShorter/domain/entity"
	"urlShorter/usecase"
)

type EncodeURL struct {
	useCaseURL usecase.URL
	encoder    encode.URL
}

func NewEncodeURL(useCaseURL usecase.URL, encoder encode.URL) EncodeURL {
	return EncodeURL{
		useCaseURL: useCaseURL,
		encoder:    encoder,
	}
}

func (e *EncodeURL) Create(c *fiber.Ctx) error {
	body := c.Body()
	var input entity.URL
	if err := json.Unmarshal(body, &input); err != nil {
		return err
	}

	encodedLink := e.encoder.Encode(input.FullURL)

	url, err := e.useCaseURL.CreateURL(c.Context(), db.CreateURLParams{
		UrlEncoded:     encodedLink,
		UrlOriginal:    input.FullURL,
		ExpirationDate: time.Now(),
	})

	if err != nil {
		return err
	}
	return c.JSON(url)
}

func (e *EncodeURL) Delete(ctx *fiber.Ctx) error {
	urlEncoded := ctx.Params("id")
	err := e.useCaseURL.DeleteURL(ctx.Context(), urlEncoded)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		return ctx.JSON(err)
	}
	ctx.Status(http.StatusOK)
	return err
}
