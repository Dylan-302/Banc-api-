package infraestructure

import (
	config "banc-api/src/common/config"
	types "banc-api/src/common/types"
	adapter "banc-api/src/infraestructure/db/adapter"

	"go.uber.org/fx"
)

type ProvidersStore struct {
	Providers []fx.Option
}

func (ps *ProvidersStore) Init() {
	ps.Providers = []fx.Option{
		fx.Provide(types.NewHandlersStore),
		fx.Provide(config.NewConfig),
		fx.Provide(adapter.NewDBConnection),
	}
}
func (ps *ProvidersStore) AddModule(p []fx.Option) {
	ps.Providers = append(ps.Providers, p...)
}

func (ps *ProvidersStore) Up(lp ...[]fx.Option) {
	ps.Providers = append(ps.Providers, fx.Invoke(NewHttpFiberServer))
	fx.New(ps.Providers...).Run()
}
