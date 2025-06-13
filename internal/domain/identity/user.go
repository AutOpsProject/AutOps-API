// Package identity defines entities related to user identities and access control within the AutOps platform.
package identity

import (
	"regexp"
	"strings"

	"github.com/AutOpsProject/AutOps-API/internal/domain/common"
	"github.com/AutOpsProject/AutOps-API/internal/domain/policy"
)

var (
	emailRegex    = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	usernameRegex = regexp.MustCompile(`^[a-zA-Z0-9_]+$`)
)

// User represents an individual user account in the AutOps platform. It includes user identity, email verification status,
// a unique username, and associated security policies through inheritance from RestrictedEntity.
type User struct {
	common.TimestampedEntity
	RestrictedEntity
	email    string
	verified bool
	username string
}

// NewUser creates a new User instance with a generated identifier and default timestamp.
// It validates the provided email and username.
func NewUser(email string, username string) (*User, error) {
	date := common.CurrentTimestamp()
	id, err := common.BuildIdentifier("autops:", "user")
	if err != nil {
		return nil, err
	}
	return ExistingUser(id.ToString(), email, false, username, []*policy.Policy{}, date, date)
}

// ExistingUser reconstructs a User from existing persisted data such as identifier, email, verification status,
// username, associated policies, and timestamps. It validates the input data before returning the user.
func ExistingUser(id string, email string, verified bool, username string, attachedPolicies []*policy.Policy, createdAt string, updatedAt string) (*User, error) {
	identifier, err := common.BuildUserIdentifier()
	if err != nil {
		return nil, err
	}
	timedEntity, err := common.ExistingTimestampedEntity(identifier.ToString(), createdAt, updatedAt)
	if err != nil {
		return nil, err
	}
	user := User{
		TimestampedEntity: *timedEntity,
		RestrictedEntity:  *ExistingRestrictedEntity(attachedPolicies),
		email:             "",
		verified:          false,
		username:          "",
	}
	err = user.SetEmail(email)
	if err != nil {
		return nil, err
	}
	err = user.SetUsername(username)
	if err != nil {
		return nil, err
	}
	user.verified = verified
	return &user, nil
}

// SetEmail updates the user's email after validating its format. It also resets the verification status.
func (u *User) SetEmail(email string) error {
	email = strings.TrimSpace(email)
	if !emailRegex.MatchString(email) {
		return ErrInvalidEmail
	}
	u.email = email
	u.verified = false
	return nil
}

// SetUsername updates the user's username after validating its format and length.
func (u *User) SetUsername(username string) error {
	username = strings.TrimSpace(username)
	if len(username) < 3 || len(username) > 30 || !usernameRegex.MatchString(username) {
		return ErrInvalidUsername
	}
	u.username = username
	return nil
}

// VerifyEmail marks the user's email as verified.
func (u *User) VerifyEmail() {
	u.verified = true
}

// IsVerified returns whether the user's email is verified.
func (u *User) IsVerified() bool {
	return u.verified
}

// GetEmail returns the user's email address.
func (u *User) GetEmail() string {
	return u.email
}

// GetUsername returns the user's username.
func (u *User) GetUsername() string {
	return u.username
}
