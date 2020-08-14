module github.com/iotexproject/iotex-core-rosetta-gateway

go 1.13

require (
	github.com/ethereum/go-ethereum v1.8.27
	github.com/iotexproject/go-pkgs v0.1.2-0.20200523040337-5f1d9ddaa8ee
	github.com/iotexproject/iotex-address v0.2.1
	github.com/iotexproject/iotex-antenna-go/v2 v2.4.0 // indirect
	github.com/iotexproject/iotex-core v0.8.1-0.20200713031334-9be4cb0f24ed
	github.com/iotexproject/iotex-proto v0.4.0
	github.com/stretchr/testify v1.4.0
	google.golang.org/grpc v1.27.0
)

replace github.com/ethereum/go-ethereum => github.com/iotexproject/go-ethereum v0.3.0
replace github.com/iotexproject/iotex-antenna-go/v2 => github.com/lzxm160/iotex-antenna-go/v2 v2.4.1