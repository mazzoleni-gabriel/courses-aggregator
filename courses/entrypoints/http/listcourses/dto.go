package listcourses

import "github.com/mazzoleni-gabriel/courses-aggregator/courses/domain/entities"

type (
	Response struct {
		Courses []Course `json:"courses"`
	}

	Course struct {
		Name        string       `json:"name"`
		Heading     string       `json:"heading"`
		Enrollments []Enrollment `json:"enrollments"`
	}

	Enrollment struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}
)

func buildResponse(courses entities.Courses) Response {
	coursesDTO := make([]Course, len(courses.List))
	for i, c := range courses.List {
		coursesDTO[i] = buildCourse(c)
	}
	return Response{
		Courses: coursesDTO,
	}
}

func buildCourse(course entities.Course) Course {
	enrollments := make([]Enrollment, len(course.Enrollments))
	for i, e := range course.Enrollments {
		enrollments[i] = buildEnrollment(e)
	}
	return Course{
		Name:        course.Name,
		Heading:     course.Heading,
		Enrollments: enrollments,
	}
}

func buildEnrollment(enrollment entities.Enrollment) Enrollment {
	return Enrollment{
		Name:  enrollment.Name,
		Email: enrollment.Email,
	}
}
