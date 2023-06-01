/*
 * SPDX-FileCopyrightText: (C) 2023 Intel Corporation
 * SPDX-License-Identifier: LicenseRef-Intel
 */

package main

import (
	net_http "net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/intel-innersource/infrastructure.edge.iaas.platform/internal/http"
)

// Version is set at build time.
var Version string

func main() {
	mux := chi.NewRouter()
	mux.Use(middleware.Logger)
	mux.Use(middleware.Recoverer)
	mux.Mount("/v1", (&http.API{
		Version: Version,
	}).Root())

	srv := &net_http.Server{
		Addr:           ":3000",
		Handler:        mux,
		ReadTimeout:    15 * time.Second,
		WriteTimeout:   15 * time.Second,
		MaxHeaderBytes: net_http.DefaultMaxHeaderBytes,
	}

	_ = srv.ListenAndServe()
}
