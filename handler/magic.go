package handler

import (
	"net/http"
	"strconv"

	"golang-starter-pack/model"
	"golang-starter-pack/utils"

	"github.com/labstack/echo/v4"
)

func (h *Handler) GetCard(c echo.Context) error {
	name := c.Param("name")
	cd, err := h.magicStore.GetByName(name)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}
	if cd == nil {
		return c.JSON(http.StatusNotFound, utils.NotFound())
	}
	return c.JSON(http.StatusOK, newCardResponse(c, cd))
}

func (h *Handler) Cards(c echo.Context) error {
	offset, err := strconv.Atoi(c.QueryParam("offset"))
	if err != nil {
		offset = 0
	}
	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil {
		limit = 20
	}
	var cards []model.Card
	var count int

	cards, count, err = h.magicStore.List(offset, limit)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, nil)

	}
	return c.JSON(http.StatusOK, newCardListResponse(h.userStore, userIDFromToken(c), cards, count))
}

func (h *Handler) CreateCard(c echo.Context) error {
	var a model.Article
	req := &articleCreateRequest{}
	if err := req.bind(c, &a); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	a.AuthorID = userIDFromToken(c)
	err := h.articleStore.CreateArticle(&a)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}

	return c.JSON(http.StatusCreated, newArticleResponse(c, &a))
}

func (h *Handler) DeleteCard(c echo.Context) error {
	slug := c.Param("slug")
	a, err := h.articleStore.GetUserArticleBySlug(userIDFromToken(c), slug)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}
	if a == nil {
		return c.JSON(http.StatusNotFound, utils.NotFound())
	}
	err = h.articleStore.DeleteArticle(a)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}
	return c.JSON(http.StatusOK, map[string]interface{}{"result": "ok"})
}
