package identity_test

import (
	"testing"

	"github.com/AutOpsProject/AutOps-API/internal/domain/identity"
)

func TestNewUser_Success(t *testing.T) {
	user, err := identity.NewUser("user@example.com", "valid_user")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if user == nil {
		t.Fatal("expected user, got nil")
	}
	if user.GetEmail() != "user@example.com" {
		t.Errorf("expected email to be 'user@example.com', got %s", user.GetEmail())
	}
	if user.IsVerified() {
		t.Error("new user should not be verified")
	}
	if user.GetUsername() != "valid_user" {
		t.Errorf("expected username to be 'valid_user', got %s", user.GetUsername())
	}
}

func TestNewUser_InvalidEmail(t *testing.T) {
	_, err := identity.NewUser("invalid-email", "username")
	if err != identity.ErrInvalidEmail {
		t.Errorf("expected ErrInvalidEmail, got %v", err)
	}
}

func TestNewUser_InvalidUsername(t *testing.T) {
	_, err := identity.NewUser("user@example.com", "in$valid")
	if err != identity.ErrInvalidUsername {
		t.Errorf("expected ErrInvalidUsername, got %v", err)
	}
}

func TestSetEmail(t *testing.T) {
	user := createTestUser(t)
	err := user.SetEmail("new@example.com")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if user.GetEmail() != "new@example.com" {
		t.Errorf("email not set correctly")
	}
	if user.IsVerified() {
		t.Errorf("expected verified to be false after changing email")
	}
}

func TestSetEmail_Invalid(t *testing.T) {
	user := createTestUser(t)
	err := user.SetEmail("notanemail")
	if err != identity.ErrInvalidEmail {
		t.Errorf("expected ErrInvalidEmail, got %v", err)
	}
}

func TestSetUsername(t *testing.T) {
	user := createTestUser(t)
	err := user.SetUsername("new_username")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if user.GetUsername() != "new_username" {
		t.Errorf("username not updated correctly")
	}
}

func TestSetUsername_Invalid(t *testing.T) {
	user := createTestUser(t)
	err := user.SetUsername("no")
	if err != identity.ErrInvalidUsername {
		t.Errorf("expected ErrInvalidUsername for short name")
	}
	err = user.SetUsername("thisusernameiswaytoolongtobevalid")
	if err != identity.ErrInvalidUsername {
		t.Errorf("expected ErrInvalidUsername for long name")
	}
	err = user.SetUsername("bad*char")
	if err != identity.ErrInvalidUsername {
		t.Errorf("expected ErrInvalidUsername for bad characters")
	}
}

func TestVerifyEmail(t *testing.T) {
	user := createTestUser(t)
	user.VerifyEmail()
	if !user.IsVerified() {
		t.Error("expected verified to be true after calling VerifyEmail")
	}
}

// helper
func createTestUser(t *testing.T) *identity.User {
	t.Helper()
	user, err := identity.NewUser("user@example.com", "validuser")
	if err != nil {
		t.Fatalf("failed to create test user: %v", err)
	}
	return user
}
