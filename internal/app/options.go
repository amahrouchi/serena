package app

import (
	"github.com/amahrouchi/serena/internal/blockchain"
	"github.com/amahrouchi/serena/internal/core"
	"github.com/amahrouchi/serena/internal/utils"
	"go.uber.org/fx"
)

// Options registers the app package FX options.
var Options = fx.Options(
	core.Options,
	utils.Options,
	blockchain.Options,
)
