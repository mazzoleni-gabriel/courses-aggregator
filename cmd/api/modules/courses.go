package modules

import (
	"go.uber.org/fx"

	"github.com/mazzoleni-gabriel/courses-aggregator/courses/domain/services/buildenrollmentsmap"
	"github.com/mazzoleni-gabriel/courses-aggregator/courses/domain/services/decoratecoursesenrollments"
	"github.com/mazzoleni-gabriel/courses-aggregator/courses/domain/services/listallcourses"
	"github.com/mazzoleni-gabriel/courses-aggregator/courses/domain/usecases/listcourses"
	listcoursesEntrypoint "github.com/mazzoleni-gabriel/courses-aggregator/courses/entrypoints/http/listcourses"
	"github.com/mazzoleni-gabriel/courses-aggregator/courses/repository/http/courseenrollment"
	"github.com/mazzoleni-gabriel/courses-aggregator/courses/repository/http/courses"
	"github.com/mazzoleni-gabriel/courses-aggregator/courses/repository/http/users"
)

var coursesFactories = fx.Provide(
	//Entrypoints
	listcoursesEntrypoint.NewHandler,

	//Use cases
	listcourses.NewUseCase,

	//Services
	listallcourses.NewService,
	buildenrollmentsmap.NewService,
	decoratecoursesenrollments.NewService,

	//Repositories
	courses.NewClient,
	courseenrollment.NewClient,
	users.NewClient,
)

var coursesTranslations = fx.Provide(

	func(u listcourses.UseCase) listcoursesEntrypoint.UseCase { return u },
	func(s listallcourses.Service) listcourses.CoursesFinder { return s },
	func(s buildenrollmentsmap.Service) listcourses.EnrollmentsMapBuilder { return s },
	func(s decoratecoursesenrollments.Service) listcourses.CoursesEnrollmentsDecorator { return s },
	func(c courses.Client) listallcourses.CoursesRepository { return c },
	func(c users.Client) buildenrollmentsmap.UsersRepository { return c },
	func(c courseenrollment.Client) decoratecoursesenrollments.CourseEnrollmentsRepository { return c },
)

var coursesEndpoints = fx.Invoke(
	listcoursesEntrypoint.RegisterHandler,
)

var coursesModule = fx.Options(
	coursesFactories,
	coursesTranslations,
	coursesEndpoints,
)
