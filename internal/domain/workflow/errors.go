package workflow

import "errors"

var (
	ErrUnsupportedAttributeType         = errors.New("unsupported attribute type")
	ErrInvalidListFormat                = errors.New("invalid LIST format: expected format like [a, b, c]")
	ErrEmptyItemInList                  = errors.New("invalid LIST format: empty item")
	ErrUnsupportedWorkflowAttributeType = errors.New("unsupported attribute type")
	ErrWorkflowInputAlreadyPresent      = errors.New("a workflow input with the same identifier is already attached")
	ErrWorkflowOutputAlreadyPresent     = errors.New("a workflow output with the same identifier is already attached")
	ErrWorkflowRunAlreadyPresent        = errors.New("a workflow run with the same date is already attached")
	ErrWorkflowInputNotFound            = errors.New("cannot find a workflow input with the specified identifier")
	ErrWorkflowOutputNotFound           = errors.New("cannot find a workflow output with the specified identifier")
	ErrWorkflowStepNotFound             = errors.New("cannot find a workflow step with the specified step number")
)
