package models

import (
	"fmt"

	. "github.com/SHP-Association/E-learningWeb/backend/pkg/ui/components"
	. "maragu.dev/gomponents"
)

type DashboardStats struct {
	TotalUsers         int
	TotalCourses       int
	TotalEnrollments   int
	CertificatesIssued int
	RecentActivity     int
	SystemTasks        int
}

func (s *DashboardStats) Render() Node {
	return Stats(
		Stat{
			Title:       "Total Platform Users",
			Value:       fmt.Sprintf("%d", s.TotalUsers),
			Description: "Active learner accounts",
			Icon:        Icon("Users", "w-6 h-6"),
		},
		Stat{
			Title:       "Curriculum Inventory",
			Value:       fmt.Sprintf("%d", s.TotalCourses),
			Description: "Active published courses",
			Icon:        Icon("BookOpen", "w-6 h-6"),
		},
		Stat{
			Title:       "Total Enrollments",
			Value:       fmt.Sprintf("%d", s.TotalEnrollments),
			Description: "Lifetime student signups",
			Icon:        Icon("AcademicCap", "w-6 h-6"),
		},
		Stat{
			Title:       "Certificates Issued",
			Value:       fmt.Sprintf("%d", s.CertificatesIssued),
			Description: "Validated completions",
			Icon:        Icon("DocumentCheck", "w-6 h-6"),
		},
	)
}
