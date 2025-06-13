package template

import (
	"testing"

	"github.com/AutOpsProject/AutOps-API/internal/domain/common"
)

func TestNewTemplateAttribute(t *testing.T) {
	templateId := "autops::project:1234567890:template:abcdefghij"
	name := "name"
	description := "desc"
	attributeType := NUMBER
	defaultValue := "13.4"
	attribute, err := NewTemplateAttribute(templateId, name, description, attributeType, defaultValue)
	if err != nil {
		t.Error("expected err to be nil")
	} else if attribute.GetIdentifier().ToString() == "" {
		t.Errorf("expected identifier to be not empty")
	} else if attribute.GetName() != name {
		t.Errorf("expected %s, got %s", name, attribute.GetName())
	} else if attribute.GetDescription() != description {
		t.Errorf("expected %s, got %s", description, attribute.GetDescription())
	} else if attribute.GetType() != attributeType {
		t.Errorf("expected %d, got %d", attributeType, attribute.GetType())
	} else if attribute.GetDefaultValue() != defaultValue {
		t.Errorf("expected %s, got %s", defaultValue, attribute.GetDefaultValue())
	}
}

func TestExistingTemplateAttribute(t *testing.T) {
	templateId := "autops::project:1234567890:template:abcdefghij"
	name := "name"
	description := "desc"
	attributeType := NUMBER
	defaultValue := "13.4"
	attribute, err := ExistingTemplateAttribute(templateId, name, description, attributeType, defaultValue)
	if err != nil {
		t.Error("expected err to be nil")
	} else if attribute.GetIdentifier().ToString() != templateId {
		t.Errorf("expected %s, got %s", templateId, attribute.GetIdentifier())
	}

	name = "invalid name"
	description = "desc"
	attributeType = NUMBER
	defaultValue = "13.4"
	_, err = ExistingTemplateAttribute(templateId, name, description, attributeType, defaultValue)
	if err != common.ErrInvalidName {
		t.Error("expected err to be ErrInvalidName")
	}
}

func TestSetDefaultValue(t *testing.T) {
	templateId := "autops::project:1234567890:template:abcdefghij"
	name := "name"
	description := "desc"
	attributeType := NUMBER
	defaultValue := "13.4"
	attribute, _ := NewTemplateAttribute(templateId, name, description, attributeType, defaultValue)

	// Testing NUMBER type
	err := attribute.SetDefaultValue("abcd")
	if err == nil {
		t.Error("expected err to be ErrSyntax")
	}
	err = attribute.SetDefaultValue("-123.18780")
	if err != nil {
		t.Error("expected err to be nil")
	}

	// Testing LIST type
	attributeType = LIST
	defaultValue = "[1, 2, 3]"
	attribute, err = NewTemplateAttribute(templateId, name, description, attributeType, defaultValue)
	if err != nil {
		t.Error("expected error to be nil")
	}

	defaultValue = "[1, , 3]"
	err = attribute.SetDefaultValue(defaultValue)
	if err != ErrInvalidListFormat {
		t.Error("expected error to be ErrInvalidListFormat")
	}

	defaultValue = "[{\"a\": \"b\"}, {\"c\": \"d\"}]"
	err = attribute.SetDefaultValue(defaultValue)
	if err != nil {
		t.Error("expected error to be nil")
	}

	defaultValue = "12"
	err = attribute.SetDefaultValue(defaultValue)
	if err != ErrInvalidListFormat {
		t.Error("expected error to be ErrInvalidListFormat")
	}

	defaultValue = "string"
	err = attribute.SetDefaultValue(defaultValue)
	if err != ErrInvalidListFormat {
		t.Error("expected error to be ErrInvalidListFormat")
	}

	// Testing OBJECT type

	defaultValue = "{\"a\": \"b\", \"c\": \"d\"}"
	attributeType = OBJECT
	_, err = NewTemplateAttribute(templateId, name, description, attributeType, defaultValue)
	if err != nil {
		t.Error("expected error to be nil")
	}

	defaultValue = "[{\"a\": \"b\", \"c\": \"d\"}, {\"a\": \"b\", \"c\": \"d\"}]"
	attributeType = LIST
	_, err = NewTemplateAttribute(templateId, name, description, attributeType, defaultValue)
	if err != nil {
		t.Error("expected error to be nil")
	}

	defaultValue = "[{\"a\": \"b\", \"c\": \"d\"}, {\"a\": \"b\" \"c\": \"d\"}]"
	attributeType = LIST
	_, err = NewTemplateAttribute(templateId, name, description, attributeType, defaultValue)
	if err != ErrInvalidListFormat {
		t.Error("expected error to be ErrInvalidListFormat")
	}
}

func TestCompare(t *testing.T) {
	id := "autops::project:1234567890:template:AAAAAAAAAA"
	name := "name"
	description := "desc"
	attributeType := NUMBER
	defaultValue := "13.4"
	attributeA, errA := ExistingTemplateAttribute(id, name, description, attributeType, defaultValue)
	if errA != nil {
		t.Errorf("expected err to be nil")
	}

	id = "autops::project:1234567890:template:BBBBBBBBBB"
	name = "name"
	description = "desc"
	attributeType = NUMBER
	defaultValue = "13.4"
	attributeB, errB := ExistingTemplateAttribute(id, name, description, attributeType, defaultValue)
	if errB != nil {
		t.Errorf("expected err to be nil")
	}

	id = "autops::project:1234567890:template:AAAAAAAAAA"
	name = "name"
	description = "desc"
	attributeType = NUMBER
	defaultValue = "13.4"
	attributeC, errC := ExistingTemplateAttribute(id, name, description, attributeType, defaultValue)
	if errC != nil {
		t.Errorf("expected err to be nil")
	}

	var comparator TemplateAttributeComparator
	if comparator.Compare(attributeA, attributeB) != -1 {
		t.Errorf("expected comparison to return -1, got %d", comparator.Compare(attributeA, attributeB))
	}
	if comparator.Compare(attributeB, attributeA) != 1 {
		t.Errorf("expected comparison to return 1, got %d", comparator.Compare(attributeB, attributeA))
	}
	if comparator.Compare(attributeC, attributeA) != 0 {
		t.Errorf("expected comparison to return 0, got %d", comparator.Compare(attributeC, attributeA))
	}
}
