package common

import "strings"

// Tag represents a key-value pair used to annotate objects.
type Tag struct {
	key   string
	value string
}

// NewTag creates a new Tag with the specified key and value.
func NewTag(key string, value string) *Tag {
	return &Tag{
		key:   key,
		value: value,
	}
}

// GetKey returns the key of the Tag.
func (t *Tag) GetKey() string {
	return t.key
}

// GetValue returns the value of the Tag.
func (t *Tag) GetValue() string {
	return t.value
}

// SetValue updates the value of the Tag.
func (t *Tag) SetValue(value string) {
	t.value = value
}

// TagComparator provides a comparison function for Tags based on their key.
type TagComparator struct{}

// Compare compares two Tags by their key.
//
// Returns -1 if a < b, 1 if a > b, and 0 if equal.
func (TagComparator) Compare(a *Tag, b *Tag) int {
	return strings.Compare(a.GetKey(), b.GetKey())
}

// TaggedEntity represents an object that can have multiple unique tags.
type TaggedEntity struct {
	tags *List[*Tag]
}

// NewTaggedEntity creates a new TaggedEntity with an empty tag list.
func NewTaggedEntity() *TaggedEntity {
	return &TaggedEntity{
		tags: NewList(TagComparator{}, []*Tag{}),
	}
}

// AddTag adds or updates a tag on the TaggedEntity.
//
// If a tag with the same key already exists, its value is updated.
func (t *TaggedEntity) AddTag(tag *Tag) {
	index, _ := t.tags.GetItem(tag)
	if index != -1 {
		t.RemoveTag(tag.GetKey())
	}
	t.tags.Append(tag)
}

// RemoveTag removes the tag with the specified key from the TaggedEntity.
func (t *TaggedEntity) RemoveTag(key string) {
	t.tags.Remove(&Tag{key, ""})
}

// HasTag checks whether a tag with the given key exists on the TaggedEntity.
func (t *TaggedEntity) HasTag(key string) bool {
	return t.tags.Contains(&Tag{key, ""})
}

// GetTag returns the tag with the given key, or nil if not found.
func (t *TaggedEntity) GetTag(key string) *Tag {
	_, val := t.tags.GetItem(&Tag{key, ""})
	if val == nil {
		return nil
	}
	return *val
}

// ListTags returns a slice of all tags attached to the TaggedEntity.
func (t *TaggedEntity) ListTags() []*Tag {
	return t.tags.Items()
}
