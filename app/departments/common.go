package departments

import (
	"github.com/google/wire"
)

var Provides = wire.NewSet(
	wire.Struct(new(Service), "*"),
)
