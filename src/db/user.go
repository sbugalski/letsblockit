package db

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/volatiletech/authboss/v3"
)

type UserStore struct {
	db Store
}

func NewUserStore(db Store) *UserStore {
	return &UserStore{db: db}
}

func (s UserStore) New(_ context.Context) authboss.User {
	return &User{}
}

func (s UserStore) Load(ctx context.Context, email string) (authboss.User, error) {
	u, err := s.db.GetUserByEmail(ctx, email)
	switch err {
	case nil:
		return &u, nil
	case NotFound:
		return nil, authboss.ErrUserNotFound
	default:
		return nil, err
	}
}

func (s UserStore) Save(ctx context.Context, user authboss.User) error {
	u, ok := user.(*User)
	if !ok || u.Email == "" {
		return errors.New("invalid user object")
	}
	err := s.db.CreateUser(ctx, CreateUserParams{
		Email:            u.Email,
		EmailConfirmed:   u.EmailConfirmed,
		PasswordHash:     u.PasswordHash,
		ConfirmSelector:  u.ConfirmSelector,
		ConfirmVerifier:  u.ConfirmVerifier,
		RecoverySelector: u.RecoverySelector,
		RecoveryVerifier: u.RecoveryVerifier,
		RecoveryExpiry:   u.RecoveryExpiry,
	})
	switch err {
	case nil:
		return nil
	case NotFound:
		return authboss.ErrUserNotFound
	default:
		return err
	}
}

func (s UserStore) Create(ctx context.Context, user authboss.User) error {
	u, ok := user.(*User)
	if !ok || u.Email == "" {
		return errors.New("invalid user object")
	}
	return s.db.RunCtxTx(ctx, func(ctx context.Context, q Querier) error {
		if _, e := q.GetUserByEmail(ctx, u.Email); e != NotFound {
			return authboss.ErrUserFound
		}
		return q.CreateUser(ctx, CreateUserParams{
			Email:            u.Email,
			EmailConfirmed:   u.EmailConfirmed,
			PasswordHash:     u.PasswordHash,
			ConfirmSelector:  u.ConfirmSelector,
			ConfirmVerifier:  u.ConfirmVerifier,
			RecoverySelector: u.RecoverySelector,
			RecoveryVerifier: u.RecoveryVerifier,
			RecoveryExpiry:   u.RecoveryExpiry,
		})
	})
}

func (u *User) GetPID() string {
	return u.Email
}

func (u *User) PutPID(pid string) {
	u.Email = pid
}

func (u *User) GetEmail() string {
	return u.Email
}

func (u *User) GetPassword() string {
	return u.PasswordHash.String
}

func (u *User) GetConfirmed() bool {
	return u.EmailConfirmed
}

func (u *User) GetConfirmSelector() string {
	return u.ConfirmSelector.String
}

func (u *User) GetConfirmVerifier() string {
	return u.ConfirmVerifier.String
}

func (u *User) GetRecoverSelector() string {
	return u.RecoverySelector.String
}

func (u *User) GetRecoverVerifier() string {
	return u.RecoveryVerifier.String
}

func (u *User) GetRecoverExpiry() time.Time {
	return u.RecoveryExpiry.Time
}

func (u *User) PutEmail(email string) {
	u.Email = email
}

func (u *User) PutPassword(password string) {
	u.PasswordHash = toNullString(password)
}

func (u *User) PutConfirmed(confirmed bool) {
	u.EmailConfirmed = confirmed
}

func (u *User) PutConfirmSelector(selector string) {
	u.ConfirmSelector = toNullString(selector)
}

func (u *User) PutConfirmVerifier(verifier string) {
	u.ConfirmVerifier = toNullString(verifier)
}

func (u *User) PutRecoverSelector(selector string) {
	u.RecoverySelector = toNullString(selector)
}

func (u *User) PutRecoverVerifier(verifier string) {
	u.RecoveryVerifier = toNullString(verifier)
}

func (u *User) PutRecoverExpiry(expiry time.Time) {
	if expiry.IsZero() {
		u.RecoveryExpiry = sql.NullTime{}
	} else {
		u.RecoveryExpiry = sql.NullTime{
			Time:  expiry,
			Valid: true,
		}
	}
}

func toNullString(input string) sql.NullString {
	if input == "" {
		return sql.NullString{}
	} else {
		return sql.NullString{
			String: input,
			Valid:  true,
		}
	}
}
