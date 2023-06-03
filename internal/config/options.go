package config

import (
	"time"

	"github.com/urfave/cli/v3"
)

type Options struct {
	Listen string
	Env    string
	Url    string
	Mailer struct {
		Enabled       bool
		Name, Address string
		SMTP          struct {
			Address       string
			AuthAnonymous struct {
				Enabled bool
				Trace   string
			}

			AuthOAUTHBearer struct {
				Enabled               bool
				Username, Token, Host string
				Port                  int
			}
			AuthPlain struct {
				Enabled                      bool
				Identity, Username, Password string
			}
			EnableMailHog bool
		}
	}
	DataPath    string
	EnableEmail bool
	LogLevel    string
	Secrets     struct {
		Age, Secret string
	}
	Cors struct {
		Origin                string
		Credentials           bool
		MaxAge                int
		Headers               []string
		Expose                []string
		Methods               []string
		SendPreflightResponse bool
	}
	SuperUserId []uint64
	Firewall    struct {
		Enabled bool
		BlockIP []string
		AllowIP []string
	}
	Intervals struct {
		SiteCache, TSSync time.Duration
	}
	Backup struct {
		Enabled bool
		Dir     string
	}
	Acme struct {
		Enabled       bool
		Email, Domain string
	}
	TLS struct {
		Enabled            bool
		Address, Key, Cert string
	}
	Bootstrap struct {
		Enabled                    bool
		Name, Email, Password, Key string
	}
	Alerts struct {
		Enabled bool
		Source  string
	}
	EnableProfile bool
}

