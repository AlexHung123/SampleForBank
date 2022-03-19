package api

import (
	"testing"

	db "github.com/alexhung/simplebank/db/sqlc"
	"github.com/alexhung/simplebank/util"
	"github.com/stretchr/testify/require"
)

func randowUser(t *testing.T) (user db.User, password string) {
	password = util.RandomString(6)
	hashedPassword, err := util.HashPassword(password)

	require.NoError(t, err)

	user = db.User{
		Username:       util.RandomOwner(),
		HashedPassword: hashedPassword,
		FullName:       util.RandomOwner(),
		Email:          util.RandomEmail(),
	}

	return
}
