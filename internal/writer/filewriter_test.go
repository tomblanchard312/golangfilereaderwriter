package writer

import (
	"os"
	"testing"
)

func TestWriteOrAppendToFile(t *testing.T) {
	content := "hello\ngo\n"
	tmpFile, err := os.CreateTemp("", "example")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpFile.Name())

	err = WriteOrAppendToFile(tmpFile.Name(), content)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	result, err := os.ReadFile(tmpFile.Name())
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if string(result) != content {
		t.Fatalf("expected %s, got %s", content, string(result))
	}

	// Append content
	appendContent := "appended text\n"
	err = WriteOrAppendToFile(tmpFile.Name(), appendContent)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	expectedContent := content + appendContent
	result, err = os.ReadFile(tmpFile.Name())
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if string(result) != expectedContent {
		t.Fatalf("expected %s, got %s", expectedContent, string(result))
	}
}
