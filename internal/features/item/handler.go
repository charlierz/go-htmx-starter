package item

import (
	"github.com/charlierz/go-htmx-starter/internal/model"
	"github.com/charlierz/go-htmx-starter/internal/template/pages"
	"github.com/charlierz/go-htmx-starter/internal/template/partials"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	service *Service
}

func NewHandler(s *Service) *Handler {
	return &Handler{
		service: s,
	}
}

func Register(e *echo.Echo, h *Handler) {
	items := e.Group("/items")
	items.GET("", h.List)
	items.POST("", h.Create)
	items.DELETE("/:itemId", h.Delete)
}

func (h *Handler) List(c echo.Context) error {
	items, err := h.service.List()
	if err != nil {
		return err
	}
	component := pages.ItemListPage(items)
	return component.Render(c.Request().Context(), c.Response().Writer)
}

func (h *Handler) Create(c echo.Context) error {
	newItem := model.NewItem{
		Name: c.FormValue("name"),
	}
	item, err := h.service.Create(newItem)
	if err != nil {
		return err
	}

	component := partials.ItemRow(item)
	return component.Render(c.Request().Context(), c.Response().Writer)
}

func (h *Handler) Delete(c echo.Context) error {
	itemId := c.Param("itemId")
	err := h.service.Delete(itemId)
	if err != nil {
		return err
	}
	return c.NoContent(200)
}
