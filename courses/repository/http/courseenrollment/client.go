package courseenrollment

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/mazzoleni-gabriel/courses-aggregator/courses/domain/entities"
	"github.com/mazzoleni-gabriel/courses-aggregator/internal/config"
)

const path = "v1/courses/%d/enrollments"

type Client struct {
	BaseURL string
}

func NewClient(cfg config.Configuration) Client {
	return Client{
		BaseURL: cfg.CoursesAPI.BaseURL,
	}
}

func (c Client) Get(ctx context.Context, courseID int32, token string) ([]entities.Enrollment, error) {
	withParam := fmt.Sprintf(path, courseID)
	url := fmt.Sprintf("%s/%s", c.BaseURL, withParam)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return []entities.Enrollment{}, fmt.Errorf("could not build request on course enrollments http client: %w", err)
	}
	req.Header.Add("accept", "application/json")
	req.Header.Add("apiKey", token)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return []entities.Enrollment{}, fmt.Errorf("error making http request: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		b, _ := io.ReadAll(res.Body)
		return []entities.Enrollment{}, fmt.Errorf("http request failed, status: %s | msg: %s", res.Status, string(b))
	}

	var response Response
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&response)
	if err != nil {
		return []entities.Enrollment{}, fmt.Errorf("could not decode http response body: %w", err)
	}

	return response.toDomain(), nil
}
