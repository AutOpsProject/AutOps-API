package common

import (
	"testing"
)

func TestStatusToString(t *testing.T) {
	state := PENDING
	if state.ToString() != "pending" {
		t.Errorf("expected pending, got %s", state.ToString())
	}
	state = RUNNING
	if state.ToString() != "running" {
		t.Errorf("expected running, got %s", state.ToString())
	}
	state = SUCCESS
	if state.ToString() != "success" {
		t.Errorf("expected success, got %s", state.ToString())
	}
	state = FAILURE
	if state.ToString() != "failure" {
		t.Errorf("expected failure, got %s", state.ToString())
	}
}

func TestStatusParsing(t *testing.T) {
	state, err := ParseStatus("Pending")
	if err != nil {
		t.Errorf("expected error to be nil, got %s", err.Error())
	}
	if state != PENDING {
		t.Errorf("expected state to be %d, got %d", PENDING, state)
	}

	state, err = ParseStatus("RUNNING")
	if err != nil {
		t.Errorf("expected error to be nil, got %s", err.Error())
	}
	if state != RUNNING {
		t.Errorf("expected state to be %d, got %d", RUNNING, state)
	}

	state, err = ParseStatus("suCCess")
	if err != nil {
		t.Errorf("expected error to be nil, got %s", err.Error())
	}
	if state != SUCCESS {
		t.Errorf("expected state to be %d, got %d", SUCCESS, state)
	}

	state, err = ParseStatus("failure")
	if err != nil {
		t.Errorf("expected error to be nil, got %s", err.Error())
	}
	if state != FAILURE {
		t.Errorf("expected state to be %d, got %d", FAILURE, state)
	}

	state, err = ParseStatus("not-existing")
	if err != ErrStatusParseError {
		t.Errorf("expected error to be %s, got %s", ErrStatusParseError.Error(), err.Error())
	}
	if state != FAILURE {
		t.Errorf("expected state to be %d, got %d", FAILURE, state)
	}
}

func TestExecutionLogCreation(t *testing.T) {
	val, err := NewExecutionLog("./../../path/to/file.txt")
	if val == nil {
		t.Error("expected val to not be nil")
	}
	if err != nil {
		t.Errorf("expected no error to be returned, got %s", err.Error())
	}

	val, err = NewExecutionLog("s3://path-to/the_file.zip")
	if val == nil {
		t.Error("expected val to not be nil")
	}
	if err != nil {
		t.Errorf("expected no error to be returned, got %s", err.Error())
	}

	val, err = NewExecutionLog("<?php hack() ?>")
	if val != nil {
		t.Error("expected val to be nil")
	}
	if err != ErrInvalidPathOrUrl {
		t.Errorf("expected error to be %s, got %s", ErrInvalidPathOrUrl.Error(), err.Error())
	}

	val, err = NewExecutionLog("Not a url")
	if val != nil {
		t.Error("expected val to be nil")
	}
	if err != ErrInvalidPathOrUrl {
		t.Errorf("expected error to be %s, got %s", ErrInvalidPathOrUrl.Error(), err.Error())
	}
}

func TestExecutionLogBasicOperations(t *testing.T) {
	logPath := "/var/log/autops::project:1234567890.log"
	val, _ := NewExecutionLog(logPath)

	if val.GetLogPath() != logPath {
		t.Errorf("expected %s, got %s", logPath, val.GetLogPath())
	}
}

func TestStatefulNamedEntityConstructor(t *testing.T) {
	entity, _ := NewStatefulNamedEntity("autops::project:1234567890", "some-name", "some description", PENDING)
	if entity == nil {
		t.Error("expected entity to not be nil.")
	} else if entity.GetIdentifier().ToString() != "autops::project:1234567890" {
		t.Errorf("expected autops::project:1234567890, got %s", entity.GetIdentifier())
	} else if entity.GetName() != "some-name" {
		t.Errorf("expected some-name, got %s", entity.GetName())
	} else if entity.GetDescription() != "some description" {
		t.Errorf("expected some description, got %s", entity.GetDescription())
	} else if entity.GetStatus() != PENDING {
		t.Errorf("expected %d, got %d", PENDING, entity.GetStatus())
	} else if entity.GetExecutionLog() != nil {
		t.Error("expected execution log to be nil")
	}
}

func TestStatefulNamedEntitySetters(t *testing.T) {
	entity, _ := NewStatefulNamedEntity("autops::project:1234567890", "some-name", "some description", PENDING)
	if entity == nil {
		t.Error("expected entity to not be nil.")
	} else if entity.GetStatus() != PENDING {
		t.Errorf("expected %d, got %d", PENDING, entity.GetStatus())
	} else if entity.GetExecutionLog() != nil {
		t.Error("expected execution log to be nil")
	}

	entity.SetStatus(SUCCESS)
	if entity.GetStatus() != SUCCESS {
		t.Errorf("expected %d, got %d", SUCCESS, entity.GetStatus())
	}

	log, _ := NewExecutionLog("some/path")
	if log == nil {
		t.Errorf("expected log to not be nil.")
	}

	entity.SetExecutionLog(log)
	if entity.GetExecutionLog() != log {
		t.Errorf("expected %p, got %p", log, entity.GetExecutionLog())
	}
}
