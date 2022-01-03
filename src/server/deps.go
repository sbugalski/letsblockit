package server

import (
	"io"

	"github.com/labstack/echo/v4"
	"github.com/xvello/letsblockit/src/filters"
	"github.com/xvello/letsblockit/src/pages"
)

type PageRenderer interface {
	RegisterHelpers(helpers map[string]interface{})
	RenderRaw(name string, data *pages.Context) ([]byte, error)
	Render(c echo.Context, name string, data *pages.Context) error
	RenderWithSidebar(c echo.Context, name, sidebar string, data *pages.Context) error
}

type FilterRepository interface {
	GetFilter(name string) (*filters.Filter, error)
	GetFilters() []*filters.Filter
	GetTags() []string
	Render(w io.Writer, name string, data map[string]interface{}) error
}
