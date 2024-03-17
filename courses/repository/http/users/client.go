package users

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/mazzoleni-gabriel/courses-aggregator/courses/domain/entities"
	"github.com/mazzoleni-gabriel/courses-aggregator/internal/config"
)

const path = "v1/users"

type Client struct {
	BaseURL string
}

func NewClient(cfg config.Configuration) Client {
	return Client{
		BaseURL: cfg.CoursesAPI.BaseURL,
	}
}

func (c Client) Get(ctx context.Context, filters entities.UsersFilters, token string) (entities.Enrollments, error) {
	url := fmt.Sprintf("%s/%s", c.BaseURL, path)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return entities.Enrollments{}, fmt.Errorf("could not build request on users http client: %w", err)
	}
	req.Header.Add("accept", "application/json")
	req.Header.Add("apiKey", token)

	queryParams := req.URL.Query()
	queryParams.Add("page", fmt.Sprintf("%d", filters.Page))
	queryParams.Add("per", fmt.Sprintf("%d", filters.Per))
	req.URL.RawQuery = queryParams.Encode()

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return entities.Enrollments{}, fmt.Errorf("error making http request: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		b, _ := io.ReadAll(res.Body)
		return entities.Enrollments{}, fmt.Errorf("http request failed, status: %s | msg: %s", res.Status, string(b))
	}

	var response Response
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&response)
	if err != nil {
		return entities.Enrollments{}, fmt.Errorf("could not decode http response body: %w", err)
	}

	return response.toDomain(), nil
}
