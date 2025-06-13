package workflow

import (
	"testing"

	"github.com/AutOpsProject/AutOps-API/internal/domain/common"
)

func TestWorkflowAttributeType(t *testing.T) {
	if STRING.ToString() != "string" {
		t.Errorf("expected string, got %s", STRING.ToString())
	}
	if NUMBER.ToString() != "number" {
		t.Errorf("expected number, got %s", NUMBER.ToString())
	}
	if LIST.ToString() != "list" {
		t.Errorf("expected list, got %s", LIST.ToString())
	}
	if OBJECT.ToString() != "object" {
		t.Errorf("expected object, got %s", OBJECT.ToString())
	}
	if BOOL.ToString() != "bool" {
		t.Errorf("expected bool, got %s", BOOL.ToString())
	}
	var a WorkflowAttributeType = 100
	if a.ToString() != "unknown" {
		t.Errorf("expected unknown, got %s", a.ToString())
	}

	attrType, err := ParseWorkflowAttributeType("string")
	if attrType != STRING || err != nil {
		t.Errorf("expected %d, got %d", STRING, attrType)
	}

	attrType, err = ParseWorkflowAttributeType("number")
	if attrType != NUMBER || err != nil {
		t.Errorf("expected %d, got %d", NUMBER, attrType)
	}

	attrType, err = ParseWorkflowAttributeType("bool")
	if attrType != BOOL || err != nil {
		t.Errorf("expected %d, got %d", BOOL, attrType)
	}

	attrType, err = ParseWorkflowAttributeType("list")
	if attrType != LIST || err != nil {
		t.Errorf("expected %d, got %d", LIST, attrType)
	}

	attrType, err = ParseWorkflowAttributeType("object")
	if attrType != OBJECT || err != nil {
		t.Errorf("expected %d, got %d", OBJECT, attrType)
	}

	attrType, err = ParseWorkflowAttributeType("none")
	if attrType != -1 || err == nil {
		t.Errorf("expected -1, got %d", attrType)
	}
}

func TestNewWorkflowAttribute(t *testing.T) {
	workflowId := "autops::project:ABCDEFGHIJ:workflow:1234567890"
	name := "name"
	description := "desc"
	attributeType := NUMBER
	defaultValue := "13.4"
	attribute, err := NewWorkflowAttribute(workflowId, name, description, attributeType, defaultValue)
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

func TestExistingWorkflowAttribute(t *testing.T) {
	identifier := "autops::project:ABCDEFGHIJ:workflow:1234567890"
	name := "name"
	description := "desc"
	attributeType := NUMBER
	defaultValue := "13.4"
	attribute, err := ExistingWorkflowAttribute(identifier, name, description, attributeType, defaultValue)
	if err != nil {
		t.Error("expected err to be nil")
	} else if attribute.GetIdentifier().ToString() != identifier {
		t.Errorf("expected %s, got %s", identifier, attribute.GetIdentifier())
	}

	name = "invalid name"
	description = "desc"
	attributeType = NUMBER
	defaultValue = "13.4"
	_, err = ExistingWorkflowAttribute(identifier, name, description, attributeType, defaultValue)
	if err != common.ErrInvalidName {
		t.Error("expected err to be ErrInvalidName")
	}
}

func TestSetDefaultValue(t *testing.T) {
	workflowId := "autops::project:ABCDEFGHIJ:workflow:1234567890"
	name := "name"
	description := "desc"
	attributeType := NUMBER
	defaultValue := "13.4"
	attribute, _ := NewWorkflowAttribute(workflowId, name, description, attributeType, defaultValue)

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
	attribute, err = NewWorkflowAttribute(workflowId, name, description, attributeType, defaultValue)
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
	_, err = NewWorkflowAttribute(workflowId, name, description, attributeType, defaultValue)
	if err != nil {
		t.Error("expected error to be nil")
	}

	defaultValue = "[{\"a\": \"b\", \"c\": \"d\"}, {\"a\": \"b\", \"c\": \"d\"}]"
	attributeType = LIST
	_, err = NewWorkflowAttribute(workflowId, name, description, attributeType, defaultValue)
	if err != nil {
		t.Error("expected error to be nil")
	}

	defaultValue = "[{\"a\": \"b\", \"c\": \"d\"}, {\"a\": \"b\" \"c\": \"d\"}]"
	attributeType = LIST
	_, err = NewWorkflowAttribute(workflowId, name, description, attributeType, defaultValue)
	if err != ErrInvalidListFormat {
		t.Error("expected error to be ErrInvalidListFormat")
	}
}

func TestCompare(t *testing.T) {
	id := "autops::project:ABCDEFGHIJ:workflow:AAAAAAAAAA"
	name := "name"
	description := "desc"
	attributeType := NUMBER
	defaultValue := "13.4"
	attributeA, errA := ExistingWorkflowAttribute(id, name, description, attributeType, defaultValue)
	if errA != nil {
		t.Errorf("expected err to be nil")
	}

	id = "autops::project:ABCDEFGHIJ:workflow:BBBBBBBBBB"
	name = "name"
	description = "desc"
	attributeType = NUMBER
	defaultValue = "13.4"
	attributeB, errB := ExistingWorkflowAttribute(id, name, description, attributeType, defaultValue)
	if errB != nil {
		t.Errorf("expected err to be nil")
	}

	id = "autops::project:ABCDEFGHIJ:workflow:AAAAAAAAAA"
	name = "name"
	description = "desc"
	attributeType = NUMBER
	defaultValue = "13.4"
	attributeC, errC := ExistingWorkflowAttribute(id, name, description, attributeType, defaultValue)
	if errC != nil {
		t.Errorf("expected err to be nil")
	}

	var comparator WorkflowAttributeComparator
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
