package reader

import (
	"os"
	"testing"
)

func TestReadFile(t *testing.T) {
	content := "hello\ngo\n"
	tmpFile, err := os.CreateTemp("", "example")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpFile.Name())

	if _, err := tmpFile.Write([]byte(content)); err != nil {
		t.Fatal(err)
	}
	if err := tmpFile.Close(); err != nil {
		t.Fatal(err)
	}

	result, err := ReadFile(tmpFile.Name())
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if result != content {
		t.Fatalf("expected %s, got %s", content, result)
	}
}
