package modules

import (
	"go.uber.org/fx"

	"github.com/mazzoleni-gabriel/courses-aggregator/internal/config"
)

func NewApp() *fx.App {
	cfg, err := config.NewConfig()
	if err != nil {
		panic(err)
	}

	return newAppWihConfig(cfg)
}

func newAppWihConfig(cfg config.Configuration) *fx.App {
	options := []fx.Option{
		fx.Provide(func() config.Configuration { return cfg }),
		internalModule,
		coursesModule,
	}

	return fx.New(
		fx.Options(options...),
	)
}
