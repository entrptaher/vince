package api

import (
	"database/sql"
	"net/http"

	"github.com/vinceanalytics/vince/internal/cmd/output"
	"github.com/vinceanalytics/vince/internal/core"
	"github.com/vinceanalytics/vince/internal/pj"
	"github.com/vinceanalytics/vince/internal/query"
	"github.com/vinceanalytics/vince/internal/render"
	v1 "github.com/vinceanalytics/vince/proto/v1"
	"google.golang.org/protobuf/types/known/durationpb"
)

func Query(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var qr v1.Query_RequestOptions
	err := pj.UnmarshalDefault(&qr, r.Body)
	if err != nil {
		render.ERROR(w, http.StatusBadRequest)
		return
	}
	if qr.Query == "" {
		render.ERROR(w, http.StatusBadRequest, "query string is required")
		return
	}
	params := make([]any, len(qr.Params))
	for i := range params {
		params[i] = sql.Named(qr.Params[i].Name, qr.Params[i].Value.Interface())
	}
	db := query.GetInternalClient(ctx)
	start := core.Now(ctx)
	rows, err := db.Query(qr.Query, params...)
	elapsed := core.Now(ctx).Sub(start)
	if err != nil {
		render.ERROR(w, http.StatusBadRequest, err.Error())
		return
	}
	defer rows.Close()
	result, err := output.Build(rows)
	if err != nil {
		render.ERROR(w, http.StatusBadRequest, err.Error())
		return
	}
	result.Elapsed = durationpb.New(elapsed)
	render.JSON(w, http.StatusOK, result)
}