package roles

import (
	"github.com/google/wire"
)

var Provides = wire.NewSet(
	wire.Struct(new(Service), "*"),
)
