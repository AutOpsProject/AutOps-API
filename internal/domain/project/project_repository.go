package project

import "github.com/AutOpsProject/AutOps-API/internal/domain/common"

type ProjectRepository interface {
	Create(project *Project) error
	Update(project *Project) error
	Delete(projectId common.Identifier) error

	FindById(id common.Identifier) (*Project, error)
	FindAll(offset int, limit int) ([]*Project, error)

	FindWithAllTags(tags []*common.Tag, offset int, limit int) ([]*Project, error)
	FindWithAnyTags(tags []*common.Tag, offset int, limit int) ([]*Project, error)
}