func (o *Options) Flags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Category:    "core",
			Name:        "listen",
			Usage:       "http address to listen to",
			Value:       ":8080",
			Destination: &o.Listen,
			EnvVars:     []string{"VINCE_LISTEN"},
		},
		&cli.StringFlag{
			Category:    "core",
			Name:        "log-level",
			Usage:       "log level, values are (trace,debug,info,warn,error,fatal,panic)",
			Value:       "debug",
			Destination: &o.LogLevel,
			EnvVars:     []string{"VINCE_LOG_LEVEL"},
		},
		&cli.BoolFlag{
			Category:    "core.tls",
			Name:        "enable-tls",
			Usage:       "Enables serving https traffic.",
			Destination: &o.TLS.Enabled,
			EnvVars:     []string{"VINCE_ENABLE_TLS"},
		},
		&cli.StringFlag{
			Category:    "core.tls",
			Name:        "tls-address",
			Usage:       "https address to listen to. You must provide tls-key and tls-cert or configure auto-tls",
			Value:       ":8443",
			Destination: &o.TLS.Address,
			EnvVars:     []string{"VINCE_TLS_LISTEN"},
		},
		&cli.StringFlag{
			Category:    "core.tls",
			Name:        "tls-key",
			Usage:       "Path to key file used for https",
			Destination: &o.TLS.Key,
			EnvVars:     []string{"VINCE_TLS_KEY"},
		},
		&cli.StringFlag{
			Category:    "core.tls",
			Name:        "tls-cert",
			Usage:       "Path to certificate file used for https",
			Destination: &o.TLS.Cert,
			EnvVars:     []string{"VINCE_TLS_CERT"},
		},
		&cli.StringFlag{
			Category:    "core",
			Name:        "data",
			Usage:       "path to data directory",
			Value:       ".vince",
			Destination: &o.DataPath,
			EnvVars:     []string{"VINCE_DATA"},
		},
		&cli.StringFlag{
			Category:    "core",
			Name:        "env",
			Usage:       "environment on which vince is run (dev,staging,production)",
			Value:       "dev",
			Destination: &o.Env,
			EnvVars:     []string{"VINCE_ENV"},
		},
		&cli.StringFlag{
			Category:    "core",
			Name:        "url",
			Usage:       "url for the server on which vince is hosted(it shows up on emails)",
			Value:       "http://localhost:8080",
			Destination: &o.Url,
			EnvVars:     []string{"VINCE_URL"},
		},
		&cli.BoolFlag{
			Category:    "core.backup",
			Name:        "enable-backup",
			Usage:       "Allows backing up and restoring",
			Destination: &o.Backup.Enabled,
			EnvVars:     []string{"VINCE_BACKUP_ENABLED"},
		},
		&cli.StringFlag{
			Category:    "core.backup",
			Name:        "backup-dir",
			Usage:       "directory where backups are stored",
			Destination: &o.Backup.Dir,
			EnvVars:     []string{"VINCE_BACKUP_DIR"},
		},

		&cli.BoolFlag{
			Category:    "core.mailer",
			Name:        "enable-email",
			Usage:       "allows sending emails",
			Destination: &o.Mailer.Enabled,
			EnvVars:     []string{"VINCE_ENABLE_EMAIL"},
		},

		&cli.StringFlag{
			Category:    "core.mailer",
			Name:        "mailer-address",
			Usage:       "email address used for the sender of outgoing emails ",
			Value:       "vince@mailhog.example",
			Destination: &o.Mailer.Address,
			EnvVars:     []string{"VINCE_MAILER_ADDRESS"},
		},
		&cli.StringFlag{
			Category:    "core.mailer",
			Name:        "mailer-address-name",
			Usage:       "email address name  used for the sender of outgoing emails ",
			Value:       "gernest from vince analytics",
			Destination: &o.Mailer.Name,
			EnvVars:     []string{"VINCE_MAILER_ADDRESS_NAME"},
		},
		&cli.StringFlag{
			Category:    "core.mailer.smtp",
			Name:        "mailer-smtp-address",
			Usage:       "host:port address of the smtp server used for outgoing emails",
			Value:       "localhost:1025",
			Destination: &o.Mailer.SMTP.Address,
			EnvVars:     []string{"VINCE_MAILER_SMTP_ADDRESS"},
		},
		&cli.BoolFlag{
			Category:    "core.mailer.smtp",
			Name:        "mailer-smtp-enable-mailhog",
			Usage:       "port address of the smtp server used for outgoing emails",
			Value:       true,
			Destination: &o.Mailer.SMTP.EnableMailHog,
			EnvVars:     []string{"VINCE_MAILER_SMTP_ENABLE_MAILHOG"},
		},
		&cli.BoolFlag{
			Category:    "core.mailer.smtp.auth.anonymous",
			Name:        "mailer-smtp-anonymous-enable",
			Usage:       "enables anonymous authenticating smtp client",
			Destination: &o.Mailer.SMTP.AuthAnonymous.Enabled,
			EnvVars:     []string{"VINCE_MAILER_SMTP_ANONYMOUS_ENABLED"},
		},
		&cli.StringFlag{
			Category:    "core.mailer.smtp.auth.anonymous",
			Name:        "mailer-smtp-anonymous-trace",
			Usage:       "trace value for anonymous smtp auth",
			Destination: &o.Mailer.SMTP.AuthAnonymous.Trace,
			EnvVars:     []string{"VINCE_MAILER_SMTP_ANONYMOUS_TRACE"},
		},
		&cli.BoolFlag{
			Category:    "core.mailer.smtp.auth.plain",
			Name:        "mailer-smtp-plain-enabled",
			Usage:       "enables PLAIN authentication of smtp client",
			Destination: &o.Mailer.SMTP.AuthPlain.Enabled,
			EnvVars:     []string{"VINCE_MAILER_SMTP_PLAIN_ENABLED"},
		},
		&cli.StringFlag{
			Category:    "core.mailer.smtp.auth.plain",
			Name:        "mailer-smtp-plain-identity",
			Usage:       "identity value for plain smtp auth",
			Destination: &o.Mailer.SMTP.AuthPlain.Identity,
			EnvVars:     []string{"VINCE_MAILER_SMTP_PLAIN_IDENTITY"},
		},
		&cli.StringFlag{
			Category:    "core.mailer.smtp.auth.plain",
			Name:        "mailer-smtp-plain-username",
			Usage:       "username value for plain smtp auth",
			Destination: &o.Mailer.SMTP.AuthPlain.Username,
			EnvVars:     []string{"VINCE_MAILER_SMTP_PLAIN_USERNAME"},
		},
		&cli.StringFlag{
			Category:    "core.mailer.smtp.auth.plain",
			Name:        "mailer-smtp-plain-password",
			Usage:       "password value for plain smtp auth",
			Destination: &o.Mailer.SMTP.AuthPlain.Password,
			EnvVars:     []string{"VINCE_MAILER_SMTP_PLAIN_PASSWORD"},
		},
		&cli.BoolFlag{
			Category:    "core.mailer.smtp.auth.oauth",
			Name:        "mailer-smtp-oauth-username",
			Usage:       "allows oauth authentication on smtp client",
			Destination: &o.Mailer.SMTP.AuthOAUTHBearer.Enabled,
			EnvVars:     []string{"VINCE_MAILER_SMTP_OAUTH_USERNAME"},
		},
		&cli.StringFlag{
			Category:    "core.mailer.smtp.auth.oauth",
			Name:        "mailer-smtp-oauth-token",
			Usage:       "token value for oauth bearer smtp auth",
			Destination: &o.Mailer.SMTP.AuthOAUTHBearer.Token,
			EnvVars:     []string{"VINCE_MAILER_SMTP_OAUTH_TOKEN"},
		},
		&cli.StringFlag{
			Category:    "core.mailer.smtp.auth.oauth",
			Name:        "mailer-smtp-oauth-host",
			Usage:       "host value for oauth bearer smtp auth",
			Destination: &o.Mailer.SMTP.AuthOAUTHBearer.Host,
			EnvVars:     []string{"VINCE_MAILER_SMTP_OAUTH_HOST"},
		},
		&cli.IntFlag{
			Category:    "core.mailer.smtp.auth.oauth",
			Name:        "mailer-smtp-oauth-port",
			Usage:       "port value for oauth bearer smtp auth",
			Destination: &o.Mailer.SMTP.AuthOAUTHBearer.Port,
			EnvVars:     []string{"VINCE_MAILER_SMTP_OAUTH_PORT"},
		},
		&cli.DurationFlag{
			Category:    "core.intervals",
			Name:        "cache-refresh-interval",
			Usage:       "window for refreshing sites cache",
			Value:       15 * time.Minute,
			Destination: &o.Intervals.SiteCache,
			EnvVars:     []string{"VINCE_SITE_CACHE_REFRESH_INTERVAL"},
		},
		&cli.DurationFlag{
			Category: "core.intervals",
			Name:     "ts-buffer-sync-interval",
			Usage:    "window for buffering timeseries in memory before savin them",
			// This seems reasonable to avoid users to wait for a long time between
			// creating the site and seeing something on the dashboard. A bigger
			// duration is better though, to reduce pressure on our kv store
			Value:       time.Minute,
			Destination: &o.Intervals.TSSync,
			EnvVars:     []string{"VINCE_TS_BUFFER_INTERVAL"},
		},
		// secrets
		&cli.StringFlag{
			Category:    "core.secrets",
			Name:        "secret",
			Usage:       "path to a file with  ed25519 private key",
			Destination: &o.Secrets.Secret,
			EnvVars:     []string{"VINCE_SECRET"},
		},
		&cli.StringFlag{
			Category:    "core.secrets",
			Name:        "secret-age",
			Usage:       "path to file with age.X25519Identity",
			Destination: &o.Secrets.Age,
			EnvVars:     []string{"VINCE_SECRET_AGE"},
		},
		&cli.BoolFlag{
			Category:    "core.tls.acme",
			Name:        "enable-auto-tls",
			Usage:       "Enables using acme for automatic https.",
			Destination: &o.Acme.Enabled,
			EnvVars:     []string{"VINCE_AUTO_TLS"},
		},
		&cli.StringFlag{
			Category:    "core.tls.acme",
			Name:        "acme-email",
			Usage:       "Email address to use with letsencrypt.",
			Destination: &o.Acme.Email,
			EnvVars:     []string{"VINCE_ACME_EMAIL"},
		},
		&cli.StringFlag{
			Category:    "core.tls.acme",
			Name:        "acme-domain",
			Usage:       "Domain to use with letsencrypt.",
			Destination: &o.Acme.Domain,
			EnvVars:     []string{"VINCE_ACME_DOMAIN"},
		},
		&cli.BoolFlag{
			Category:    "core.bootstrap",
			Name:        "enable-bootstrap",
			Usage:       "allows creating a user and api key on startup.",
			Destination: &o.Bootstrap.Enabled,
			EnvVars:     []string{"VINCE_ENABLE_BOOTSTRAP"},
		},
		&cli.StringFlag{
			Category:    "core.bootstrap",
			Name:        "bootstrap-name",
			Usage:       "Full name of the user to bootstrap.",
			Destination: &o.Bootstrap.Name,
			EnvVars:     []string{"VINCE_BOOTSTRAP_NAME"},
		},
		&cli.StringFlag{
			Category:    "core.bootstrap",
			Name:        "bootstrap-email",
			Usage:       "Email address of the user to bootstrap.",
			Destination: &o.Bootstrap.Email,
			EnvVars:     []string{"VINCE_BOOTSTRAP_EMAIL"},
		},
		&cli.StringFlag{
			Category:    "core.bootstrap",
			Name:        "bootstrap-password",
			Usage:       "Password of the user to bootstrap.",
			Destination: &o.Bootstrap.Password,
			EnvVars:     []string{"VINCE_BOOTSTRAP_PASSWORD"},
		},
		&cli.StringFlag{
			Category:    "core.bootstrap",
			Name:        "bootstrap-key",
			Usage:       "API Key of the user to bootstrap.",
			Destination: &o.Bootstrap.Key,
			EnvVars:     []string{"VINCE_BOOTSTRAP_KEY"},
		},
		&cli.BoolFlag{
			Category:    "core",
			Name:        "enable-profile",
			Usage:       "Expose /debug/pprof endpoint",
			Destination: &o.EnableProfile,
			EnvVars:     []string{"VINCE_ENABLE_PROFILE"},
		},
		&cli.BoolFlag{
			Category:    "core.alerts",
			Name:        "enable-alerts",
			Usage:       "allows loading and executing alerts",
			Destination: &o.Alerts.Enabled,
			EnvVars:     []string{"VINCE_ENABLE_ALERTS"},
		},
		&cli.StringFlag{
			Category:    "core.alerts",
			Name:        "alerts-source",
			Usage:       "path to directory with alerts scripts",
			Destination: &o.Alerts.Source,
			EnvVars:     []string{"VINCE_ALERTS_SOURCE"},
		},

		&cli.StringFlag{
			Category:    "core.cors",
			Name:        "cors-origin",
			Value:       "*",
			Destination: &o.Cors.Origin,
			EnvVars:     []string{"VINCE_CORS_ORIGIN"},
		},
		&cli.BoolFlag{
			Category:    "core.cors",
			Name:        "cors-credentials",
			Value:       true,
			Destination: &o.Cors.Credentials,
			EnvVars:     []string{"VINCE_CORS_ORIGIN"},
		},
		&cli.IntFlag{
			Category:    "core.cors",
			Name:        "cors-max-age",
			Value:       1_728_000,
			Destination: &o.Cors.MaxAge,
			EnvVars:     []string{"VINCE_CORS_MAX_AGE"},
		},
		&cli.StringSliceFlag{
			Category:    "core.cors",
			Name:        "cors-headers",
			Value:       []string{"Authorization", "Content-Type", "Accept", "Origin", "User-Agent", "DNT", "Cache-Control", "X-Mx-ReqToken", "Keep-Alive", "X-Requested-With", "If-Modified-Since", "X-CSRF-Token"},
			Destination: &o.Cors.Headers,
			EnvVars:     []string{"VINCE_CORS_HEADERS"},
		},
		&cli.StringSliceFlag{
			Category:    "core.cors",
			Name:        "cors-expose",
			Destination: &o.Cors.Expose,
			EnvVars:     []string{"VINCE_CORS_EXPOSE"},
		},
		&cli.StringSliceFlag{
			Category:    "core.cors",
			Name:        "cors-methods",
			Value:       []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
			Destination: &o.Cors.Methods,
			EnvVars:     []string{"VINCE_CORS_METHODS"},
		},
		&cli.BoolFlag{
			Category:    "core.cors",
			Name:        "cors-send-preflight-response",
			Value:       true,
			Destination: &o.Cors.SendPreflightResponse,
			EnvVars:     []string{"VINCE_CORS_SEND_PREFLIGHT_RESPONSE"},
		},
		&cli.Uint64SliceFlag{
			Category:    "core",
			Name:        "super-users",
			Usage:       "a list of user ID with super privilege",
			Destination: &o.SuperUserId,
			EnvVars:     []string{"VINCE_SUPER_USERS"},
		},
		&cli.BoolFlag{
			Category:    "core.firewall",
			Name:        "enable-firewall",
			Usage:       "allow blocking ip address",
			Destination: &o.Firewall.Enabled,
			EnvVars:     []string{"VINCE_ENABLE_FIREWALL"},
		},
		&cli.StringSliceFlag{
			Category:    "core.firewall",
			Name:        "firewall-block-list",
			Usage:       "block  ip address from this list",
			Destination: &o.Firewall.BlockIP,
			EnvVars:     []string{"VINCE_FIREWALL_BLOCK_LIST"},
		},
		&cli.StringSliceFlag{
			Category:    "core.firewall",
			Name:        "firewall-allow-list",
			Usage:       "allow  ip address from this list",
			Destination: &o.Firewall.AllowIP,
			EnvVars:     []string{"VINCE_FIREWALL_ALLOW_LIST"},
		},
	}
}

func (o *Options) IsSuperUser(uid uint64) bool {
	for _, v := range o.SuperUserId {
		if v == uid {
			return true
		}
	}
	return false
}