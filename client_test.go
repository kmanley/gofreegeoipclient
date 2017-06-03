package gofreegeoipclient

import (
	"testing"
	"time"

	"github.com/davecgh/go-spew/spew"
)

func TestGetCountryForIP(t *testing.T) {
	client := Locator{10 * time.Second}

	if ret, err := client.Locate("92.50.115.170"); ret.CountryCode != "DE" {
		t.Errorf("Locate(%v) = %v, want %v, err=%v", "92.50.115.170", ret.CountryCode, "DE", err)
	} else {
		spew.Dump(ret)
	}
	if ret, err := client.Locate("92.50.115.a"); err == nil {
		t.Errorf("Locate(%v) = %v, want %v", "92.50.115.a", ret, nil)
	}
}
