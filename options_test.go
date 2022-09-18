package goshopify

import (
	"fmt"
	"net/http"
	"testing"
	"time"
)

func TestWithVersion(t *testing.T) {
	c := NewClient(app, testShopName, testToken, WithVersion(testApiVersion))
	expected := fmt.Sprintf("admin/api/%s", testApiVersion)
	if c.pathPrefix != expected {
		t.Errorf("WithVersion client.pathPrefix = %s, expected %s", c.pathPrefix, expected)
	}
}

func TestWithVersionNoVersion(t *testing.T) {
	c := NewClient(app, testShopName, testToken, WithVersion(""))
	expected := "admin"
	if c.pathPrefix != expected {
		t.Errorf("WithVersion client.pathPrefix = %s, expected %s", c.pathPrefix, expected)
	}
}

func TestWithoutVersionInInitiation(t *testing.T) {
	c := NewClient(app, testShopName, testToken)
	expected := "admin"
	if c.pathPrefix != expected {
		t.Errorf("WithVersion client.pathPrefix = %s, expected %s", c.pathPrefix, expected)
	}
}

func TestWithVersionInvalidVersion(t *testing.T) {
	c := NewClient(app, testShopName, testToken, WithVersion("9999-99b"))
	expected := "admin"
	if c.pathPrefix != expected {
		t.Errorf("WithVersion client.pathPrefix = %s, expected %s", c.pathPrefix, expected)
	}
}

func TestWithRetry(t *testing.T) {
	c := NewClient(app, testShopName, testToken, WithRetry(5))
	expected := 5
	if c.retries != expected {
		t.Errorf("WithRetry client.retries = %d, expected %d", c.retries, expected)
	}
}

func TestWithLogger(t *testing.T) {
	logger := &LeveledLogger{Level: LevelDebug}
	c := NewClient(app, testShopName, testToken, WithLogger(logger))

	if c.log != logger {
		t.Errorf("WithLogger expected logs to match %v != %v", c.log, logger)
	}
}

func TestWithUnstableVersion(t *testing.T) {
	c := NewClient(app, testShopName, testToken, WithVersion(UnstableApiVersion))
	expected := fmt.Sprintf("admin/api/%s", UnstableApiVersion)
	if c.pathPrefix != expected {
		t.Errorf("WithVersion client.pathPrefix = %s, expected %s", c.pathPrefix, expected)
	}
}

func TestWithHTTPClient(t *testing.T) {
	c := NewClient(app, testShopName, testToken, WithHTTPClient(&http.Client{Timeout: 30 * time.Second}))
	expected := 30 * time.Second

	if c.Client.Timeout.String() != expected.String() {
		t.Errorf("WithVersion client.Client = %s, expected %s", c.Client.Timeout, expected)
	}
}
