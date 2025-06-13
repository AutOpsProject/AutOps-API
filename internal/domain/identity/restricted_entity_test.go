package identity_test

import (
	"testing"

	"github.com/AutOpsProject/AutOps-API/internal/domain/common"
	"github.com/AutOpsProject/AutOps-API/internal/domain/identity"
	"github.com/AutOpsProject/AutOps-API/internal/domain/policy"
)

func makeTestPolicy(id string) *policy.Policy {
	identifier, _ := common.NewIdentifier(id)
	p, _ := policy.NewPolicy(identifier.ToString(), "test", "test", []*policy.PolicyStatement{})
	return p
}

func TestNewRestrictedEntity(t *testing.T) {
	entity := identity.NewRestrictedEntity()
	if len(entity.ListAttachedPolicies()) != 0 {
		t.Errorf("expected no policies, got %d", len(entity.ListAttachedPolicies()))
	}
}

func TestAttachPolicy(t *testing.T) {
	entity := identity.NewRestrictedEntity()
	p := makeTestPolicy("autops::project:1234567890:policy:abcdefghij")
	entity.AttachPolicy(p)

	policies := entity.ListAttachedPolicies()
	if len(policies) != 1 || policies[0] != p {
		t.Errorf("AttachPolicy failed, got %+v", policies)
	}
}

func TestGetAttachedPolicy(t *testing.T) {
	p := makeTestPolicy("autops::project:1234567890:policy:abcdefghij")
	entity := identity.ExistingRestrictedEntity([]*policy.Policy{p})

	got := entity.GetAttachedPolicy(p.GetIdentifier())
	if got == nil || got.GetIdentifier().ToString() != p.GetIdentifier().ToString() {
		t.Errorf("expected to find policy %s, got %v", p.GetIdentifier(), got)
	}
}

func TestDetachPolicySuccess(t *testing.T) {
	p := makeTestPolicy("autops::project:1234567890:policy:abcdefghij")
	entity := identity.ExistingRestrictedEntity([]*policy.Policy{p})

	err := entity.DetachPolicy(p.GetIdentifier())
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	if len(entity.ListAttachedPolicies()) != 0 {
		t.Errorf("expected policy to be detached")
	}
}

func TestDetachPolicyNotFound(t *testing.T) {
	entity := identity.NewRestrictedEntity()
	id, _ := common.NewIdentifier("autops::project:1234567890:policy:notfounddd")
	err := entity.DetachPolicy(id)

	if err != identity.ErrAttachedPolicyNotFound {
		t.Errorf("expected ErrAttachedPolicyNotFound, got %v", err)
	}
}
