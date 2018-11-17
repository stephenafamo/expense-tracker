package auth

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/antonlindstrom/pgstore"
	"github.com/volatiletech/authboss-clientstate"
)

func NewSessionStore(db *sql.DB, hashKey, blockKey []byte, options ...string) abclientstate.SessionStorer {

	name := "sessions"

	switch len(options) {
	case 1:
		name = options[0]
	}

	store, err := pgstore.NewPGStoreFromPool(db, hashKey, blockKey)
	if err != nil {
		fmt.Printf("err in pg session store %#v \n", err)
	}

	store.Options.Secure = false
	store.Options.HttpOnly = true
	store.Options.MaxAge = 60 * 60 * 24 * 60

	store.Cleanup(1 * time.Hour)

	return abclientstate.NewSessionStorerFromExisting(name, store)
}
