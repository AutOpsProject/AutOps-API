package project

import (
	"testing"

	"github.com/AutOpsProject/AutOps-API/internal/domain/common"
	"github.com/AutOpsProject/AutOps-API/internal/domain/policy"
	"github.com/AutOpsProject/AutOps-API/internal/domain/template"
	"github.com/AutOpsProject/AutOps-API/internal/domain/workflow"
)

func TestNewProject(t *testing.T) {
	p, err := NewProject("MyProject", "Some description")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if p.GetName() != "MyProject" {
		t.Errorf("expected name 'MyProject', got '%s'", p.GetName())
	}
}

func TestAddAndRemoveTemplate(t *testing.T) {
	p, _ := NewProject("Project", "Desc")

	tmpl, _ := template.NewTemplate(p.GetIdentifier().ToString(), "Test", "Desc", common.PENDING, template.ANSIBLE, "/path/to/file.zip")
	p.AddTemplate(tmpl)

	found := p.GetTemplate(tmpl.GetIdentifier())
	if found == nil {
		t.Error("expected to find added template")
	}

	if err := p.RemoveTemplate(tmpl.GetIdentifier()); err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if p.GetTemplate(tmpl.GetIdentifier()) != nil {
		t.Error("expected template to be removed")
	}
}

func TestAddAndRemoveWorkflow(t *testing.T) {
	p, _ := NewProject("Project", "Desc")

	wf, _ := workflow.NewWorkflow(p.GetIdentifier().ToString(), "Wf", "Desc", "/path/to/file.yml")
	p.AddWorkflow(wf)

	if p.GetWorkflow(wf.GetIdentifier()) == nil {
		t.Error("expected to find added workflow")
	}

	if err := p.RemoveWorkflow(wf.GetIdentifier()); err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if p.GetWorkflow(wf.GetIdentifier()) != nil {
		t.Error("expected workflow to be removed")
	}
}

func TestAddAndRemovePolicy(t *testing.T) {
	p, _ := NewProject("Project", "Desc")

	pl, _ := policy.NewPolicy(p.GetIdentifier().ToString(), "Policy", "Description", []*policy.PolicyStatement{})
	p.AddPolicy(pl)

	if p.GetPolicy(pl.GetIdentifier()) == nil {
		t.Error("expected to find added policy")
	}

	if err := p.RemovePolicy(pl.GetIdentifier()); err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if p.GetPolicy(pl.GetIdentifier()) != nil {
		t.Error("expected policy to be removed")
	}
}
