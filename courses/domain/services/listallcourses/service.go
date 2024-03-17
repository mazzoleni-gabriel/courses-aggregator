package listallcourses

import (
	"context"

	"github.com/mazzoleni-gabriel/courses-aggregator/courses/domain/entities"
)

const pageSize = 2 // @todo make it a configuration

type Service struct {
	coursesRepository CoursesRepository
}

func NewService(coursesRepository CoursesRepository) Service {
	return Service{
		coursesRepository: coursesRepository,
	}
}

type (
	CoursesRepository interface {
		Get(ctx context.Context, filters entities.CoursesFilters, token string) (entities.Courses, error)
	}
)

func (s Service) ListAllPages(ctx context.Context, filters entities.CoursesFilters, token string) (entities.Courses, error) {
	filters.Page = 1
	filters.Per = pageSize
	courses, err := s.coursesRepository.Get(ctx, filters, token)
	if err != nil {
		return entities.Courses{}, err
	}
	if courses.NumberOfPages == 1 {
		return courses, nil
	}

	for page := int32(2); page <= courses.NumberOfPages; page++ {
		filters.Page = page
		page, err := s.coursesRepository.Get(ctx, filters, token)
		if err != nil {
			return entities.Courses{}, err
		}
		courses.List = append(courses.List, page.List...)
	}

	return courses, nil
}
