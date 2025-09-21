package handler

import (
	"fmt"
	"net/http"

	"github.com/loadept/loadept.com/internal/service"
	"github.com/loadept/loadept.com/pkg/logger"
	"github.com/loadept/loadept.com/pkg/respond"
)

type SitemapHandler struct {
	service *service.SitemapService
}

func NewSitemapHandler(service *service.SitemapService) *SitemapHandler {
	return &SitemapHandler{service: service}
}

func (h *SitemapHandler) GetSitemap(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet && r.Method != http.MethodHead {
		respond.JSON(w, respond.Map{
			"detail": "Method '" + r.Method + "' not allowed",
		}, http.StatusMethodNotAllowed)
		return
	}

	requestCtx := r.Context()

	articles, err := h.service.GetArticles(requestCtx)
	if err != nil {
		respond.JSON(w, respond.Map{
			"detail": "An error occurred while retrieving results",
		}, http.StatusInternalServerError)

		logger.ERROR.Printf("An error occurred while retrieving categories: %v\n", err)
		return
	}

	w.Header().Set("Content-Type", "application/xml")

	fmt.Fprint(w, `<?xml version="1.0" encoding="utf-8" standalone="yes"?>`)
	fmt.Fprint(w, `<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9" xmlns:xhtml="http://www.w3.org/1999/xhtml">`)
	fmt.Fprint(w, `
		<url>
			<loc>https://loadept.com/</loc>
			<lastmod>2025-02-19</lastmod>
		</url>
		<url>
			<loc>https://loadept.com/about</loc>
			<lastmod>2025-03-01</lastmod>
		</url>
	`)

	for _, article := range articles {
		fmt.Fprintf(w, `
		  <url>
			<loc>https://loadept.com/%s</loc>
			<lastmod>2025-03-01</lastmod>
		</url>`, article)
	}

	fmt.Fprint(w, `</urlset>`)
}
