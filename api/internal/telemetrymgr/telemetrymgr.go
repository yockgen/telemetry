/*
 * SPDX-FileCopyrightText: (C) 2023 Intel Corporation
 * SPDX-License-Identifier: LicenseRef-Intel
 */
package telemetrymgr

import (
	"context"
	"fmt"
	"log"
	"net"
	sync "sync"

	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// Configuration defaults
const defaultPort = "3000"
const portName = "Port"
const addrName = "Address"
const defaultAddr = "localhost"
const configID = "telemetrymgr_config"

// Service interface defines the operations supported by the Telemetry Service
type TelemetrySvc interface {
	Start(ctx context.Context) error
	Stop()
}

// Represents service instance
type telemetrysvc struct {
	addr                                        string // Server string (addr:port)
	UnimplementedTelemetryServiceExternalServer        // gRPC required struct
}

var (
	instance telemetrysvc
	once     sync.Once
)

func (s *telemetrysvc) GetGroups(ctx context.Context, in *GetGroupsRequest) (*GetGroupsResponse, error) {

	log.Printf("Received request: %v", in.ProtoReflect().Descriptor().FullName())
	return getGroups()
}

func (s *telemetrysvc) GetGroup(ctx context.Context, in *GetGroupRequest) (*GetGroupResponse, error) {

	log.Printf("yockgen2 Received request: %v", in.ProtoReflect().Descriptor().FullName())
	log.Printf("idx: %d", in.Index)
	return getGroup(in)
}

// Creates a new singelton instance of the Telemeetry service
func New(ctx context.Context) (TelemetrySvc, error) {

	fmt.Println("yockgen....from telemetrymgr")

	// Atomic singleton
	once.Do(func() {
		cfg := ctx.Value(configID).(string)

		viper.SetDefault(portName, defaultPort)
		viper.SetDefault(addrName, defaultAddr)
		viper.SetConfigName(cfg)
		viper.SetConfigType("yaml")
		viper.AddConfigPath(".")
		viper.AddConfigPath("../../internal/config")

		if err := viper.ReadInConfig(); err != nil {
			if _, ok := err.(viper.ConfigFileNotFoundError); ok {
				log.Printf("Configuration file not found, using defaults")
			} else {
				log.Fatalf("Configuration file error: %d", err)
			}
		}
		instance.addr = fmt.Sprintf("%s:%s", viper.GetString(addrName), viper.GetString(portName))

	})
	return instance, nil
}

// Starts the Telemetry server and regsters gRPC services
func (tsvc telemetrysvc) Start(ctx context.Context) error {

	fmt.Println("Bind Address : Port = ", instance.addr)

	lis, err := net.Listen("tcp", instance.addr)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	RegisterTelemetryServiceExternalServer(s, &telemetrysvc{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %s", err)
	}
	return nil
}

func (tsvc telemetrysvc) Stop() {
}
