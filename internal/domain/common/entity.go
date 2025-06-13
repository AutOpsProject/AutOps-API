package common

import (
	"regexp"
	"strings"
	"time"
)

var nameRegex = regexp.MustCompile(`^[a-zA-Z0-9_-]{1,128}$`)

const _NAME_MIN_LENGTH = 1
const _NAME_MAX_LENGTH = 128
const _DESCRIPTION_MAX_LENGTH = 512

func CurrentTimestamp() string {
	return time.Now().Format(time.RFC3339)
}

// TimestampedEntity represents an identified object with creation and update timestamps.
type TimestampedEntity struct {
	identifier *Identifier
	createdAt  string
	updatedAt  string
}

// NewTimestampedEntity creates a TimestampedEntity with the current time as both creation and update time.
func NewTimestampedEntity(identifier string) (*TimestampedEntity, error) {
	date := CurrentTimestamp()
	return ExistingTimestampedEntity(identifier, date, date)
}

// ExistingTimestampedEntity creates a TimestampedEntity with the provided creation and update timestamps.
//
// Both timestamps should be valid RFC3339 date strings.
func ExistingTimestampedEntity(identifier string, createdAt string, updatedAt string) (*TimestampedEntity, error) {
	id, err := NewIdentifier(identifier)
	if err != nil {
		return nil, err
	}
	return &TimestampedEntity{
		identifier: id,
		createdAt:  createdAt,
		updatedAt:  updatedAt,
	}, nil
}

// GetIdentifier returns the identifier of the TimestampedEntity.
func (r *TimestampedEntity) GetIdentifier() *Identifier {
	return r.identifier
}

// GetCreatedAt returns the creation timestamp of the TimestampedEntity.
func (t *TimestampedEntity) GetCreatedAt() string {
	return t.createdAt
}

// GetUpdatedAt returns the update timestamp of the TimestampedEntity.
func (t *TimestampedEntity) GetUpdatedAt() string {
	return t.updatedAt
}

// UpdateModificationDate sets the updatedAt timestamp to the current time.
func (t *TimestampedEntity) UpdateModificationDate() {
	t.updatedAt = CurrentTimestamp()
}

// NamedEntity represents a timestamped and identified object with a name and a description.
type NamedEntity struct {
	TimestampedEntity
	name        string
	description string
}

// NewNamedEntity creates a new NamedEntity with the provided name and description.
func NewNamedEntity(identifier string, name string, description string) (*NamedEntity, error) {
	date := CurrentTimestamp()
	return ExistingNamedEntity(identifier, name, description, date, date)
}

// ExistingNamedEntity creates a NamedEntity with the provided parameters.
func ExistingNamedEntity(identifier string, name string, description string, createdAt string, updatedAt string) (*NamedEntity, error) {
	timedEntity, err := NewTimestampedEntity(identifier)
	if err != nil {
		return nil, err
	}
	namedEntity := &NamedEntity{
		TimestampedEntity: *timedEntity,
		name:              "",
		description:       "",
	}
	err = namedEntity.SetName(name)
	if err != nil {
		return nil, err
	}
	err = namedEntity.SetDescription(description)
	if err != nil {
		return nil, err
	}

	return namedEntity, nil
}

// GetName returns the name of the NamedEntity.
func (n *NamedEntity) GetName() string {
	return n.name
}

// SetName sets the name of the NamedEntity if it is valid.
//
// The name must be non-empty, up to 128 characters, and contain only letters, numbers, hyphens, or underscores.
func (n *NamedEntity) SetName(name string) error {
	if len(name) < _NAME_MIN_LENGTH || len(name) > _NAME_MAX_LENGTH || !nameRegex.MatchString(name) {
		return ErrInvalidName
	}
	n.name = name
	return nil
}

// GetDescription returns the description of the NamedEntity.
func (n *NamedEntity) GetDescription() string {
	return n.description
}

// SetDescription sets the description of the NamedEntity if it does not exceed 512 characters.
func (n *NamedEntity) SetDescription(description string) error {
	description = strings.TrimSpace(description)
	if len(description) > _DESCRIPTION_MAX_LENGTH {
		return ErrInvalidDescription
	}
	n.description = description
	return nil
}
