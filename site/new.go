package site

import (
	"net/http"

	"github.com/gernest/vince/assets/ui/templates"
	"github.com/gernest/vince/models"
	"github.com/gernest/vince/render"
)

func New(w http.ResponseWriter, r *http.Request) {
	u := models.GetUser(r.Context())
	owned := u.CountOwnedSites(r.Context())
	limit := u.SitesLimit(r.Context())
	render.HTML(r.Context(), w, templates.SiteNew, http.StatusOK, func(ctx *templates.Context) {
		ctx.NewSite = &templates.NewSite{
			IsFirstSite: owned == 0,
			IsAtLimit:   limit != -1 && owned >= int64(limit),
			SiteLimit:   limit,
		}
		ctx.Page = "add_site"
	})
}
