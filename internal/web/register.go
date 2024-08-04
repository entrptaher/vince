package web

import (
	"net/http"
	"strings"

	"github.com/vinceanalytics/vince/internal/kv"
	"github.com/vinceanalytics/vince/internal/web/db"
)

func RegisterForm(db *db.Config, w http.ResponseWriter, r *http.Request) {
	register.Execute(w, db.Context(make(map[string]any)))
}

func Register(db *db.Config, w http.ResponseWriter, r *http.Request) {
	usr := new(kv.User)
	m, err := usr.NewUser(r)
	if err != nil {
		e500.Execute(w, db.Context(make(map[string]any)))
		db.Logger().Error("creating new user", "err", err)
		return
	}
	validCaptcha := db.VerifyCaptchaSolution(r)
	if len(m) > 0 || !validCaptcha {
		db.SaveCsrf(w)
		db.SaveCaptcha(w)
		if len(m) == 0 {
			m = make(map[string]any)
		}
		if !validCaptcha {
			m["validation.captcha"] = "invalid captcha"
		}
		register.Execute(w, db.Context(m))
		return
	}
	err = usr.Save(db.Get())
	if err != nil {
		if strings.Contains(err.Error(), "UNIQUE constraint failed") {
			db.SaveCsrf(w)
			db.SaveCaptcha(w)
			register.Execute(w, db.Context(map[string]any{
				"validation.email": "email already exists",
			}))
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/sites/new", http.StatusFound)
}