package buildenrollmentsmap

import (
	"context"

	"github.com/mazzoleni-gabriel/courses-aggregator/courses/domain/entities"
)

const pageSize = 2 // @todo make it a configuration

type Service struct {
	courseEnrollmentsRepository UsersRepository
}

func NewService(courseEnrollmentsRepository UsersRepository) Service {
	return Service{
		courseEnrollmentsRepository: courseEnrollmentsRepository,
	}
}

type (
	UsersRepository interface {
		Get(ctx context.Context, filters entities.UsersFilters, token string) (entities.Enrollments, error)
	}
)

func (s Service) Build(ctx context.Context, token string) (map[int32]entities.Enrollment, error) {
	filters := entities.UsersFilters{
		Per:  pageSize,
		Page: 1,
	}
	enrollments, err := s.courseEnrollmentsRepository.Get(ctx, filters, token)
	if err != nil {
		return map[int32]entities.Enrollment{}, err
	}
	if enrollments.NumberOfPages == 1 {
		return map[int32]entities.Enrollment{}, nil
	}

	for page := int32(2); page <= enrollments.NumberOfPages; page++ {
		filters.Page = page
		page, err := s.courseEnrollmentsRepository.Get(ctx, filters, token)
		if err != nil {
			return map[int32]entities.Enrollment{}, err
		}
		enrollments.List = append(enrollments.List, page.List...)
	}

	return buildMap(enrollments), nil
}

func buildMap(enrollments entities.Enrollments) map[int32]entities.Enrollment {
	m := make(map[int32]entities.Enrollment, len(enrollments.List))
	for _, e := range enrollments.List {
		m[e.ID] = e
	}
	return m
}
