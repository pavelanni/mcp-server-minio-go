package fsutils

import (
	"os"
	"path/filepath"
	"testing"
)

func TestNormalizePath(t *testing.T) {
	// Get user's home directory for testing
	home, err := os.UserHomeDir()
	if err != nil {
		t.Fatalf("Failed to get user home directory: %v", err)
	}

	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Home directory expansion",
			input:    "~/documents",
			expected: filepath.Join(home, "documents"),
		},
		{
			name:     "Regular path",
			input:    "/usr/local/bin",
			expected: "/usr/local/bin",
		},
		{
			name:     "Path with dots",
			input:    "./test/../folder",
			expected: "folder",
		},
		{
			name:     "Just tilde",
			input:    "~",
			expected: home,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := NormalizePath(tt.input)
			if result != tt.expected {
				t.Errorf("NormalizePath(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestValidatePath(t *testing.T) {
	allowedPaths := []string{
		"/allowed/path",
		"/another/allowed",
		filepath.Join(os.TempDir(), "test"),
	}

	tests := []struct {
		name        string
		input       string
		shouldError bool
	}{
		{
			name:        "Allowed path",
			input:       "/allowed/path/subdir",
			shouldError: false,
		},
		{
			name:        "Another allowed path",
			input:       "/another/allowed/file.txt",
			shouldError: false,
		},
		{
			name:        "Disallowed path",
			input:       "/not/allowed/path",
			shouldError: true,
		},
		{
			name:        "Empty path",
			input:       "",
			shouldError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidatePath(tt.input, allowedPaths)
			if tt.shouldError && err == nil {
				t.Errorf("ValidatePath(%q) should have returned an error", tt.input)
			}
			if !tt.shouldError && err != nil {
				t.Errorf("ValidatePath(%q) returned unexpected error: %v", tt.input, err)
			}
		})
	}
}
