package common

import (
	"net/url"
	"path/filepath"
	"regexp"
	"strings"
)

// Status represents the execution status of a stateful entity.
type Status int

const (
	// PENDING indicates that execution is pending.
	PENDING Status = iota
	// RUNNING indicates that execution is in progress.
	RUNNING
	// SUCCESS indicates that execution has completed successfully.
	SUCCESS
	// FAILURE indicates that execution has failed.
	FAILURE
)

// ToString returns the string representation of the Status.
func (s Status) ToString() string {
	switch s {
	case PENDING:
		return "pending"
	case RUNNING:
		return "running"
	case SUCCESS:
		return "success"
	default:
		return "failure"
	}
}

// ParseStatus converts a string to a Status.
//
// It returns an error if the string does not match a known status.
func ParseStatus(str string) (Status, error) {
	str = strings.ToLower(str)
	switch str {
	case "pending":
		return PENDING, nil
	case "running":
		return RUNNING, nil
	case "success":
		return SUCCESS, nil
	case "failure":
		return FAILURE, nil
	default:
		return FAILURE, ErrStatusParseError
	}
}

// ExecutionLog holds a reference to a log path, which may be local or a URL.
type ExecutionLog struct {
	logPath string
}

var validPathRegex = regexp.MustCompile(`^[a-zA-Z0-9._/\-:\\]+$`)

// IsSyntacticallySafePath returns true if the path contains only allowed characters.
func IsSyntacticallySafePath(path string) bool {
	return validPathRegex.MatchString(path)
}

// IsPlausibleLocalPath returns true if the path appears to be a valid local file path.
func IsPlausibleLocalPath(path string) bool {
	cleaned := filepath.Clean(path)
	return cleaned != "" && !strings.ContainsRune(cleaned, 0)
}

// IsValidURL returns true if the path is a syntactically valid URL.
func IsValidURL(path string) bool {
	u, err := url.ParseRequestURI(path)
	return err == nil && u.Scheme != "" && u.Host != ""
}

// NewExecutionLog creates and returns a new ExecutionLog.
//
// It returns an error if the path is not syntactically safe or not a valid URL or local path.
func NewExecutionLog(logPath string) (*ExecutionLog, error) {
	if !IsSyntacticallySafePath(logPath) || (!IsValidURL(logPath) && !IsPlausibleLocalPath(logPath)) {
		return nil, ErrInvalidPathOrUrl
	}
	return &ExecutionLog{
		logPath: logPath,
	}, nil
}

// GetLogPath returns the path associated with the ExecutionLog.
func (l *ExecutionLog) GetLogPath() string {
	return l.logPath
}

// StatefulNamedEntity represents an identified, timestamped, and named object
// that also has a status and an optional execution log.
type StatefulNamedEntity struct {
	NamedEntity
	TaggedEntity
	status Status
	log    *ExecutionLog
}

// NewStatefulNamedEntity creates a new StatefulNamedEntity with the given name, description, and status.
// The ExecutionLog is initially nil.
func NewStatefulNamedEntity(identifier string, name string, description string, status Status) (*StatefulNamedEntity, error) {
	namedEntity, err := NewNamedEntity(identifier, name, description)
	if err != nil {
		return nil, err
	}
	return &StatefulNamedEntity{
		NamedEntity:  *namedEntity,
		TaggedEntity: *NewTaggedEntity(),
		status:       status,
		log:          nil,
	}, nil
}

// GetStatus returns the current Status of the StatefulNamedEntity.
func (s *StatefulNamedEntity) GetStatus() Status {
	return s.status
}

// SetStatus sets the Status of the StatefulNamedEntity.
func (s *StatefulNamedEntity) SetStatus(status Status) {
	s.status = status
}

// GetExecutionLog returns the ExecutionLog of the StatefulNamedEntity, or nil if none is set.
func (s *StatefulNamedEntity) GetExecutionLog() *ExecutionLog {
	return s.log
}

// SetExecutionLog assigns a new ExecutionLog to the StatefulNamedEntity.
func (s *StatefulNamedEntity) SetExecutionLog(log *ExecutionLog) {
	s.log = log
}
