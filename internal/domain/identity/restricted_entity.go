// Package identity defines domain models related to access-controlled entities.
// These entities can have security policies attached or detached at runtime.
package identity

import (
	"strings"

	"github.com/AutOpsProject/AutOps-API/internal/domain/common"
	"github.com/AutOpsProject/AutOps-API/internal/domain/policy"
)

// RestrictedEntity represents any entity that can have security policies attached to it.
// It abstracts common behavior for access-controlled resources such as users, access keys, or roles.
type RestrictedEntity struct {
	attachedPolicies *common.List[*policy.Policy]
}

// NewRestrictedEntity creates a new RestrictedEntity with no attached policies.
func NewRestrictedEntity() *RestrictedEntity {
	return ExistingRestrictedEntity([]*policy.Policy{})
}

// ExistingRestrictedEntity creates a RestrictedEntity with an initial list of attached policies.
// It is intended to be used when reconstructing the entity from storage.
func ExistingRestrictedEntity(attachedPolicies []*policy.Policy) *RestrictedEntity {
	return &RestrictedEntity{
		attachedPolicies: common.NewList(policy.PolicyComparator{}, attachedPolicies),
	}
}

// ListAttachedPolicies returns a slice of all currently attached policies.
func (r *RestrictedEntity) ListAttachedPolicies() []*policy.Policy {
	return r.attachedPolicies.Items()
}

// GetAttachedPolicy returns the policy with the given identifier if found.
// Returns nil if no matching policy is attached.
func (r *RestrictedEntity) GetAttachedPolicy(id *common.Identifier) *policy.Policy {
	policy, _ := r.attachedPolicies.SelectOne(func(p *policy.Policy) bool {
		return strings.Compare(id.ToString(), p.GetIdentifier().ToString()) == 0
	})
	return policy
}

// DetachPolicy removes a policy identified by its identifier from the entity.
// Returns an error if no such policy is currently attached.
func (r *RestrictedEntity) DetachPolicy(id *common.Identifier) error {
	policy := r.GetAttachedPolicy(id)
	if policy == nil {
		return ErrAttachedPolicyNotFound
	}
	r.attachedPolicies.Remove(policy)
	return nil
}

// AttachPolicy appends the given policy to the entity's attached policies list.
func (r *RestrictedEntity) AttachPolicy(policy *policy.Policy) {
	r.attachedPolicies.Append(policy)
}
