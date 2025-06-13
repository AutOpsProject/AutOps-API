package policy

import "github.com/AutOpsProject/AutOps-API/internal/domain/common"

type PolicyRepository interface {
	Create(policy *Policy) error
	Update(policy *Policy) error
	Delete(policyId common.Identifier) error

	FindById(policyId common.Identifier) (*Policy, error)
	FindAll(offset int, limit int) ([]*Policy, error)
	FindByEntity(entityId common.Identifier, offset int, limit int) ([]*Policy, error)
	AttachToEntity(policyId common.Identifier, entityId common.Identifier) error
	DetachFromEntity(policyId common.Identifier, entityId common.Identifier) error

	FindWithAllTags(tags []*common.Tag, offset int, limit int) ([]*Policy, error)
	FindWithAnyTags(tags []*common.Tag, offset int, limit int) ([]*Policy, error)
}
