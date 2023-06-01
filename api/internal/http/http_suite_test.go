/*
 * SPDX-FileCopyrightText: (C) 2023 Intel Corporation
 * SPDX-License-Identifier: LicenseRef-Intel
 */

package http_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestPlatform(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "HTTP Suite")
}
