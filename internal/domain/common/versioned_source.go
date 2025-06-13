package common

// VersionedSource encapsulates a source path and its version number.
// It is used by entities like Template and Workflow to manage versioning
// and track the source location (either a local file path or a URL).
type VersionedSource struct {
	sourcePath string
	version    int
}

// NewVersionnedSource creates a new VersionedSource with the provided sourcePath and version.
// It validates the sourcePath to ensure it is either a syntactically safe local path or a valid URL.
// Returns an error if the provided path is invalid.
func NewVersionedSource(sourcePath string, version int) (*VersionedSource, error) {
	entity := VersionedSource{
		sourcePath: "",
		version:    version,
	}
	err := entity.SetSourcePath(sourcePath)
	if err != nil {
		return nil, err
	}
	return &entity, nil
}

// GetSourcePath returns the source path of the entity.
func (t *VersionedSource) GetSourcePath() string {
	return t.sourcePath
}

// GetVersion returns the current version of the entity.
func (t *VersionedSource) GetVersion() int {
	return t.version
}

// setSourcePath validates and sets the template's source path.
// The path must be either a valid local file path or a URL.
// This is an internal method not exported outside the domain.
func (t *VersionedSource) SetSourcePath(path string) error {
	if !IsSyntacticallySafePath(path) || (!IsValidURL(path) && !IsPlausibleLocalPath(path)) {
		return ErrInvalidPathOrUrl
	}
	t.sourcePath = path
	return nil
}

// ForkWithNewVersion creates a deep copy of the current template with an incremented version,
// and assigns it a new source path.
// Returns an error if the new source path is invalid.
func (v *VersionedSource) ForkWithNewVersion(newSourcePath string) (*VersionedSource, error) {
	newVersion := *v
	err := newVersion.SetSourcePath(newSourcePath)
	if err != nil {
		return nil, err
	}
	newVersion.version = v.version + 1
	return &newVersion, nil
}
