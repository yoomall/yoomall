package jutuike_test

import (
	"testing"

	"lazyfury.github.com/yoomall-server/libs/jutuike"
)

func TestJutuike(t *testing.T) {

	_, err := jutuike.NewJtkClient(&jutuike.JtkConfig{
		PubId:  "pub_id",
		ApiKey: "api_key",
		JtkUrl: "https://api.jutuike.com",
	})
	if err != nil {
		t.Error(err)
	}

}
