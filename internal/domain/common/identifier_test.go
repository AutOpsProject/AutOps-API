package common

import (
	"testing"
)

func TestValidateIdentifier_Valid(t *testing.T) {
	valids := []string{
		"autops::project:abcDEF1234",
		"autops::project:abcDEF1234:template:XYZxyz7890",
		"autops::project:abcDEF1234:workflow:testID1234",
		"autops::project:abcDEF1234:policy:abcdEFG789",
	}

	for _, id := range valids {
		if err := ValidateIdentifier(id); err != nil {
			t.Errorf("Expected valid identifier, got error: %v", err)
		}
	}
}

func TestValidateIdentifier_Invalid(t *testing.T) {
	invalids := []string{
		"project:abcDEF1234",                       // missing prefix
		"autops::project",                          // too short
		"autops::project:abcDEF1234:project:badID", // nested project
		"autops::project:abcDEF1234:template",      // missing template id
		"autops::project:abcDEF123",                // id too short
		"autops::project:abcDEF1234:template:too_long_and_invalid_id",
		"autops::invalidType:abcDEF1234", // invalid resource type
		"autops::template:abcDEF1234:workflow:1234567890",
		"autops::project:abcDEF1234:something:1234567890",
	}

	for _, id := range invalids {
		if err := ValidateIdentifier(id); err == nil {
			t.Errorf("Expected error for invalid identifier: %s", id)
		}
	}
}

func TestNewIdentifier(t *testing.T) {
	id := "autops::project:abcDEF1234"
	identifier, err := NewIdentifier(id)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if identifier.ToString() != id {
		t.Errorf("Expected %s, got %s", id, identifier.ToString())
	}
}

func TestSegments(t *testing.T) {
	id := "autops::project:abcDEF1234:template:XYZxyz7890"
	identifier, _ := NewIdentifier(id)
	segments := identifier.Segments()
	if len(segments) != 6 {
		t.Errorf("Expected 6 segments, got %d", len(segments))
	}
}

func TestGetType(t *testing.T) {
	id := "autops::project:abcDEF1234:template:XYZxyz7890"
	identifier, _ := NewIdentifier(id)
	if identifier.GetType() != TEMPLATE {
		t.Errorf("Expected type TEMPLATE, got %v", identifier.GetType())
	}
}

func TestBuildIdentifiers(t *testing.T) {
	projId, err := BuildProjectIdentifier()
	if err != nil {
		t.Fatalf("BuildProjectIdentifier failed: %v", err)
	}

	tplId, err := BuildTemplateIdentifier(projId.ToString())
	if err != nil {
		t.Fatalf("BuildTemplateIdentifier failed: %v", err)
	}

	wfId, err := BuildWorkflowIdentifier(projId.ToString())
	if err != nil {
		t.Fatalf("BuildWorkflowIdentifier failed: %v", err)
	}

	polId, err := BuildPolicyIdentifier(projId.ToString())
	if err != nil {
		t.Fatalf("BuildPolicyIdentifier failed: %v", err)
	}

	attrId, err := BuildAttributeIdentifier(tplId.ToString(), "input")
	if err != nil {
		t.Fatalf("BuildAttributeIdentifier failed: %v", err)
	}

	identifiers := []*Identifier{projId, tplId, wfId, polId, attrId}
	for _, id := range identifiers {
		if err := ValidateIdentifier(id.ToString()); err != nil {
			t.Errorf("Built identifier is invalid: %s (%v)", id.ToString(), err)
		}
	}
}
