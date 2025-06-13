package project

import (
	"strings"

	"github.com/AutOpsProject/AutOps-API/internal/domain/common"
	"github.com/AutOpsProject/AutOps-API/internal/domain/policy"
	"github.com/AutOpsProject/AutOps-API/internal/domain/template"
	"github.com/AutOpsProject/AutOps-API/internal/domain/workflow"
)

// Project represents an AutOps project, grouping templates, workflows, and policies
// under a single logical unit.
type Project struct {
	common.NamedEntity
	common.TaggedEntity
	templates *common.List[*template.Template]
	workflows *common.List[*workflow.Workflow]
	policies  *common.List[*policy.Policy]
}

// NewProject creates a new Project with a generated identifier and current timestamps.
func NewProject(name string, description string) (*Project, error) {
	id, err := common.BuildProjectIdentifier()
	if err != nil {
		return nil, err
	}
	date := common.CurrentTimestamp()
	return ExistingProject(
		id.ToString(),
		name,
		description,
		date,
		date,
		[]*template.Template{},
		[]*workflow.Workflow{},
		[]*policy.Policy{},
	)
}

// ExistingProject rebuilds a Project from persisted data, including associated templates, workflows, and policies.
func ExistingProject(id string, name string, description string, createdAt string, updatedAt string, templates []*template.Template, workflows []*workflow.Workflow, policies []*policy.Policy) (*Project, error) {
	namedEntity, err := common.ExistingNamedEntity(id, name, description, createdAt, updatedAt)
	if err != nil {
		return nil, err
	}
	return &Project{
		NamedEntity:  *namedEntity,
		TaggedEntity: *common.NewTaggedEntity(),
		templates:    common.NewList(template.TemplateComparator{}, templates),
		workflows:    common.NewList(workflow.WorkflowComparator{}, workflows),
		policies:     common.NewList(policy.PolicyComparator{}, policies),
	}, nil
}

// ListTemplates returns all templates associated with the project.
func (p *Project) ListTemplates() []*template.Template {
	return p.templates.Items()
}

// GetTemplate returns a template associated with the given identifier, or nil if not found.
func (p *Project) GetTemplate(id *common.Identifier) *template.Template {
	template, _ := p.templates.SelectOne((func(t *template.Template) bool {
		return strings.Compare(id.ToString(), t.GetIdentifier().ToString()) == 0
	}))
	return template
}

// AddTemplate adds a template to the project.
func (p *Project) AddTemplate(template *template.Template) {
	p.templates.Append(template)
}

// RemoveTemplate removes a template from the project by its identifier.
// Returns an error if the template is not found.
func (p *Project) RemoveTemplate(id *common.Identifier) error {
	template := p.GetTemplate(id)
	if template == nil {
		return ErrTemplateNotFound
	}
	p.templates.Remove(template)
	return nil
}

// ListWorkflows returns all workflows associated with the project.
func (p *Project) ListWorkflows() []*workflow.Workflow {
	return p.workflows.Items()
}

// GetWorkflow returns a workflow associated with the given identifier, or nil if not found.
func (p *Project) GetWorkflow(id *common.Identifier) *workflow.Workflow {
	workflow, _ := p.workflows.SelectOne((func(t *workflow.Workflow) bool {
		return strings.Compare(id.ToString(), t.GetIdentifier().ToString()) == 0
	}))
	return workflow
}

// AddWorkflow adds a workflow to the project.
func (p *Project) AddWorkflow(workflow *workflow.Workflow) {
	p.workflows.Append(workflow)
}

// RemoveWorkflow removes a workflow from the project by its identifier.
// Returns an error if the workflow is not found.
func (p *Project) RemoveWorkflow(id *common.Identifier) error {
	workflow := p.GetWorkflow(id)
	if workflow == nil {
		return ErrWorkflowNotFound
	}
	p.workflows.Remove(workflow)
	return nil
}

// ListPolicies returns all policies associated with the project.
func (p *Project) ListPolicies() []*policy.Policy {
	return p.policies.Items()
}

// GetPolicy returns a policy associated with the given identifier, or nil if not found.
func (p *Project) GetPolicy(id *common.Identifier) *policy.Policy {
	policy, _ := p.policies.SelectOne((func(t *policy.Policy) bool {
		return strings.Compare(id.ToString(), t.GetIdentifier().ToString()) == 0
	}))
	return policy
}

// AddPolicy adds a policy to the project.
func (p *Project) AddPolicy(policy *policy.Policy) {
	p.policies.Append(policy)
}

// RemovePolicy removes a policy from the project by its identifier.
// Returns an error if the policy is not found.
func (p *Project) RemovePolicy(id *common.Identifier) error {
	policy := p.GetPolicy(id)
	if policy == nil {
		return ErrPolicyNotFound
	}
	p.policies.Remove(policy)
	return nil
}
