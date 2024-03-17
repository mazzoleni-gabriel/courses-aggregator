package entities

type (
	Courses struct {
		List          []Course
		Total         int32
		NumberOfPages int32
	}

	Course struct {
		ID          int32
		Name        string
		Heading     string
		Enrollments []Enrollment
	}

	Enrollment struct {
		ID    int32
		Name  string
		Email string
	}

	Enrollments struct {
		List          []Enrollment
		Total         int32
		NumberOfPages int32
	}
)
