package common

import (
	"testing"
)

func TestSetGetName_Valid(t *testing.T) {
	name := "valid_name-1234567890"
	entity := &NamedEntity{}
	err := entity.SetName(name)
	if err != nil {
		t.Errorf("expected no error, got: %v", err)
	}
	if entity.GetName() != name {
		t.Errorf("expected %s to be %s", name, entity.GetName())
	}
}

func TestSetName_Invalid(t *testing.T) {
	entity := &NamedEntity{}

	invalidNames := []string{
		"",                                    // empty
		"with spaces",                         // space
		"@invalid!",                           // special chars
		"toolong" + string(make([]byte, 130)), // > 128
	}

	for _, name := range invalidNames {
		if err := entity.SetName(name); err == nil {
			t.Errorf("expected error for name %q, got nil", name)
		}
	}
}

func TestSetGetDescription_Valid(t *testing.T) {
	entity := &NamedEntity{}
	description := "This is a description."
	err := entity.SetDescription(description)
	if err != nil {
		t.Errorf("expected no error, got: %v", err)
	}
	if entity.GetDescription() != description {
		t.Errorf("expected %s to be %s", description, entity.GetDescription())
	}
}

func TestSetDescription_TooLong(t *testing.T) {
	entity := &NamedEntity{}
	longDesc := make([]byte, 513)
	for i := range longDesc {
		longDesc[i] = 'a'
	}
	err := entity.SetDescription(string(longDesc))
	if err == nil {
		t.Error("expected error for too long description, got nil")
	}
}

func TestNewTimestampedEntity_Valid(t *testing.T) {
	date := CurrentTimestamp()
	_, err := NewTimestampedEntity("1234567890")
	if err != ErrInvalidIdentifierFormat {
		t.Error("expected err to be ErrInvalidIdentifierFormat")
	}
	entity, err := NewTimestampedEntity("autops::project:1234567890")
	if err != nil {
		t.Error("expected err to be nil")
	}
	if entity.GetIdentifier().ToString() != "autops::project:1234567890" {
		t.Errorf("expected 1234567890, got %s", entity.GetIdentifier())
	}
	if entity.GetCreatedAt() != date {
		t.Errorf("expected %s to be %s", date, entity.GetCreatedAt())
	}
	if entity.GetUpdatedAt() != date {
		t.Errorf("expected %s to be %s", date, entity.GetUpdatedAt())
	}
}

func TestNewTimestampedEntityWithAttribute_Valid(t *testing.T) {
	creationDate := "1976-01-01 00:00:00+02:00:00"
	updatedAt := "2002-01-01 00:00:00+02:00:00"
	entity, err := ExistingTimestampedEntity("autops::project:1234567890", creationDate, updatedAt)
	if err != nil {
		t.Errorf("expected err to be nil")
	}
	if entity.GetCreatedAt() != creationDate {
		t.Errorf("expected %s to be %s", creationDate, entity.GetCreatedAt())
	}
	if entity.GetUpdatedAt() != updatedAt {
		t.Errorf("expected %s to be %s", updatedAt, entity.GetUpdatedAt())
	}
}

func TestUpdateModificationDate(t *testing.T) {
	creationDate := "1976-01-01 00:00:00+02:00:00"
	updatedAt := "2002-01-01 00:00:00+02:00:00"
	entity, _ := ExistingTimestampedEntity("autops::project:1234567890", creationDate, updatedAt)
	currentDate := CurrentTimestamp()
	entity.UpdateModificationDate()
	if entity.GetUpdatedAt() != currentDate {
		t.Errorf("expected %s to be %s", entity.GetUpdatedAt(), currentDate)
	}
}

func TestNewNamedEntityWithInvalidFields(t *testing.T) {
	_, err := NewNamedEntity("autops::project:1234567890", "invalid name", "")
	if err != ErrInvalidName {
		t.Errorf("expected err to not be ErrInvalidName")
	}
	longDesc := make([]byte, 513)
	for i := range longDesc {
		longDesc[i] = 'a'
	}

	_, err = NewNamedEntity("autops::project:1234567890", "valid-name", string(longDesc))
	if err != ErrInvalidDescription {
		t.Errorf("expected err to be ErrInvalidDescription")
	}

	_, err = NewStatefulNamedEntity("autops::project:1234567890", "invalid name", "", PENDING)
	if err != ErrInvalidName {
		t.Errorf("expected err to not be ErrInvalidName")
	}

	_, err = NewStatefulNamedEntity("autops::project:1234567890", "valid-name", string(longDesc), FAILURE)
	if err != ErrInvalidDescription {
		t.Errorf("expected err to be ErrInvalidDescription")
	}
}
