package common

import (
	"testing"
)

func TestTagBasicFunctions(t *testing.T) {
	key := "key"
	value := "value"

	tag := NewTag(key, value)

	if tag.GetKey() != key {
		t.Errorf("expected %s to be %s", key, tag.GetKey())
	}
	if tag.GetValue() != value {
		t.Errorf("expected %s to be %s", value, tag.GetValue())
	}

	newValue := "changed"
	tag.SetValue(newValue)

	if tag.GetKey() != key {
		t.Errorf("expected %s to be %s", key, tag.GetKey())
	}
	if tag.GetValue() != newValue {
		t.Errorf("expected %s to be %s", newValue, tag.GetValue())
	}
}

func TestTaggedEntityBasicOperations(t *testing.T) {
	// Test constructor
	taggedEntity := NewTaggedEntity()
	tags := taggedEntity.ListTags()
	if len(tags) > 0 {
		t.Errorf("expected intialized tags list to be empty, got length : %d", len(tags))
	}

	tag1 := NewTag("key1", "value1")
	tag2 := NewTag("key2", "value2")
	tag3 := NewTag("key3", "value3")
	// Test AddTag
	taggedEntity.AddTag(tag1)
	taggedEntity.AddTag(tag2)

	tag1.SetValue("other")
	// Test HasTag
	if !taggedEntity.HasTag(tag1.GetKey()) {
		t.Errorf("expected tags to have key %s, but not found", tag1.GetKey())
	}
	if !taggedEntity.HasTag(tag2.GetKey()) {
		t.Errorf("expected tags to have key %s, but not found", tag2.GetKey())
	}
	if taggedEntity.HasTag(tag3.GetKey()) {
		t.Errorf("expected tags to do not have key %s, but found it", tag3.GetKey())
	}

	// Test GetTag
	foundTag := taggedEntity.GetTag(tag1.GetKey())
	if foundTag.GetKey() != tag1.GetKey() || foundTag.GetValue() != tag1.GetValue() {
		t.Errorf("expected tag (%s, %s) to be (%s, %s)", foundTag.GetKey(), foundTag.GetValue(), tag1.GetKey(), tag1.GetValue())
	}
	foundTag = taggedEntity.GetTag(tag3.GetKey())
	if foundTag != nil {
		t.Errorf("expected tag with key %s to be absent, but found.", tag3.GetKey())
	}

	// Test AddTag with a tag key already present in tags list
	existingTag := NewTag("key1", "other")
	taggedEntity.AddTag(existingTag)
	foundTag = taggedEntity.GetTag(existingTag.GetKey())
	if foundTag.GetKey() != existingTag.GetKey() || foundTag.GetValue() != existingTag.GetValue() {
		t.Errorf("expected tag (%s, %s) to be (%s, %s)", foundTag.GetKey(), foundTag.GetValue(), existingTag.GetKey(), existingTag.GetValue())
	}
	if len(taggedEntity.ListTags()) != 2 {
		t.Errorf("expected 2 as a number of attached tags, found %d", len(taggedEntity.ListTags()))
	}
}
