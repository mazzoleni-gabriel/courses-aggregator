package courseenrollment

import "github.com/mazzoleni-gabriel/courses-aggregator/courses/domain/entities"

type (
	Response struct {
		Enrollments []Enrollment `json:"enrollments"`
	}

	Enrollment struct {
		UserID int32 `json:"user_id"`
	}
)

func (r Response) toDomain() []entities.Enrollment {
	var enrollments []entities.Enrollment
	for _, u := range r.Enrollments {
		enrollments = append(enrollments, u.toDomain())
	}
	return enrollments
}

func (r Enrollment) toDomain() entities.Enrollment {
	return entities.Enrollment{
		ID: r.UserID,
	}
}
