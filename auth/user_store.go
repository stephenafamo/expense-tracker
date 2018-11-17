package auth

import (
	"context"
	"database/sql"
	
	"github.com/pkg/errors"
	"github.com/volatiletech/authboss"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
	amodels "github.com/stephenafamo/expense-tracker/auth_models"
	"github.com/stephenafamo/expense-tracker/models"
)

type UserStore struct {
	db *sql.DB
}

func NewUserStore(db *sql.DB) *UserStore {
	return &UserStore{
		db: db,
	}
}

// Load will look up the user based on the passed the PrimaryID
func (m *UserStore) Load(ctx context.Context, key string) (authboss.User, error) {
	user, err := models.Users(qm.Where("email=?", key)).One(ctx, m.db)
	if err != nil && errors.Cause(err) == sql.ErrNoRows {
		return nil, authboss.ErrUserNotFound
	}
	return amodels.NewUser(user), err
}

// Save persists the user in the database, this should never
// create a user and instead return ErrUserNotFound if the user
// does not exist.
func (m *UserStore) Save(ctx context.Context, user authboss.User) (err error) {
	theM := user.(*amodels.User)

	exists, err := models.Users(qm.Where("email=?", theM.Model.Email)).Exists(ctx, m.db)
	if !exists {
		return authboss.ErrUserNotFound
	}

	_, err = theM.Model.Update(ctx, m.db, boil.Infer())
	theM.Model.Reload(ctx, m.db)
	return
}

// New creates a blank user, it is not yet persisted in the database
// but is just for storing data
func (m *UserStore) New(ctx context.Context) (user authboss.User) {
	return amodels.NewUser()
}

// Create the user in storage, it should not overwrite a user
// and should return ErrUserFound if it currently exists.
func (m *UserStore) Create(ctx context.Context, user authboss.User) error {
	theM := user.(*amodels.User)

	exists, err := models.Users(qm.Where("email=?", theM.Model.Email)).Exists(ctx, m.db)
	if exists {
		return authboss.ErrUserFound
	}

	err = theM.Model.Insert(ctx, m.db, boil.Infer())
	return err
}

// AddRememberToken to a user
func (m *UserStore) AddRememberToken(ctx context.Context, pid, token string) (err error) {
	user, err := models.Users(qm.Where("email=?", pid)).One(ctx, m.db)
	if err != nil {
		return authboss.ErrUserNotFound
	}

	tokenModel := models.UserToken{Token: token}
	err = user.AddUserTokens(ctx, m.db, true, &tokenModel)

	return err
}

// DelRememberTokens removes all tokens for the given pid
func (m *UserStore) DelRememberTokens(ctx context.Context, pid string) error {
	user, err := models.Users(qm.Where("email=?", pid)).One(ctx, m.db)
	if err != nil {
		return authboss.ErrUserNotFound
	}

	_, err = user.UserTokens().DeleteAll(ctx, m.db)
	return err
}

// UseRememberToken finds the pid-token pair and deletes it.
// If the token could not be found return ErrTokenNotFound
func (m *UserStore) UseRememberToken(ctx context.Context, pid, token string) error {
	user, err := models.Users(qm.Where("email=?", pid)).One(ctx, m.db)
	if err != nil {
		return authboss.ErrUserNotFound
	}

	exists, err := user.UserTokens(qm.Where("token=?", token)).Exists(ctx, m.db)
	if !exists {
		return authboss.ErrTokenNotFound
	}
	if err != nil {
		return err
	}

	_, err = user.UserTokens(qm.Where("token=?", token)).DeleteAll(ctx, m.db)
	return err
}
