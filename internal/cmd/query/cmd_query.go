package query

import (
	"os"

	"github.com/urfave/cli/v3"
	"github.com/vinceanalytics/vince/internal/cmd/ansi"
	"github.com/vinceanalytics/vince/internal/cmd/auth"
	"github.com/vinceanalytics/vince/internal/cmd/output"
	"github.com/vinceanalytics/vince/internal/query"
)

func CMD() *cli.Command {
	return &cli.Command{
		Name:  "query",
		Usage: "connect to vince and execute sql query",
		Action: func(ctx *cli.Context) error {
			a := ctx.Args().First()
			if a == "" {
				return nil
			}
			o, _ := auth.LoadClient()
			if o.Active == nil {
				ansi.Err("no active account found")
				ansi.Suggestion(
					"log in to a vince instance with [vince login] command",
					"select existing vince instance/account using [vince use] command",
				)
				os.Exit(1)
			}
			dsn := query.ParseDSN(&o)
			db, err := query.Open(dsn)
			if err != nil {
				ansi.Err("failed connecting to instance :%v", err)
				return nil
			}
			err = db.Ping()
			if err != nil {
				ansi.Err("can't reach vince instance :%v", err)
				return err
			}
			defer db.Close()
			rows, err := db.Query(a)
			if err != nil {
				return ansi.ERROR(err)
			}
			defer rows.Close()
			cols, err := output.Build(rows)
			if err != nil {
				return ansi.ERROR(err)
			}
			return ansi.ERROR(
				output.Tab(os.Stdout, cols),
			)
		},
	}
}