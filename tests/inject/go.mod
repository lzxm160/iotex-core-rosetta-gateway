module github.com/iotexproject/iotex-core-rosetta-gateway

go 1.13

require (
	github.com/ethereum/go-ethereum v1.8.27
	github.com/iotexproject/go-pkgs v0.1.2-0.20200523040337-5f1d9ddaa8ee
	github.com/iotexproject/iotex-address v0.2.1
	github.com/iotexproject/iotex-antenna-go/v2 v2.3.2
	github.com/iotexproject/iotex-core v0.8.1-0.20200715220723-64cde17743f2 // indirect
	github.com/iotexproject/iotex-proto v0.3.1-0.20200713201308-abf112fff818
	github.com/stretchr/testify v1.4.0
	google.golang.org/grpc v1.21.0
)

replace github.com/ethereum/go-ethereum => github.com/iotexproject/go-ethereum v0.3.0
