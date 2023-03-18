package main

import (
	"log"
	"os"

	"github.com/sekibomazic/hex-arch-go-grpc/internal/adapters/app/api"
	"github.com/sekibomazic/hex-arch-go-grpc/internal/adapters/core/arithmetic"
	gRPC "github.com/sekibomazic/hex-arch-go-grpc/internal/adapters/framework/left/grpc"
	"github.com/sekibomazic/hex-arch-go-grpc/internal/adapters/framework/right/db"
	"github.com/sekibomazic/hex-arch-go-grpc/internal/ports"
)

func main() {
	var err error

	// ports
	var dbaseAdapter ports.DbPort
	var core ports.ArithmeticPort
	var appAdapter ports.APIPort
	var gRPCAdapter ports.GRPCPort

	dbaseDriver := os.Getenv("DB_DRIVER")
	dsourceName := os.Getenv("DS_NAME")

	dbaseAdapter, err = db.NewAdapter(dbaseDriver, dsourceName)
	if err != nil {
		log.Fatalf("failed to initiate db connection %v", err)
	}

	defer dbaseAdapter.CloseDbConnection()

	core = arithmetic.NewAdapter()
	appAdapter = api.NewAdapter(dbaseAdapter, core)

	gRPCAdapter = gRPC.NewAdapter(appAdapter)
	gRPCAdapter.Run()
}
