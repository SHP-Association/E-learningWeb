package models

import (
	"fmt"

	. "github.com/SHP-Association/E-learningWeb/backend/pkg/ui/components"
	. "maragu.dev/gomponents"
)

type DashboardStats struct {
	TotalUsers       int
	TotalCourses     int
	TotalEnrollments int
	RecentActivity   int
}

func (s *DashboardStats) Render() Node {
	return Stats(
		Stat{
			Title:       "Total Platform Users",
			Value:       fmt.Sprintf("%d", s.TotalUsers),
			Description: "Active user accounts",
			Icon:        Icon("Users", "w-6 h-6"),
		},
		Stat{
			Title:       "Course Inventory",
			Value:       fmt.Sprintf("%d", s.TotalCourses),
			Description: "Published courses",
			Icon:        Icon("BookOpen", "w-6 h-6"),
		},
		Stat{
			Title:       "Total Enrollments",
			Value:       fmt.Sprintf("%d", s.TotalEnrollments),
			Description: "Lifetime student signups",
			Icon:        Icon("AcademicCap", "w-6 h-6"),
		},
		Stat{
			Title:       "Recent Activity",
			Value:       fmt.Sprintf("%d", s.RecentActivity),
			Description: "New signups (Last 24h)",
			Icon:        Icon("ChartBar", "w-6 h-6"),
		},
	)
}
