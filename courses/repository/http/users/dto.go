package users

import "github.com/mazzoleni-gabriel/courses-aggregator/courses/domain/entities"

type (
	Response struct {
		Users []User `json:"users"`
		Meta  Meta   `json:"meta"`
	}

	User struct {
		ID    int32  `json:"id"`
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	Meta struct {
		Total         int32 `json:"total"`
		NumberOfPages int32 `json:"number_of_pages"`
	}
)

func (r Response) toDomain() entities.Enrollments {
	var enrollments entities.Enrollments
	for _, u := range r.Users {
		enrollments.List = append(enrollments.List, u.toDomain())
	}
	enrollments.Total = r.Meta.Total
	enrollments.NumberOfPages = r.Meta.NumberOfPages
	return enrollments
}

func (u User) toDomain() entities.Enrollment {
	return entities.Enrollment{
		ID:    u.ID,
		Name:  u.Name,
		Email: u.Email,
	}
}
