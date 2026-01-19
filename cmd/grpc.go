package cmd

import (
	"E-Wallet-wallet/helpers"
	"log"
	"net"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func ServerGRPC() {
	list, err := net.Listen("tcp", ":"+helpers.GetEnv("GRPC_PORT", ""))
	if err != nil {
		log.Fatal("failed serve to GRPC", err)
	}

	s := grpc.NewServer()

	logrus.Info("start listening GRPC on port: " + helpers.GetEnv("GRPC_PORT", ""))
	if err := s.Serve(list); err != nil {
		log.Fatal("failed serve to GRPC", err)
	}
}
