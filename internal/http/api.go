/*
 * SPDX-FileCopyrightText: (C) 2023 Intel Corporation
 * SPDX-License-Identifier: LicenseRef-Intel
 */

package http

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

// API implements all endpoints.
type API struct {
	Version string

	// Dependency injection happens here by composing interfaces
}

// Root returns a chi.Router with all routes.
func (api *API) Root() chi.Router {
	r := chi.NewRouter()

	r.Get("/version", func(w http.ResponseWriter, r *http.Request) {
		version := struct {
			Version string `json:"version"`
		}{
			Version: api.Version,
		}

		w.Header().Add("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(version); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	})

	r.Get("/clusters", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotImplemented)
	})

	r.Get("/nodes", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotImplemented)
	})

	r.Get("/organizations", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotImplemented)
	})

	return r
}
