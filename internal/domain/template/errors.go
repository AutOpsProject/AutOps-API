package template

import "errors"

var (
	ErrUnsupportedAttributeType     = errors.New("unsupported attribute type")
	ErrInvalidListFormat            = errors.New("invalid LIST format: expected format like [a, b, c]")
	ErrEmptyItemInList              = errors.New("invalid LIST format: empty item")
	ErrTemplateInputAlreadyPresent  = errors.New("an input with the same identifier is already present in the template")
	ErrTemplateOutputAlreadyPresent = errors.New("an output with the same identifier is already present in the template")
	ErrParseInvalidTemplateType     = errors.New("cannot parse the string into a TemplateType")
	ErrTemplateOutputNotFound       = errors.New("cannot find an output with the specified identifier")
	ErrTemplateInputNotFound        = errors.New("cannot find an input with the specified identifier")
	ErrInvalidAttributeType         = errors.New("invalid attribute type")
)
