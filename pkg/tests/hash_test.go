package tests

import (
	"testing"

	"github.com/Arshdeep54/Shelflove-mvc/pkg/utils"
	"golang.org/x/crypto/bcrypt"
)

type HashData struct {
	password string
	err      error
}

var hashPayloads = []HashData{
	{
		password: "cosign",
		err:      nil,
	},
	{
		password: "this is random oasswrod whose length i guess would be more then 72 characters to get an error",
		err:      bcrypt.ErrPasswordTooLong,
	},
	{
		password: "",
		err: nil,
	},
}

func TestHashPassword(t *testing.T) {
	for _, value := range hashPayloads {
		_, err := utils.HashedPassword(value.password)
		if err != nil {
			if err != value.err {
				t.Fatal(err.Error(), value.err)
			}
		}

	}
}
