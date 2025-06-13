package workflow

import "github.com/AutOpsProject/AutOps-API/internal/domain/common"

type WorkflowRepository interface {
	Create(workflow *Workflow) error
	Update(workflow *Workflow) error
	Delete(workflowId common.Identifier) error

	FindByProject(projectId common.Identifier, offset int, limit int) ([]*Workflow, error)
	FindAllVersions(workflowId common.Identifier, offset int, limit int) ([]*Workflow, error)

	FindById(workflowId common.Identifier) (*Workflow, error)
	FindAll(offset int, limit int) ([]*Workflow, error)
	FindWithAllTags(tags []*common.Tag, offset int, limit int) ([]*Workflow, error)
	FindWithAnyTags(tags []*common.Tag, offset int, limit int) ([]*Workflow, error)
}
