package handlers

import (
	"time"

	"github.com/SHP-Association/E-learningWeb/backend/ent"
	"github.com/SHP-Association/E-learningWeb/backend/ent/enrollment"
	"github.com/SHP-Association/E-learningWeb/backend/pkg/middleware"
	"github.com/SHP-Association/E-learningWeb/backend/pkg/pager"
	"github.com/SHP-Association/E-learningWeb/backend/pkg/routenames"
	"github.com/SHP-Association/E-learningWeb/backend/pkg/services"
	"github.com/SHP-Association/E-learningWeb/backend/pkg/ui/models"
	"github.com/SHP-Association/E-learningWeb/backend/pkg/ui/pages"
	"github.com/labstack/echo/v4"
)

type Pages struct {
	container *services.Container
}

func init() {
	Register(new(Pages))
}

func (h *Pages) Init(c *services.Container) error {
	h.container = c
	return nil
}

func (h *Pages) Routes(g *echo.Group) {
	g.GET("/", h.Home, middleware.RequireAuthentication).Name = routenames.Home
	g.GET("/about", h.About, middleware.RequireAuthentication).Name = routenames.About
}

func (h *Pages) Home(ctx echo.Context) error {
	pgr := pager.NewPager(ctx, 4)

	return pages.Home(ctx, &models.Posts{
		Posts: h.fetchPosts(ctx, &pgr),
		Pager: pgr,
	}, h.fetchDashboardStats(ctx))
}

func (h *Pages) fetchPosts(ctx echo.Context, pgr *pager.Pager) []models.Post {
	count, _ := h.container.ORM.Course.Query().Count(ctx.Request().Context())
	pgr.SetItems(count)

	courses, err := h.container.ORM.Course.
		Query().
		Limit(pgr.ItemsPerPage).
		Offset(pgr.GetOffset()).
		Order(ent.Desc("created_at")).
		All(ctx.Request().Context())

	if err != nil {
		return nil
	}

	posts := make([]models.Post, len(courses))
	for i, c := range courses {
		posts[i] = models.Post{
			ID:     c.ID,
			Title:  c.Title,
			Body:   c.ShortDescription,
			Author: "Platform Instructor",
			Date:   c.CreatedAt.Format("Jan 02, 2006"),
		}
	}
	return posts
}

func (h *Pages) fetchDashboardStats(ctx echo.Context) *models.DashboardStats {
	stats := &models.DashboardStats{}

	// Total Users
	stats.TotalUsers, _ = h.container.ORM.User.Query().Count(ctx.Request().Context())

	// Total Courses
	stats.TotalCourses, _ = h.container.ORM.Course.Query().Count(ctx.Request().Context())

	// Total Enrollments
	stats.TotalEnrollments, _ = h.container.ORM.Enrollment.Query().Count(ctx.Request().Context())

	// Recent Enrollments (Last 24h)
	stats.RecentActivity, _ = h.container.ORM.Enrollment.
		Query().
		Where(enrollment.EnrolledAtGTE(time.Now().Add(-24 * time.Hour))).
		Count(ctx.Request().Context())

	return stats
}

func (h *Pages) About(ctx echo.Context) error {
	return pages.About(ctx)
}
