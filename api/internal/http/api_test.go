/*
 * SPDX-FileCopyrightText: (C) 2023 Intel Corporation
 * SPDX-License-Identifier: LicenseRef-Intel
 */

package http_test

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	net_http "net/http"
	"net/http/httptest"
	"path/filepath"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/intel-innersource/infrastructure.edge.iaas.platform/internal/http"
)

var _ = Describe("REST API", func() {
	It("Validates the API Schema", func() {
		schemaPath, err := filepath.Abs("../../schemas/openapi.yaml")
		Expect(err).NotTo(HaveOccurred())

		schema, err := openapi3.NewLoader().LoadFromFile(schemaPath)
		Expect(err).NotTo(HaveOccurred())

		Expect(schema.Validate(context.Background())).To(Succeed())
	})

	It("Has routes defined for all schema paths", func() {
		schemaPath, err := filepath.Abs("../../schemas/openapi.yaml")
		Expect(err).NotTo(HaveOccurred())

		schema, err := openapi3.NewLoader().LoadFromFile(schemaPath)
		Expect(err).NotTo(HaveOccurred())

		router := (&http.API{}).Root()

		var errors []error
		for path, item := range schema.Paths {
			for method := range item.Operations() {
				if !router.Match(chi.NewRouteContext(), method, path) {
					errors = append(errors, fmt.Errorf("no route defined for %s %s", method, path))
				}
			}
		}
		Expect(errors).To(BeEmpty())
	})

	Context("With an unauthenticated client", func() {
		It("Returns the semantic version", func() {
			version := "1.2.3"
			api := &http.API{
				Version: version,
			}

			mux := chi.NewRouter()
			mux.Use(middleware.Logger)
			mux.Use(middleware.Recoverer)
			mux.Mount("/v1", api.Root())
			s := httptest.NewServer(mux)
			defer s.Close()
			cli := s.Client()

			resp, err := cli.Get(s.URL + "/v1/version")
			Expect(err).NotTo(HaveOccurred())
			Expect(resp.StatusCode).To(Equal(net_http.StatusOK))

			// Check content against expected version
			body, err := ioutil.ReadAll(resp.Body)
			Expect(err).NotTo(HaveOccurred())

			ver := struct {
				Version string `json:"version"`
			}{}
			err = json.Unmarshal(body, &ver)
			Expect(err).NotTo(HaveOccurred())
			Expect(ver.Version).To(Equal(version))
		})
	})
})
