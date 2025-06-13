package policy

import (
	"strings"

	"github.com/AutOpsProject/AutOps-API/internal/domain/common"
)

// Policy represents a security policy defining permissions over resources.
type Policy struct {
	common.NamedEntity
	common.TaggedEntity
	statements *common.List[*PolicyStatement]
}

type PolicyComparator struct{}

func (PolicyComparator) Compare(p1, p2 *Policy) int {
	return strings.Compare(p1.GetIdentifier().ToString(), p2.GetIdentifier().ToString())
}

// NewPolicy creates a new Policy instance with the given project identifier, name,
// description, and policy statements. It automatically generates an identifier and sets timestamps.
func NewPolicy(projectIdentifier string, name string, description string, statements []*PolicyStatement) (*Policy, error) {
	date := common.CurrentTimestamp()
	identifier, err := common.BuildPolicyIdentifier(projectIdentifier)
	if err != nil {
		return nil, err
	}
	return ExistingPolicy(identifier.ToString(), name, description, date, date, statements)
}

// ExistingPolicy reconstructs an existing Policy from stored data, typically retrieved from a database.
func ExistingPolicy(policyIdentifier string, name string, description string, createdAt string, updatedAt string, statements []*PolicyStatement) (*Policy, error) {
	namedEntity, err := common.ExistingNamedEntity(policyIdentifier, name, description, createdAt, updatedAt)
	if err != nil {
		return nil, err
	}
	return &Policy{
		NamedEntity:  *namedEntity,
		TaggedEntity: *common.NewTaggedEntity(),
		statements:   common.NewList(common.Comparator[*PolicyStatement](PolicyStatementComparator{}), statements),
	}, nil
}

// ListStatements returns all the statements attached to the policy.
func (p *Policy) ListStatements() []*PolicyStatement {
	return p.statements.Items()
}

// GetPermission determines the policy effect (ALLOW, DENY, or UNSPECIFIED) for a given action
// on a specified resource identifier. DENY takes precedence over ALLOW.
func (p *Policy) GetPermission(resourceIdentifier *common.Identifier, action PolicyAction) PolicyEffect {
	allowed := false
	for _, p := range p.statements.Items() {
		effect := p.GetPermission(resourceIdentifier, action)
		if effect == DENY {
			return DENY
		} else if effect == ALLOW {
			allowed = true
		}
	}
	if allowed {
		return ALLOW
	}
	return UNSPECIFIED
}
