package auth

import (
	"github.com/volatiletech/authboss-clientstate"
)

func NewCookieStore(hashKey, blockKey []byte) abclientstate.CookieStorer {
	return abclientstate.NewCookieStorer(hashKey, blockKey)
}
