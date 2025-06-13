package template

import (
	"testing"

	"github.com/AutOpsProject/AutOps-API/internal/domain/common"
)

func TestTemplateTypeToString(t *testing.T) {
	templateType := TERRAFORM
	if templateType.ToString() != "terraform" {
		t.Errorf("expected terraform, got %s", templateType.ToString())
	}

	templateType = ANSIBLE
	if templateType.ToString() != "ansible" {
		t.Errorf("expected ansible, got %s", templateType.ToString())
	}

	templateType = PACKER
	if templateType.ToString() != "packer" {
		t.Errorf("expected packer, got %s", templateType.ToString())
	}

	templateType = OPENTOFU
	if templateType.ToString() != "opentofu" {
		t.Errorf("expected opentofu, got %s", templateType.ToString())
	}
}

func TestParseTemplateType(t *testing.T) {
	str := "terraform"
	value, err := ParseTemplateType(str)
	if err != nil {
		t.Error("expected err to be nil")
	}
	if value != TERRAFORM {
		t.Errorf("expected %d, got %d", TERRAFORM, value)
	}

	str = "Ansible"
	value, err = ParseTemplateType(str)
	if err != nil {
		t.Error("expected err to be nil")
	}
	if value != ANSIBLE {
		t.Errorf("expected %d, got %d", ANSIBLE, value)
	}

	str = "PACKER"
	value, err = ParseTemplateType(str)
	if err != nil {
		t.Error("expected err to be nil")
	}
	if value != PACKER {
		t.Errorf("expected %d, got %d", PACKER, value)
	}

	str = "OpenTofu"
	value, err = ParseTemplateType(str)
	if err != nil {
		t.Error("expected err to be nil")
	}
	if value != OPENTOFU {
		t.Errorf("expected %d, got %d", OPENTOFU, value)
	}

	str = "invalid type"
	_, err = ParseTemplateType(str)
	if err != ErrParseInvalidTemplateType {
		t.Error("expected err to be ErrParseInvalidTemplateType")
	}
}

func TestNewTemplate(t *testing.T) {
	projectIdentifier := "autops::project:ABCDEFGHIJ:template:1234567890"
	templateName := "my-first-project"
	templateDescription := "my description"
	templateStatus := common.PENDING
	templateType := TERRAFORM
	templateSource := "/path/to/file.zip"
	template, err := NewTemplate(projectIdentifier, templateName, templateDescription, templateStatus, templateType, templateSource)
	if err != nil {
		t.Error("expected err to be nil")
	} else {
		if template.GetName() != templateName {
			t.Errorf("expected template name to be %s, got %s", templateName, template.GetName())
		}
		if template.GetDescription() != templateDescription {
			t.Errorf("expected template description to be %s, got %s", templateDescription, template.GetDescription())
		}
		if template.GetSourcePath() != templateSource {
			t.Errorf("expected template source path to be %s, got %s", templateSource, template.GetSourcePath())
		}
		if template.GetTemplateType() != templateType {
			t.Errorf("expected template type to be %d, got %d", templateType, template.GetTemplateType())
		}
		if template.GetVersion() != 1 {
			t.Errorf("expected template version to be %d, got %d", 1, template.GetVersion())
		}
	}
}

func TestExistingTemplate(t *testing.T) {
	templateIdentifier := "autops::project:ABCDEFGHIJ:template:1234567890"
	templateName := "my-first-project"
	templateDescription := "my description"
	templateStatus := common.PENDING
	templateType := TERRAFORM
	templateSource := "/path/to/file.zip"
	template, err := ExistingTemplate(templateIdentifier, templateName, templateDescription, templateStatus, templateType, templateSource, 1)
	if err != nil {
		t.Error("expected err to be nil")
	} else if template.GetIdentifier().ToString() != "autops::project:ABCDEFGHIJ:template:1234567890" {
		t.Errorf("expeted template identifier to be autops::project:ABCDEFGHIJ:template:1234567890, got %s", template.GetIdentifier())
	}

	templateName = "Invalid template name"
	_, err = ExistingTemplate(templateIdentifier, templateName, templateDescription, templateStatus, templateType, templateSource, 1)
	if err != common.ErrInvalidName {
		t.Error("expected err to be ErrInvalidName")
	}

	templateName = "valid-name"
	templateSource = "invalid source path"
	_, err = ExistingTemplate(templateIdentifier, templateName, templateDescription, templateStatus, templateType, templateSource, 1)
	if err != common.ErrInvalidPathOrUrl {
		t.Error("expected err to be ErrInvalidPathOrUrl")
	}
}

