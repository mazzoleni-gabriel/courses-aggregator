package courses

import "github.com/mazzoleni-gabriel/courses-aggregator/courses/domain/entities"

type (
	Response struct {
		Courses []Course `json:"courses"`
		Meta    Meta     `json:"meta"`
	}

	Course struct {
		ID      int32  `json:"id"`
		Name    string `json:"name"`
		Heading string `json:"heading"`
	}

	Meta struct {
		Total         int32 `json:"total"`
		NumberOfPages int32 `json:"number_of_pages"`
	}
)

func (r Response) toDomain() entities.Courses {
	courses := make([]entities.Course, len(r.Courses))
	for i, c := range r.Courses {
		courses[i] = c.toDomain()
	}
	return entities.Courses{
		List:          courses,
		Total:         r.Meta.Total,
		NumberOfPages: r.Meta.NumberOfPages,
	}
}

func (c Course) toDomain() entities.Course {
	return entities.Course{
		ID:      c.ID,
		Name:    c.Name,
		Heading: c.Heading,
	}
}
