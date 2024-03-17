package decoratecoursesenrollments

import (
	"context"

	"github.com/mazzoleni-gabriel/courses-aggregator/courses/domain/entities"
)

type Service struct {
	courseEnrollmentsRepository CourseEnrollmentsRepository
}

func NewService(courseEnrollmentsRepository CourseEnrollmentsRepository) Service {
	return Service{
		courseEnrollmentsRepository: courseEnrollmentsRepository,
	}
}

type (
	CourseEnrollmentsRepository interface {
		Get(ctx context.Context, courseID int32, token string) ([]entities.Enrollment, error)
	}
)

func (s Service) Decorate(ctx context.Context, courses entities.Courses, token string) (entities.Courses, error) {
	for i, c := range courses.List {
		decorated, err := s.decorateCourse(ctx, c, token)
		if err != nil {
			return courses, err
		}
		courses.List[i] = decorated
	}
	return courses, nil
}

func (s Service) decorateCourse(ctx context.Context, course entities.Course, token string) (entities.Course, error) {
	enrollments, err := s.courseEnrollmentsRepository.Get(ctx, course.ID, token) //@todo validate pagination
	if err != nil {
		return course, err
	}
	course.Enrollments = enrollments
	return course, nil
}
