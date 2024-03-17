package listcourses

import (
	"context"
	"net/http"

	"github.com/go-chi/chi"

	"github.com/mazzoleni-gabriel/courses-aggregator/courses/domain/entities"
	"github.com/mazzoleni-gabriel/courses-aggregator/internal/render"
)

//go:generate mockery --all --disable-version-string --with-expecter
type (
	UseCase interface {
		List(ctx context.Context, token string) (entities.Courses, error)
	}
)

type Handler struct {
	useCase UseCase
}

func NewHandler(useCase UseCase) Handler {
	return Handler{
		useCase: useCase,
	}
}

func RegisterHandler(r *chi.Mux, l Handler) {
	r.Get("/courses", l.Handle)
}

func (h Handler) Handle(w http.ResponseWriter, r *http.Request) {
	courses, err := h.useCase.List(r.Context(), r.Header.Get("apiKey"))
	if err != nil {
		render.NewError(w, r, err.Error(), http.StatusInternalServerError)
		return
	}

	render.JSON(w, r, buildResponse(courses))
}
