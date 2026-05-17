package handlers

import (
	"net/http"
	"strings"

	"github.com/SHP-Association/E-learningWeb/backend/pkg/context"
	"github.com/SHP-Association/E-learningWeb/backend/pkg/middleware"
	"github.com/SHP-Association/E-learningWeb/backend/pkg/services"
	files "github.com/SHP-Association/E-learningWeb/backend/public"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
	echomw "github.com/labstack/echo/v4/middleware"
)

// BuildRouter builds the router.
func BuildRouter(c *services.Container) error {
	// Force HTTPS, if enabled.
	if c.Config.HTTP.TLSEnabled {
		c.Web.Use(echomw.HTTPSRedirect())
	}

	// Serve public files with cache control.
	c.Web.Group("", middleware.CacheControl(c.Config.Cache.ExpirationPublicFile)).
		Static("files", "public/files")

	// Serve static files.
	// ui.StaticFile() should be used in ui components to append a cache key to the URL to break cache
	// after each server reboot.
	c.Web.Group(
		"",
		echomw.GzipWithConfig(echomw.GzipConfig{
			Skipper: func(c echo.Context) bool {
				for _, ext := range []string{
					".js",
					".css",
				} {
					if strings.HasSuffix(c.Request().URL.Path, ext) {
						return false
					}
				}
				return true
			},
		}),
		middleware.CacheControl(c.Config.Cache.ExpirationPublicFile),
	).StaticFS("static", echo.MustSubFS(files.Static, "static"))

	// CORS configuration
	origins := strings.Split(c.Config.CORS.AllowOrigins, ",")
	for i, o := range origins {
		origins[i] = strings.TrimSpace(o)
	}

	c.Web.Use(echomw.CORSWithConfig(echomw.CORSConfig{
		AllowOrigins:     origins,
		AllowHeaders:     strings.Split(c.Config.CORS.AllowHeaders, ","),
		AllowMethods:     strings.Split(c.Config.CORS.AllowMethods, ","),
		AllowCredentials: true,
	}))

	// Non-static file route group.
	g := c.Web.Group("")

	// Create a cookie store for session data.
	cookieStore := sessions.NewCookieStore([]byte(c.Config.App.EncryptionKey))
	cookieStore.Options.HttpOnly = true
	cookieStore.Options.SameSite = http.SameSiteLaxMode

	g.Use(
		echomw.RemoveTrailingSlashWithConfig(echomw.TrailingSlashConfig{
			RedirectCode: http.StatusMovedPermanently,
		}),
		echomw.Recover(),
		echomw.Secure(),
		echomw.RequestID(),
		middleware.SetLogger(),
		middleware.LogRequest(),
		echomw.Gzip(),
		middleware.Config(c.Config),
		middleware.Session(cookieStore),
		middleware.LoadAuthenticatedUser(c.Auth),
		echomw.CSRFWithConfig(echomw.CSRFConfig{
			TokenLookup:    "header:X-CSRF-Token,form:csrf,header:X-Csrf-Token",
			CookieHTTPOnly: false,
			CookieSameSite: http.SameSiteLaxMode,
			ContextKey:     context.CSRFKey,
			Skipper: func(ctx echo.Context) bool {
				path := ctx.Path()
				return path == "/api/health" ||
					path == "/api/auth/login" ||
					path == "/api/login" ||
					path == "/api/auth/register" ||
					path == "/api/register"
			},
		}),
	)

	// Error handler.
	c.Web.HTTPErrorHandler = new(Error).Page

	// Initialize and register all handlers.
	for _, h := range GetHandlers() {
		if err := h.Init(c); err != nil {
			return err
		}

		h.Routes(g)
	}

	return nil
}
