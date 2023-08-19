package login

import (
	"context"
	"crypto/ed25519"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/urfave/cli/v3"
	"github.com/vinceanalytics/vince/internal/cmd/ansi"
	"github.com/vinceanalytics/vince/internal/cmd/auth"
	"github.com/vinceanalytics/vince/internal/klient"
	"github.com/vinceanalytics/vince/internal/must"
	"github.com/vinceanalytics/vince/internal/pj"
	"github.com/vinceanalytics/vince/internal/tokens"
	v1 "github.com/vinceanalytics/vince/proto/v1"
)

func CMD() *cli.Command {
	return &cli.Command{
		Name:  "login",
		Usage: "Authenticate into vince instance",
		Flags: auth.Flags,
		Action: func(ctx *cli.Context) error {
			name, password := auth.Load(ctx)
			uri := ctx.Args().First()
			must.Assert(uri != "")(
				"missing instance argument",
			)
			client, file := auth.LoadClient()
			priv := ed25519.PrivateKey(client.PrivateKey)
			if client.Instance == nil {
				client.Instance = make(map[string]*v1.Client_Instance)
			}
			if client.Instance[uri] == nil {
				client.Instance[uri] = &v1.Client_Instance{
					Accounts: make(map[string]*v1.Client_Auth),
				}
			}
			// Check if we already have a valid token
			if a, ok := client.Instance[uri].Accounts[name]; ok {
				token, err := jwt.Parse(a.Token, func(t *jwt.Token) (interface{}, error) {
					return priv.Public(), nil
				})
				if err == nil && token.Valid {
					ansi.Ok("you are already signed in")
					return nil
				}

			}
			token, _ := tokens.Generate(
				context.Background(),
				priv,
				v1.Token_CLIENT,
				name,
				time.Now().Add(365*24*time.Hour),
			)
			var clientAuth v1.Client_Auth
			e := klient.POST(
				context.Background(),
				uri+"/tokens",
				&v1.Token_CreateOptions{
					Name:      name,
					Password:  password,
					Token:     token,
					PublicKey: priv.Public().(ed25519.PublicKey),
				},
				&clientAuth,
			)
			if e != nil {
				must.Assert(false)(e.Error)
			}

			client.Instance[uri].Accounts[clientAuth.Name] = &clientAuth
			if client.Active == nil {
				client.Active = &v1.Client_Active{
					Instance: uri,
					Account:  clientAuth.Name,
				}
			}
			must.One(os.WriteFile(file,
				must.Must(pj.MarshalIndent(&client))(
					"failed encoding config file",
				),
				0600))(
				"failed writing client config", "path", file,
			)
			ansi.Ok("signed in %q", uri)
			return nil
		},
	}
}