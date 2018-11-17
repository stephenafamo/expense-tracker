package models

import (
	"time"

	"github.com/volatiletech/null"
	"github.com/stephenafamo/expense-tracker/models"
)

type User struct {
	Model     *models.User
	arbitrary map[string]string
}

func NewUser(user ...*models.User) *User {
	Model := models.User{}
	modelPtr := &Model
	if len(user) > 0 {
		modelPtr = user[0]
	}

	arbitrary := make(map[string]string)

	return &User{
		Model:     modelPtr,
		arbitrary: arbitrary,
	}
}

func (m User) GetPID() (pid string) {
	return m.Model.Email
}

func (m User) PutPID(pid string) {
	m.Model.Email = pid
}

func (m *User) GetPassword() (password string) {
	return m.Model.Password.String
}

func (m *User) PutPassword(password string) {
	m.Model.Password = null.StringFrom(password)
}

func (m *User) GetEmail() (email string) {
	return m.Model.Email
}

func (m *User) PutEmail(email string) {
	m.Model.Email = email
}

func (m *User) GetAttemptCount() (attempts int) {
	return m.Model.Attempts.Int
}

func (m *User) GetLastAttempt() (last time.Time) {
	return m.Model.LastAttemptTime.Time
}

func (m *User) GetLocked() (locked time.Time) {
	return m.Model.Locked.Time
}

func (m *User) PutAttemptCount(attempts int) {
	m.Model.Attempts = null.IntFrom(attempts)
}

func (m *User) PutLastAttempt(last time.Time) {
	m.Model.LastAttemptTime = null.TimeFrom(last)
}

func (m *User) PutLocked(locked time.Time) {
	m.Model.Locked = null.TimeFrom(locked)
}

func (m *User) GetArbitrary() (arbitrary map[string]string) {
	return m.arbitrary
}

func (m *User) PutArbitrary(arbitrary map[string]string) {
	for k, v := range arbitrary {
		m.arbitrary[k] = v
	}
}
