package project

import "errors"

var (
	ErrPolicyNotFound   = errors.New("cannot find a policy with the provided id in the current project")
	ErrTemplateNotFound = errors.New("cannot find a template with the provided id in the current project")
	ErrWorkflowNotFound = errors.New("cannot find a template with the provided id i, the current project")
)
