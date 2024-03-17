package listcourses

import (
	"context"

	"github.com/mazzoleni-gabriel/courses-aggregator/courses/domain/entities"
)

type UseCase struct {
	coursesFinder               CoursesFinder
	enrollmentsMapBuilder       EnrollmentsMapBuilder
	coursesEnrollmentsDecorator CoursesEnrollmentsDecorator
}

func NewUseCase(coursesFinder CoursesFinder, enrollmentsMapBuilder EnrollmentsMapBuilder, coursesEnrollmentsDecorator CoursesEnrollmentsDecorator) UseCase {
	return UseCase{
		coursesFinder:               coursesFinder,
		enrollmentsMapBuilder:       enrollmentsMapBuilder,
		coursesEnrollmentsDecorator: coursesEnrollmentsDecorator,
	}
}

type (
	CoursesFinder interface {
		ListAllPages(ctx context.Context, filters entities.CoursesFilters, token string) (entities.Courses, error)
	}

	EnrollmentsMapBuilder interface {
		Build(ctx context.Context, token string) (map[int32]entities.Enrollment, error)
	}

	CoursesEnrollmentsDecorator interface {
		Decorate(ctx context.Context, courses entities.Courses, token string) (entities.Courses, error)
	}
)

func (u UseCase) List(ctx context.Context, token string) (entities.Courses, error) {
	filters := entities.CoursesFilters{
		IsPublished: true,
	}

	courses, err := u.coursesFinder.ListAllPages(ctx, filters, token)
	if err != nil {
		return courses, err
	}

	courses, err = u.coursesEnrollmentsDecorator.Decorate(ctx, courses, token)
	if err != nil {
		return courses, err
	}

	enrollmentsMap, err := u.enrollmentsMapBuilder.Build(ctx, token)
	if err != nil {
		return courses, err
	}

	return decorateEnrollments(courses, enrollmentsMap), nil
}

// @todo extract function
func decorateEnrollments(courses entities.Courses, enrollmentsMap map[int32]entities.Enrollment) entities.Courses {
	for i, c := range courses.List {
		for j, e := range c.Enrollments {
			courses.List[i].Enrollments[j] = enrollmentsMap[e.ID]
			//@todo call users API if not found on map
		}
	}
	return courses
}
