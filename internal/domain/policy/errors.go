package policy

import "errors"

var (
	ErrInvalidPolicyAction = errors.New("invalid action name for the specified resource type")
	ErrInvalidPolicyEffect = errors.New("invalid policy effect : correct values are 'ALLOW' or 'DENY'")
)
