package telemetrymgr

import (
	"context"
	"testing"
)

func TestStartService(t *testing.T) {

	svc, err := New(context.WithValue(context.Background(), configID, "telemetrymgr.yml"))

	if err != nil || (err == nil && svc == nil) {
		t.Errorf("FAILED. Error instantiating service, svc=%s, err=%s.", svc, err)
	} else {
		t.Logf("PASSED. Expected nil, got %s.\n", err)
	}
}
