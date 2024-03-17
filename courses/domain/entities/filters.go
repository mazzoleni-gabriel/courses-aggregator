package entities

type (
	CoursesFilters struct {
		IsPublished bool
		Page        int32
		Per         int32
	}

	UsersFilters struct {
		Page int32
		Per  int32
	}
)
