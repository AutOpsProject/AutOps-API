package policy

import (
	"strings"

	"github.com/AutOpsProject/AutOps-API/internal/domain/common"
)

// Represents the effect of a policy (Allow, Deny, or Unspecified).
type PolicyEffect int

const (
	ALLOW PolicyEffect = iota
	DENY
	UNSPECIFIED
)

// Converts a PolicyEffect to a human-readable string.
func (e PolicyEffect) ToString() (string, error) {
	switch e {
	case ALLOW:
		return "Allow", nil
	case DENY:
		return "Deny", nil
	default:
		return "", ErrInvalidPolicyEffect
	}
}

// ParsePolicyEffect converts a string into a PolicyEffect.
func ParsePolicyEffect(str string) (PolicyEffect, error) {
	str = strings.ToLower(str)
	switch str {
	case "allow":
		return ALLOW, nil
	case "deny":
		return DENY, nil
	default:
		return -1, ErrInvalidPolicyEffect
	}
}

// PolicyStatement represents a statement within a policy.
// It binds a set of resource identifiers, actions, and an effect (Allow or Deny).
type PolicyStatement struct {
	resourceIdentifiers *common.List[*common.Identifier]
	actions             *common.List[PolicyAction]
	effect              PolicyEffect
}

type PolicyStatementComparator struct{}

func (PolicyStatementComparator) Compare(a, b *PolicyStatement) int {
	a_str, _ := a.effect.ToString()
	b_str, _ := b.effect.ToString()
	return strings.Compare(a_str, b_str)
}

// NewPolicyStatement creates a new policy statement.
func NewPolicyStatement(effect PolicyEffect, resources []*common.Identifier, actions []PolicyAction) (*PolicyStatement, error) {
	_, err := effect.ToString()
	if err != nil {
		return nil, err
	}
	return &PolicyStatement{
		effect:              effect,
		resourceIdentifiers: common.NewList(common.IdentifierComparator{}, resources),
		actions:             common.NewList(PolicyActionComparator{}, actions),
	}, nil
}

// ListResources returns the resource identifiers targeted by this statement.
func (p *PolicyStatement) ListResources() []*common.Identifier {
	return p.resourceIdentifiers.Items()
}

// ListActions returns the actions covered by this statement.
func (p *PolicyStatement) ListActions() []PolicyAction {
	return p.actions.Items()
}

// GetEffect returns the effect (Allow or Deny) of this statement.
func (p *PolicyStatement) GetEffect() PolicyEffect {
	return p.effect
}

// GetPermission determines the applicable effect for a given resource and action.
// Returns UNSPECIFIED if the resource or action is not covered by the statement.
func (p *PolicyStatement) GetPermission(resourceIdentifier *common.Identifier, action PolicyAction) PolicyEffect {
	if !p.resourceIdentifiers.Contains(resourceIdentifier) || !p.actions.Contains(action) {
		return UNSPECIFIED
	}
	return p.GetEffect()
}
