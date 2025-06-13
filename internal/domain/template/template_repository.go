package template

import "github.com/AutOpsProject/AutOps-API/internal/domain/common"

type TemplateRepository interface {
	Create(template *Template) error
	Update(template *Template) error
	Delete(templateId common.Identifier) error

	FindByProject(projectId common.Identifier, offset int, limit int) ([]*Template, error)
	FindAllVersions(templateId common.Identifier, offset int, limit int) ([]*Template, error)

	FindById(templateId common.Identifier) (*Template, error)
	FindAll(offset int, limit int) ([]*Template, error)
	FindWithAllTags(tags []*common.Tag, offset int, limit int) ([]*Template, error)
	FindWithAnyTags(tags []*common.Tag, offset int, limit int) ([]*Template, error)
}
