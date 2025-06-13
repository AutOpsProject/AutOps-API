package template

import (
	"strings"

	"github.com/AutOpsProject/AutOps-API/internal/domain/common"
)

// TemplateType defines the type of automation logic used in a template.
//
// The supported types are:
//   - TERRAFORM
//   - ANSIBLE
//   - PACKER
//   - OPENTOFU
type TemplateType int

const (
	TERRAFORM TemplateType = iota
	ANSIBLE
	PACKER
	OPENTOFU
)

// ToString returns the string representation of a TemplateType.
// This is useful for serialization or display in user interfaces.
func (t TemplateType) ToString() string {
	switch t {
	case TERRAFORM:
		return "terraform"
	case ANSIBLE:
		return "ansible"
	case PACKER:
		return "packer"
	default:
		return "opentofu"
	}
}

// ParseTemplateType converts a string into a TemplateType.
// It accepts "terraform", "ansible", or "cloudformation" (case-insensitive).
// Returns an error if the string does not match a known type.
func ParseTemplateType(str string) (TemplateType, error) {
	str = strings.ToLower(str)
	switch str {
	case "terraform":
		return TERRAFORM, nil
	case "ansible":
		return ANSIBLE, nil
	case "packer":
		return PACKER, nil
	case "opentofu":
		return OPENTOFU, nil
	default:
		return -1, ErrParseInvalidTemplateType
	}
}

// Template represents an infrastructure or configuration template.
// It includes its type, source path (local or remote), version, and a set of defined inputs/outputs.
// Templates are versioned and validated upon update.
type Template struct {
	common.StatefulNamedEntity
	common.VersionedSource
	templateType TemplateType
	inputs       *common.List[*TemplateAttribute]
	outputs      *common.List[*TemplateAttribute]
}

type TemplateComparator struct{}

func (TemplateComparator) Compare(t1, t2 *Template) int {
	return strings.Compare(t1.GetIdentifier().ToString(), t2.GetIdentifier().ToString())
}

// NewTemplate creates a new Template with a generated identifier.
// It validates and sets the source path, initializes version to 1, and adds timestamps.
// Returns an error if the source path is invalid.
func NewTemplate(projectIdentifier string, name string, description string, status common.Status, templateType TemplateType, sourcePath string) (*Template, error) {
	templateIdentifier, err := common.BuildTemplateIdentifier(projectIdentifier)
	if err != nil {
		return nil, err
	}
	return ExistingTemplate(templateIdentifier.ToString(), name, description, status, templateType, sourcePath, 1)
}

// ExistingTemplate creates a Template using an existing identifier.
// Typically used when loading from persistent storage.
// It validates the source path and sets version and timestamps accordingly.
func ExistingTemplate(templateIdentifier string, name string, description string, status common.Status, templateType TemplateType, sourcePath string, version int) (*Template, error) {
	statefulEntity, err := common.NewStatefulNamedEntity(templateIdentifier, name, description, status)
	if err != nil {
		return nil, err
	}
	versionedSource, err := common.NewVersionedSource(sourcePath, version)
	if err != nil {
		return nil, err
	}
	template := &Template{
		StatefulNamedEntity: *statefulEntity,
		VersionedSource:     *versionedSource,
		templateType:        templateType,
		inputs:              common.NewList(TemplateAttributeComparator{}, []*TemplateAttribute{}),
		outputs:             common.NewList(TemplateAttributeComparator{}, []*TemplateAttribute{}),
	}
	return template, nil
}

// GetTemplateType returns the TemplateType associated with the Template.
func (t *Template) GetTemplateType() TemplateType {
	return t.templateType
}

// AddInput adds a new input to the template.
// Returns an error if an input with the same identifier already exists.
func (t *Template) AddInput(input *TemplateAttribute) error {
	if t.inputs.Contains(input) {
		return ErrTemplateInputAlreadyPresent
	}
	t.inputs.Append(input)
	return nil
}

// RemoveInput removes the input with the given identifier from the template, if it exists.
func (t *Template) RemoveInput(inputIdentifier string) error {
	attribute, found := t.inputs.SelectOne(func(t *TemplateAttribute) bool {
		return t.GetIdentifier().ToString() == inputIdentifier
	})
	if !found {
		return ErrTemplateInputNotFound
	}
	t.inputs.Remove(attribute)
	return nil
}

// ListInputs returns all defined input attributes of the template.
func (t *Template) ListInputs() []*TemplateAttribute {
	return t.inputs.Items()
}

// AddOutput adds a new output to the template.
// Returns an error if an output with the same identifier already exists.
func (t *Template) AddOutput(output *TemplateAttribute) error {
	if t.outputs.Contains(output) {
		return ErrTemplateOutputAlreadyPresent
	}
	t.outputs.Append(output)
	return nil
}

// RemoveOutput removes the output with the given identifier from the template, if it exists.
func (t *Template) RemoveOutput(outputIdentifier string) error {
	attribute, found := t.outputs.SelectOne(func(t *TemplateAttribute) bool {
		return t.GetIdentifier().ToString() == outputIdentifier
	})
	if !found {
		return ErrTemplateOutputNotFound
	}
	t.outputs.Remove(attribute)
	return nil
}

// ListOutputs returns all defined output attributes of the template.
func (t *Template) ListOutputs() []*TemplateAttribute {
	return t.outputs.Items()
}
