package dtk_test

import (
	"testing"

	"yoomall/src/libs/dtk"
)

// TestNewDtkWithoutAppKey tests that calling NewDtkClient with an empty
// DtkConfig returns an error.
func TestNewDtkWithoutAppKey(t *testing.T) {

	_, err := dtk.NewDtkClient(&dtk.DtkConfig{})

	if err == nil {
		t.Error("expected error")
	}

}

func TestNewDtkWithoutAppSecret(t *testing.T) {
	_, err := dtk.NewDtkClient(&dtk.DtkConfig{
		AppKey: "app_key",
	})

	if err == nil {
		t.Error("expected error")
	}
}

func TestNewDtkWithoutAppUrl(t *testing.T) {
	_, err := dtk.NewDtkClient(&dtk.DtkConfig{
		AppKey:    "app_key",
		AppSecret: "app_secret",
	})

	if err == nil {
		t.Error("expected error")
	}
}

func TestNewDtk(t *testing.T) {
	_, err := dtk.NewDtkClient(&dtk.DtkConfig{
		AppKey:    "app_key",
		AppSecret: "app_secret",
		AppUrl:    "https://api.dtkmall.com",
	})

	if err != nil {
		t.Error(err)
	}
}
