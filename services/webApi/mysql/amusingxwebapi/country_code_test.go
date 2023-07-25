package amusingxwebapi

import (
	"amusingx.fit/amusingx/services/webApi/conf"
	"context"
	"testing"
)

func TestQueryCountryCodes(t *testing.T) {
	conf.Mock()
	Mock()

	countryCodeList, err := QueryCountryCodes(context.Background())
	if err != nil {
		t.Fatalf("some error: %s", err)
	}

	t.Logf("country code list: %v", countryCodeList)
}
