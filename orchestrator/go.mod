module github.com/InjectiveLabs/peggo/orchestrator

go 1.15

require (
	github.com/InjectiveLabs/sdk-go v1.14.1
	github.com/alexcesaro/statsd v2.0.0+incompatible
	github.com/bugsnag/panicwrap v1.3.0 // indirect
	github.com/cosmos/cosmos-sdk v0.41.0
	github.com/ethereum/go-ethereum v1.9.25
	github.com/gofrs/uuid v4.0.0+incompatible // indirect
	github.com/jawher/mow.cli v1.2.0
	github.com/pkg/errors v0.9.1
	github.com/shopspring/decimal v1.2.0
	github.com/tendermint/tendermint v0.34.3
	github.com/xlab/closer v0.0.0-20190328110542-03326addb7c2
	github.com/xlab/suplog v1.1.0
	google.golang.org/grpc v1.35.0
	gopkg.in/alexcesaro/statsd.v2 v2.0.0 // indirect
)

replace github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.2-alpha.regen.4
