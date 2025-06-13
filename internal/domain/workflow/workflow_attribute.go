package workflow

import (
	"encoding/json"
	"strconv"
	"strings"

	"github.com/AutOpsProject/AutOps-API/internal/domain/common"
)

// WorkflowAttributeType defines the type of an attribute in a workflow.
type WorkflowAttributeType int

const (
	// STRING represents a string attribute type.
	STRING WorkflowAttributeType = iota
	// NUMBER represents a numeric attribute type.
	NUMBER
	// LIST represents a list attribute type.
	LIST
	// BOOL representes a boolean attribute type.
	BOOL
	// OBJECT represents an object (JSON) attribute type.
	OBJECT
)

// ToString converts the WorkflowAttributeType to its string representation.
func (t WorkflowAttributeType) ToString() string {
	switch t {
	case STRING:
		return "string"
	case NUMBER:
		return "number"
	case BOOL:
		return "bool"
	case LIST:
		return "list"
	case OBJECT:
		return "object"
	default:
		return "unknown"
	}
}

// ParseWorkflowAttributeType parses a string into a WorkflowAttributeType.
// Returns an error if the string does not match a known type.
func ParseWorkflowAttributeType(str string) (WorkflowAttributeType, error) {
	str = strings.ToLower(str)
	switch str {
	case "string":
		return STRING, nil
	case "number":
		return NUMBER, nil
	case "bool":
		return BOOL, nil
	case "list":
		return LIST, nil
	case "object":
		return OBJECT, nil
	default:
		return -1, ErrUnsupportedWorkflowAttributeType
	}
}

// WorkflowAttribute represents a named attribute within a workflow,
// with a specified type and optional default value.
type WorkflowAttribute struct {
	common.NamedEntity
	attributeType WorkflowAttributeType
	defaultValue  string
}

// NewWorkflowAttribute creates a new WorkflowAttribute with a generated unique identifier.
// Returns an error if the name, description, or default value is invalid.
func NewWorkflowAttribute(workflowId string, name string, description string, attributeType WorkflowAttributeType, defaultValue string) (*WorkflowAttribute, error) {
	identifier, err := common.BuildAttributeIdentifier(workflowId, attributeType.ToString())
	if err != nil {
		return nil, err
	}
	return ExistingWorkflowAttribute(identifier.ToString(), name, description, attributeType, defaultValue)
}

// ExistingWorkflowAttribute creates a WorkflowAttribute with the provided identifier, name, description, type, and default value.
// Returns an error if the name, description, or default value is invalid.
func ExistingWorkflowAttribute(identifier string, name string, description string, attributeType WorkflowAttributeType, defaultValue string) (*WorkflowAttribute, error) {
	namedEntity, err := common.NewNamedEntity(identifier, name, description)
	if err != nil {
		return nil, err
	}

	attribute := WorkflowAttribute{
		NamedEntity:   *namedEntity,
		attributeType: attributeType,
		defaultValue:  "",
	}

	err = attribute.SetDefaultValue(defaultValue)
	if err != nil {
		return nil, err
	}

	return &attribute, nil
}

// GetType returns the WorkflowAttributeType of the attribute.
func (a *WorkflowAttribute) GetType() WorkflowAttributeType {
	return a.attributeType
}

// GetDefaultValue returns the default value of the attribute.
func (a *WorkflowAttribute) GetDefaultValue() string {
	return a.defaultValue
}

// SetDefaultValue validates and sets the default value for the attribute.
//   - For STRING: any string is accepted.
//   - For NUMBER: the value must be a valid float or integer.
//   - For BOOL: the value must be "true" or "false".
//   - For OBJECT: the value must be valid JSON of type object.
//   - For LIST: the value must be valid JSON of type array, and no element must be nil.
//
// Returns an error if the value is invalid for the given type.
func (a *WorkflowAttribute) SetDefaultValue(value string) error {
	value = strings.TrimSpace(value)
	if value == "" {
		return nil
	}
	switch a.attributeType {
	case STRING:
		break

	case NUMBER:
		if _, err := strconv.ParseFloat(value, 64); err != nil {
			return err
		}

	case BOOL:
		if _, err := strconv.ParseBool(value); err != nil {
			return err
		}

	case OBJECT:
		var obj map[string]interface{}
		if err := json.Unmarshal([]byte(value), &obj); err != nil {
			return err
		}

	case LIST:
		var rawList []interface{}
		if err := json.Unmarshal([]byte(value), &rawList); err != nil {
			return ErrInvalidListFormat
		}

		for _, item := range rawList {
			if item == nil {
				return ErrEmptyItemInList
			}
		}

	default:
		return ErrUnsupportedAttributeType
	}

	a.defaultValue = value
	return nil
}

// WorkflowAttributeComparator is used to compare two WorkflowAttribute instances
// based on their identifier. It enables deterministic sorting within lists.
type WorkflowAttributeComparator struct{}

// Compare returns a comparison between two WorkflowAttribute identifiers.
// It is used to order attributes consistently within lists.
func (WorkflowAttributeComparator) Compare(a *WorkflowAttribute, b *WorkflowAttribute) int {
	return strings.Compare(a.GetIdentifier().ToString(), b.GetIdentifier().ToString())
}
