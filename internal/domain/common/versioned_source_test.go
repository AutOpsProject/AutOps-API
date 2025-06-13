package common

import "testing"

func TestNewVersionedSourceErr(t *testing.T) {
	path := "Invalid file path"
	_, err := NewVersionedSource(path, 1)
	if err != ErrInvalidPathOrUrl {
		t.Error("expected err to be ErrInvalidPathOrUrl")
	}
}

func TestForkTemplate(t *testing.T) {
	path := "/path/to/file.zip"
	versionedSource, err := NewVersionedSource(path, 18)
	if err != nil {
		t.Error("expected err to be nil")
	} else if versionedSource.GetVersion() != 18 {
		t.Errorf("expected version to be 18, got %d", versionedSource.GetVersion())
	} else if versionedSource.GetSourcePath() != path {
		t.Errorf("expected path to be %s, got %s", path, versionedSource.sourcePath)
	}

	newVersion, err := versionedSource.ForkWithNewVersion("/new/path/to/file")
	if err != nil {
		t.Error("expected err to be nil")
	} else if newVersion.GetVersion() != versionedSource.GetVersion()+1 {
		t.Errorf("expected new version to be %d, got %d", versionedSource.GetVersion()+1, newVersion.GetVersion())
	}

	_, err = versionedSource.ForkWithNewVersion("invalid file path")
	if err != ErrInvalidPathOrUrl {
		t.Error("expected err to be ErrInvalidPathOrUrl")
	}
}
