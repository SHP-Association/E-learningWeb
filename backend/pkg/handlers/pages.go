package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/SHP-Association/E-learningWeb/backend/ent"
	"github.com/SHP-Association/E-learningWeb/backend/ent/certificate"
	"github.com/SHP-Association/E-learningWeb/backend/ent/enrollment"
	"github.com/SHP-Association/E-learningWeb/backend/ent/user"
	"github.com/SHP-Association/E-learningWeb/backend/pkg/context"
	"github.com/SHP-Association/E-learningWeb/backend/pkg/pager"
	"github.com/SHP-Association/E-learningWeb/backend/pkg/routenames"
	"github.com/SHP-Association/E-learningWeb/backend/pkg/services"
	"github.com/SHP-Association/E-learningWeb/backend/pkg/ui/models"
	"github.com/SHP-Association/E-learningWeb/backend/pkg/ui/pages"
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
	g.GET("/", h.Home).Name = routenames.Home
	g.GET("/about", h.About).Name = routenames.About
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
	stats := &models.DashboardStats{
		EnrolledCourses:   0,
		ActiveCourses:     0,
		CompletedCourses:  0,
		CertificatesCount: 0,
		EngagementScore:   0,
		StudyHours:        0,
	}

	u := ctx.Get(context.AuthenticatedUserKey)
	if u == nil {
		return stats
	}
	usr := u.(*ent.User)

	// Fetch enrollment counts
	enrollments, err := h.container.ORM.Enrollment.
		Query().
		Where(enrollment.HasStudentWith(user.ID(usr.ID))).
		WithCourse().
		All(ctx.Request().Context())

	if err != nil {
		return stats
	}

	stats.EnrolledCourses = len(enrollments)
	for _, e := range enrollments {
		if e.IsCompleted {
			stats.CompletedCourses++
		} else if e.Progress > 0 {
			stats.ActiveCourses++
		}
	}

	// Fetch certificate count
	stats.CertificatesCount, _ = h.container.ORM.Certificate.
		Query().
		Where(certificate.HasEnrollmentWith(enrollment.HasStudentWith(user.ID(usr.ID)))).
		Count(ctx.Request().Context())

	// Fetch most recent progress
	recent, err := h.container.ORM.Enrollment.
		Query().
		Where(enrollment.HasStudentWith(user.ID(usr.ID))).
		Where(enrollment.IsCompleted(false)).
		WithCourse(func(q *ent.CourseQuery) {
			q.WithLessons()
		}).
		Order(ent.Desc(enrollment.FieldEnrolledAt)). // Should ideally be updated_at but schema doesn't have it yet
		First(ctx.Request().Context())

	if err == nil && recent != nil && recent.Edges.Course != nil {
		stats.RecentProgress = &models.EnrollmentProgress{
			CourseTitle: recent.Edges.Course.Title,
			Progress:    recent.Progress,
			Remaining:   len(recent.Edges.Course.Edges.Lessons), // Simplified
		}
	}

	// Mock engagement and study hours for now as they aren't in schema
	stats.EngagementScore = 85.0
	stats.StudyHours = 24.5

	return stats
}

func (h *Pages) About(ctx echo.Context) error {
	return pages.About(ctx)
}
