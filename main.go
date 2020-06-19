package main

import (
	"fmt"
	"os"

	//"context"
	//"fmt"
	//"github.com/trustwallet/blockatlas/pkg/logger"
	//"net/http"
	//"os"
	//"strconv"
	//"strings"
	//"time"
	//
	//"github.com/coinbase/rosetta-sdk-go/server"
	//"github.com/oasisprotocol/oasis-core/go/common/logging"
	//
	//"github.com/iotexproject/iotex-core-rosetta-gateway/iotex-client"
	"net/http"

	"github.com/coinbase/rosetta-sdk-go/server"
	ic "github.com/iotexproject/iotex-core-rosetta-gateway/iotex-client"
	"github.com/iotexproject/iotex-core-rosetta-gateway/services"
)

const (
	// GatewayPort is the name of the environment variable that specifies
	// which port the IoTex Rosetta gateway should run on.
	GatewayPort = "GatewayPort"
	// IoTexChainPoint is the name of the environment variable that specifies
	// which the IoTex blockchain endpoint.
	IoTexChainPoint = "IoTexChainPoint"
)

// NewBlockchainRouter returns a Mux http.Handler from a collection of
// Rosetta service controllers.
func NewBlockchainRouter(client ic.IoTexClient) http.Handler {
	//networkAPIController := server.NewNetworkAPIController(services.NewNetworkAPIService(client))
	accountAPIController := server.NewAccountAPIController(services.NewAccountAPIService(client))
	//blockAPIController := server.NewBlockAPIController(services.NewBlockAPIService(client))
	//constructionAPIController := server.NewConstructionAPIController(services.NewConstructionAPIService(client))

	//return server.NewRouter(networkAPIController, accountAPIController, blockAPIController, constructionAPIController)
	return server.NewRouter(nil, accountAPIController, nil, nil)
}

func main() {
	// Get server port from environment variable or use the default.
	port := os.Getenv(GatewayPort)
	if port == "" {
		port = "8080"
	}
	addr := os.Getenv(IoTexChainPoint)
	if addr == "" {
		fmt.Fprintf(os.Stderr, "ERROR: %s environment variable missing\n", IoTexChainPoint)
		os.Exit(1)
	}

	// Prepare a new gRPC client.
	client, err := ic.NewIoTexClient(addr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: Failed to prepare Oasis gRPC client: %v\n", err)
		os.Exit(1)
	}

	// Start the server.
	router := NewBlockchainRouter(client)
	err = http.ListenAndServe(port, router)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Oasis Rosetta Gateway server exited with error: %v\n", err)
		os.Exit(1)
	}
}
