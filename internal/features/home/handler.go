package home

import (
	"github.com/charlierz/go-htmx-starter/internal/features/item"
	"github.com/charlierz/go-htmx-starter/internal/template/pages"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	homeService  *Service
	itemsService *item.Service
}

func NewHandler(hs *Service, as *item.Service) *Handler {
	return &Handler{
		homeService:  hs,
		itemsService: as,
	}
}

func Register(e *echo.Echo, h *Handler) {
	home := e.Group("")
	home.GET("", h.Home)
}

func (h *Handler) Home(c echo.Context) error {
	items, err := h.itemsService.List()
	if err != nil {
		return err
	}
	component := pages.Home(items)
	return component.Render(c.Request().Context(), c.Response().Writer)
}
