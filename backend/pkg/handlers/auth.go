package handlers

import (
	"fmt"
	"strings"

	"github.com/SHP-Association/E-learningWeb/backend/config"
	"github.com/SHP-Association/E-learningWeb/backend/ent"
	"github.com/SHP-Association/E-learningWeb/backend/ent/user"
	"github.com/SHP-Association/E-learningWeb/backend/pkg/context"
	"github.com/SHP-Association/E-learningWeb/backend/pkg/form"
	"github.com/SHP-Association/E-learningWeb/backend/pkg/log"
	"github.com/SHP-Association/E-learningWeb/backend/pkg/middleware"
	"github.com/SHP-Association/E-learningWeb/backend/pkg/msg"
	"github.com/SHP-Association/E-learningWeb/backend/pkg/redirect"
	"github.com/SHP-Association/E-learningWeb/backend/pkg/routenames"
	"github.com/SHP-Association/E-learningWeb/backend/pkg/services"
	"github.com/SHP-Association/E-learningWeb/backend/pkg/ui/forms"
	"github.com/SHP-Association/E-learningWeb/backend/pkg/ui/pages"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type Auth struct {
	config *config.Config
	auth   *services.AuthClient
	mail   *services.MailClient
	orm    *ent.Client
}

func init() {
	Register(new(Auth))
}

func (h *Auth) Init(c *services.Container) error {
	h.config = c.Config
	h.orm = c.ORM
	h.auth = c.Auth
	h.mail = c.Mail
	return nil
}

func (h *Auth) Routes(g *echo.Group) {
	g.GET("/logout", h.Logout, middleware.RequireAuthentication).Name = routenames.Logout
	g.GET("/email/verify/:token", h.VerifyEmail).Name = routenames.VerifyEmail

	noAuth := g.Group("/user", middleware.RequireNoAuthentication)
	noAuth.GET("/login", h.LoginPage).Name = routenames.Login
	noAuth.POST("/login", h.LoginSubmit).Name = routenames.LoginSubmit
	noAuth.GET("/password", h.ForgotPasswordPage).Name = routenames.ForgotPassword
	noAuth.POST("/password", h.ForgotPasswordSubmit).Name = routenames.ForgotPasswordSubmit

	resetGroup := noAuth.Group("/password/reset",
		middleware.LoadUser(h.orm),
		middleware.LoadValidPasswordToken(h.auth),
	)
	resetGroup.GET("/token/:user/:password_token/:token", h.ResetPasswordPage).Name = routenames.ResetPassword
	resetGroup.POST("/token/:user/:password_token/:token", h.ResetPasswordSubmit).Name = routenames.ResetPasswordSubmit
}

func (h *Auth) ForgotPasswordPage(ctx echo.Context) error {
	return pages.ForgotPassword(ctx, form.Get[forms.ForgotPassword](ctx))
}

func (h *Auth) ForgotPasswordSubmit(ctx echo.Context) error {
	var input forms.ForgotPassword

	succeed := func() error {
		form.Clear(ctx)
		msg.Success(ctx, "An email containing a link to reset your password will be sent to this address if it exists in our system.")
		return h.ForgotPasswordPage(ctx)
	}

	err := form.Submit(ctx, &input)

	switch err.(type) {
	case nil:
	case validator.ValidationErrors:
		return h.ForgotPasswordPage(ctx)
	default:
		return err
	}

	// Attempt to load the user.
	u, err := h.orm.User.
		Query().
		Where(user.Email(strings.ToLower(input.Email))).
		Only(ctx.Request().Context())

	switch err.(type) {
	case *ent.NotFoundError:
		return succeed()
	case nil:
	default:
		return fail(err, "error querying user during forgot password")
	}

	// Generate the token.
	token, pt, err := h.auth.GeneratePasswordResetToken(ctx, u.ID)
	if err != nil {
		return fail(err, "error generating password reset token")
	}

	log.Ctx(ctx).Info("generated password reset token",
		"user_id", u.ID,
	)

	// Email the user.
	url := ctx.Echo().Reverse(routenames.ResetPassword, u.ID, pt.ID, token)
	err = h.mail.
		Compose().
		To(u.Email).
		Subject("Reset your password").
		Body(fmt.Sprintf("Go here to reset your password: %s", h.config.App.Host+url)).
		Send(ctx)

	if err != nil {
		return fail(err, "error sending password reset email")
	}

	return succeed()
}

func (h *Auth) LoginPage(ctx echo.Context) error {
	return pages.Login(ctx, form.Get[forms.Login](ctx))
}

func (h *Auth) LoginSubmit(ctx echo.Context) error {
	var input forms.Login

	authFailed := func() error {
		input.SetFieldError("Email", "")
		input.SetFieldError("Password", "")
		msg.Error(ctx, "Invalid credentials. Please try again.")
		return h.LoginPage(ctx)
	}

	err := form.Submit(ctx, &input)

	switch err.(type) {
	case nil:
	case validator.ValidationErrors:
		return h.LoginPage(ctx)
	default:
		return err
	}

	// Attempt to load the user.
	u, err := h.orm.User.
		Query().
		Where(user.Email(strings.ToLower(input.Email))).
		Only(ctx.Request().Context())

	switch err.(type) {
	case *ent.NotFoundError:
		return authFailed()
	case nil:
	default:
		return fail(err, "error querying user during login")
	}

	// Check if the password is correct.
	err = h.auth.CheckPassword(input.Password, u.Password)
	if err != nil {
		return authFailed()
	}

	// Log the user in.
	err = h.auth.Login(ctx, u.ID)
	if err != nil {
		return fail(err, "unable to log in user")
	}

	msg.Success(ctx, fmt.Sprintf("Welcome back, %s. You are now logged in.", u.Username))

	return redirect.New(ctx).
		Route(routenames.Home).
		Go()
}

func (h *Auth) Logout(ctx echo.Context) error {
	if err := h.auth.Logout(ctx); err == nil {
		msg.Success(ctx, "You have been logged out successfully.")
	} else {
		msg.Error(ctx, "An error occurred. Please try again.")
	}
	return redirect.New(ctx).
		Route(routenames.Home).
		Go()
}


func (h *Auth) ResetPasswordPage(ctx echo.Context) error {
	return pages.ResetPassword(ctx, form.Get[forms.ResetPassword](ctx))
}

func (h *Auth) ResetPasswordSubmit(ctx echo.Context) error {
	var input forms.ResetPassword

	err := form.Submit(ctx, &input)

	switch err.(type) {
	case nil:
	case validator.ValidationErrors:
		return h.ResetPasswordPage(ctx)
	default:
		return err
	}

	// Get the requesting user.
	usr := ctx.Get(context.UserKey).(*ent.User)

	// Update the user.
	_, err = usr.
		Update().
		SetPassword(input.Password).
		Save(ctx.Request().Context())

	if err != nil {
		return fail(err, "unable to update password")
	}

	// Delete all password tokens for this user.
	err = h.auth.DeletePasswordTokens(ctx, usr.ID)
	if err != nil {
		return fail(err, "unable to delete password tokens")
	}

	msg.Success(ctx, "Your password has been updated.")
	return redirect.New(ctx).
		Route(routenames.Login).
		Go()
}

func (h *Auth) VerifyEmail(ctx echo.Context) error {
	var usr *ent.User

	// Validate the token.
	token := ctx.Param("token")
	email, err := h.auth.ValidateEmailVerificationToken(token)
	if err != nil {
		msg.Warning(ctx, "The link is either invalid or has expired.")
		return redirect.New(ctx).
			Route(routenames.Home).
			Go()
	}

	// Check if it matches the authenticated user.
	if u := ctx.Get(context.AuthenticatedUserKey); u != nil {
		authUser := u.(*ent.User)

		if authUser.Email == email {
			usr = authUser
		}
	}

	// Query to find a matching user, if needed.
	if usr == nil {
		usr, err = h.orm.User.
			Query().
			Where(user.Email(email)).
			Only(ctx.Request().Context())

		if err != nil {
			return fail(err, "query failed loading email verification token user")
		}
	}

	// Verify the user, if needed.
	if !usr.Verified {
		usr, err = usr.
			Update().
			SetVerified(true).
			Save(ctx.Request().Context())

		if err != nil {
			return fail(err, "failed to set user as verified")
		}
	}

	msg.Success(ctx, "Your email has been successfully verified.")
	return redirect.New(ctx).
		Route(routenames.Home).
		Go()
}
