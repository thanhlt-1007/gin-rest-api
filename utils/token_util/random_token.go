package token_util

import (
	"github.com/thanhpk/randstr"
)

func RandomToken() string {
	return randstr.Hex(255)
}