func TestTemplateInputs(t *testing.T) {
	templateIdentifier := "autops::project:ABCDEFGHIJ:template:1234567890"
	templateName := "my-first-project"
	templateDescription := "my description"
	templateStatus := common.PENDING
	templateType := TERRAFORM
	templateSource := "/path/to/file.zip"
	template, err := ExistingTemplate(templateIdentifier, templateName, templateDescription, templateStatus, templateType, templateSource, 18)
	if err != nil {
		t.Error("expected err to be nil")
	} else if template.GetVersion() != 18 {
		t.Errorf("expected template version to be 18, got %d", template.GetVersion())
	}

	input1, _ := NewTemplateAttribute(templateIdentifier, "inputA", "", STRING, "value")
	input2, _ := NewTemplateAttribute(templateIdentifier, "inputB", "", NUMBER, "3.14")
	input3, _ := NewTemplateAttribute(templateIdentifier, "inputC", "", LIST, "[1, 2]")
	err = template.AddInput(input1)
	if err != nil {
		t.Error("expected err to be nil")
	}
	err = template.AddInput(input2)
	if err != nil {
		t.Error("expected err to be nil")
	}
	err = template.AddInput(input1)
	if err != ErrTemplateInputAlreadyPresent {
		t.Error("expected err to be ErrInputAlreadyPresent")
	}

	if len(template.ListInputs()) != 2 {
		t.Errorf("expected len(inputs) to be 2, got %d", len(template.ListInputs()))
	}

	err = template.RemoveInput(input1.GetIdentifier().ToString())
	if err != nil {
		t.Error("expected err to be nil")
	}
	err = template.RemoveInput(input3.GetIdentifier().ToString())
	if err != ErrTemplateInputNotFound {
		t.Error("expected err to be ErrInputNotFound")
	}

	if len(template.ListInputs()) != 1 {
		t.Errorf("expected len(inputs) to be 1, got %d", len(template.ListInputs()))
	}
}

func TestTemplateOutputs(t *testing.T) {
	templateIdentifier := "autops::project:ABCDEFGHIJ:template:1234567890"
	templateName := "my-first-project"
	templateDescription := "my description"
	templateStatus := common.PENDING
	templateType := TERRAFORM
	templateSource := "/path/to/file.zip"
	template, err := ExistingTemplate(templateIdentifier, templateName, templateDescription, templateStatus, templateType, templateSource, 1)
	if err != nil {
		t.Error("expected err to be nil")
	}

	input1, _ := NewTemplateAttribute(templateIdentifier, "outputA", "", STRING, "")
	input2, _ := NewTemplateAttribute(templateIdentifier, "outputB", "", NUMBER, "")
	input3, _ := NewTemplateAttribute(templateIdentifier, "outputC", "", LIST, "")
	err = template.AddOutput(input1)
	if err != nil {
		t.Error("expected err to be nil")
	}
	err = template.AddOutput(input2)
	if err != nil {
		t.Error("expected err to be nil")
	}
	err = template.AddOutput(input1)
	if err != ErrTemplateOutputAlreadyPresent {
		t.Error("expected err to be ErrOutputAlreadyPresent")
	}

	if len(template.ListOutputs()) != 2 {
		t.Errorf("expected len(inputs) to be 2, got %d", len(template.ListOutputs()))
	}

	err = template.RemoveOutput(input1.GetIdentifier().ToString())
	if err != nil {
		t.Error("expected err to be nil")
	}
	err = template.RemoveOutput(input3.GetIdentifier().ToString())
	if err != ErrTemplateOutputNotFound {
		t.Error("expected err to be ErrOutputNotFound")
	}

	if len(template.ListOutputs()) != 1 {
		t.Errorf("expected len(inputs) to be 1, got %d", len(template.ListOutputs()))
	}
}
