package config

type (
	Configuration struct {
		Addr       string
		CoursesAPI coursesAPI
	}

	coursesAPI struct {
		BaseURL string
	}
)

func NewConfig() (Configuration, error) {
	// @todo load configs from properties
	conf := Configuration{
		Addr: ":8080",
		CoursesAPI: coursesAPI{
			BaseURL: "https://developers.teachable.com",
		},
	}
	return conf, nil
}
