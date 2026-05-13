package models

import (
	"fmt"

	. "github.com/SHP-Association/E-learningWeb/backend/pkg/ui/components"
	. "maragu.dev/gomponents"
)

type DashboardStats struct {
	EnrolledCourses    int
	ActiveCourses      int
	CompletedCourses   int
	CertificatesCount  int
	EngagementScore    float64
	StudyHours         float64
	RecentProgress     *EnrollmentProgress
}

type EnrollmentProgress struct {
	CourseTitle string
	Progress    float64
	Remaining   int
}

func (s *DashboardStats) Render() Node {
	return Stats(
		Stat{
			Title:       "Enrolled Courses",
			Value:       fmt.Sprintf("%d", s.EnrolledCourses),
			Description: fmt.Sprintf("%d Active now", s.ActiveCourses),
			Icon:        Icon("PencilSquare", "w-6 h-6"),
		},
		Stat{
			Title:       "Engagement Score",
			Value:       fmt.Sprintf("%.0f%%", s.EngagementScore),
			Description: "Based on activity",
			Icon:        Icon("Info", "w-6 h-6"),
		},
		Stat{
			Title:       "Study Hours",
			Value:       fmt.Sprintf("%.1f", s.StudyHours),
			Description: "Lifetime total",
			Icon:        Icon("Clock", "w-6 h-6"),
		},
		Stat{
			Title:       "Certificates",
			Value:       fmt.Sprintf("%02d", s.CertificatesCount),
			Description: "Earned so far",
			Icon:        Icon("CheckCircle", "w-6 h-6"),
		},
	)
}
