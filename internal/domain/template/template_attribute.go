package template

import (
	"encoding/json"
	"strconv"
	"strings"

	"github.com/AutOpsProject/AutOps-API/internal/domain/common"
)

// AttributeType defines the type of value that a TemplateAttribute can hold.
// It determines how the default value will be validated.
//
// The supported types are:
//   - STRING: a plain string
//   - NUMBER: a numeric value (integer or float)
//   - BOOL: a boolean value ("true" or "false")
//   - LIST: a JSON array of values (e.g., ["a", 1, {"k":"v"}])
//   - OBJECT: a JSON object (e.g., {"key": "value"})
type AttributeType int

const (
	STRING AttributeType = iota
	NUMBER
	BOOL
	LIST
	OBJECT
)

// TODO : add type parser and ToString
func (t AttributeType) ToString() (string, error) {
	switch t {
	case STRING:
		return "S", nil
	case NUMBER:
		return "N", nil
	case BOOL:
		return "B", nil
	case LIST:
		return "L", nil
	case OBJECT:
		return "O", nil
	default:
		return "", ErrInvalidAttributeType
	}
}

// TemplateAttribute represents a user-defined parameter for a template.
// Each attribute has a name, description, type, and a default value (as string).
// The default value is validated according to the attribute type.
type TemplateAttribute struct {
	common.NamedEntity
	attributeType AttributeType
	defaultValue  string
}

// NewTemplateAttribute creates a new TemplateAttribute instance with a generated identifier.
// It validates the default value according to the specified attribute type.
//
// Returns an error if validation fails.
func NewTemplateAttribute(templateId string, name string, description string, attributeType AttributeType, defaultValue string) (*TemplateAttribute, error) {
	typeStr, err := attributeType.ToString()
	if err != nil {
		return nil, err
	}
	identifier, err := common.BuildAttributeIdentifier(templateId, typeStr)
	if err != nil {
		return nil, err
	}
	return ExistingTemplateAttribute(identifier.ToString(), name, description, attributeType, defaultValue)
}

// ExistingTemplateAttribute creates a TemplateAttribute with an existing identifier.
// This is typically used when reloading from a data store.
//
// It validates the default value against the attribute type and returns an error if invalid.
func ExistingTemplateAttribute(identifier string, name string, description string, attributeType AttributeType, defaultValue string) (*TemplateAttribute, error) {
	namedEntity, err := common.NewNamedEntity(identifier, name, description)
	if err != nil {
		return nil, err
	}
	attribute := TemplateAttribute{
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

// GetType returns the AttributeType associated with the TemplateAttribute.
func (a *TemplateAttribute) GetType() AttributeType {
	return a.attributeType
}

// GetDefaultValue returns the default value set for the attribute as a raw string.
func (a *TemplateAttribute) GetDefaultValue() string {
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
func (a *TemplateAttribute) SetDefaultValue(value string) error {
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

// TemplateAttributeComparator is used to compare two TemplateAttribute instances
// based on their identifier. It enables deterministic sorting within lists.
type TemplateAttributeComparator struct{}

// Compare returns a comparison between two TemplateAttribute identifiers.
// It is used to order attributes consistently within lists.
func (TemplateAttributeComparator) Compare(a *TemplateAttribute, b *TemplateAttribute) int {
	return strings.Compare(a.GetIdentifier().ToString(), b.GetIdentifier().ToString())
}
