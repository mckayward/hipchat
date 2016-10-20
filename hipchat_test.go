package hipchat_test

import (
	"github.com/mckayward/hipchat"
	"testing"
)

const (
	test_token = "test-token"
	test_url   = "test.com"
)

func TestNewClientAssignsToken(t *testing.T) {
	actual := hipchat.NewClient(test_token).Token
	expected := test_token

	if actual != expected {
		t.Fatalf("Error: expected %s, but got %s", expected, actual)
	}
}

func TestNewClientUsesDefaultUrl(t *testing.T) {
	actual := hipchat.NewClient(test_token).BaseUrl
	expected := hipchat.DefaultBaseUrl

	if actual != expected {
		t.Fatalf("Error: expected %s, but got %s", expected, actual)
	}
}

func TestCanSetUrlOfClient(t *testing.T) {
	c := hipchat.NewClient(test_token)
	c.BaseUrl = test_url

	actual := c.BaseUrl
	expected := test_url

	if actual != expected {
		t.Fatalf("Error: expected %s, but got %s", expected, actual)
	}
}
