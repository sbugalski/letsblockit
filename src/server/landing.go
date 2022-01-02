package server

import (
	"github.com/labstack/echo/v4"
	"github.com/xvello/letsblockit/src/filters"
)

type carouselSlide struct {
	Title   string
	tag     string
	Icon    string
	Filters map[string]*filters.Filter
	Padding []struct{}
}

var carouselSlideTemplates = []carouselSlide{{
	Title: "De-clutter Youtube",
	Icon:  "fab fa-youtube",
	tag:   "youtube",
}, {
	Title: "Google",
	Icon:  "fab fa-google",
	tag:   "google",
}, {
	Title: "Nebula",
	Icon:  "fas fa-star fa-rotate-180",
	tag:   "nebula",
}}

// buildLandingPageHandler pre-builds the carousel slides (using maps to randomize display order)
// and returns the handler for the landing page.
// If the user is authenticated, they are redirected to the filter list page.
func (s *Server) buildLandingPageHandler() echo.HandlerFunc {
	var slides = make(map[string]*carouselSlide)
	var maxFilterCount = 0
	for _, template := range carouselSlideTemplates {
		page := *&template
		page.Filters = make(map[string]*filters.Filter)
		for _, f := range s.filters.GetFilters() {
			if f.HasTag(page.tag) {
				page.Filters[f.Name] = f
			}
		}
		slides[page.tag] = &page
		if len(page.Filters) > maxFilterCount {
			maxFilterCount = len(page.Filters)
		}
	}
	for _, slide := range slides {
		for i := len(slide.Filters); i < maxFilterCount; i++ {
			slide.Padding = append(slide.Padding, struct{}{})
		}
	}

	return func(c echo.Context) error {
		hc := s.buildPageContext(c, "")
		if hc.UserLoggedIn {
			c.Response().Header().Set("HX-Push", "/filters")
			return s.listFilters(c)
		}
		hc.Add("filter_tags", s.filters.GetTags())
		hc.Add("slides", slides)
		hc.Add("slide_list_height", 2*maxFilterCount)

		return s.pages.Render(c, "landing", hc)
	}
}
