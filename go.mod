module github.com/Factom-Asset-Tokens/fatd

go 1.13

require (
	crawshaw.io/sqlite v0.2.5
	github.com/AdamSLevy/jsonrpc2/v14 v14.0.0
	github.com/AdamSLevy/sqlbuilder v0.0.0-20191210203204-7d84c87e7f80
	github.com/AdamSLevy/sqlitechangeset v0.0.0-20191210201651-f95453d87aff
	github.com/Factom-Asset-Tokens/factom v0.0.0-20200222022040-798896758557
	github.com/goji/httpauth v0.0.0-20160601135302-2da839ab0f4d
	github.com/nightlyone/lockfile v0.0.0-20200124072040-edb130adc195
	github.com/posener/complete v1.2.3
	github.com/rs/cors v1.7.0
	github.com/sirupsen/logrus v1.4.2
	github.com/spf13/cobra v0.0.6
	github.com/spf13/pflag v1.0.5
	github.com/stretchr/testify v1.4.0
	github.com/subchen/go-trylock/v2 v2.0.0
	github.com/wasmerio/go-ext-wasm v0.0.0-20200122131904-9f2a15374d27
	golang.org/x/sync v0.0.0-20190911185100-cd5d95a43a6e
)

replace (
	// Fixes a small annoyance when displaying defaults for custom Vars.
	github.com/spf13/pflag v1.0.5 => github.com/AdamSLevy/pflag v1.0.6-0.20191204180553-73c85c9446e1

	// Uses the LLVM backend and exposes metering. Also fixes a bug with
	// InstanceContext.Data.
	github.com/wasmerio/go-ext-wasm => github.com/AdamSLevy/go-ext-wasm v0.0.0-20191212234502-d66004a8582c
)

//replace github.com/Factom-Asset-Tokens/factom => ../factom
