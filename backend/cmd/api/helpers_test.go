package main

import (
	"bytes"
	"mime/multipart"
	"net/http/httptest"
	"testing"
)

func TestValidateImageFile(t *testing.T) {
	app := &application{}

	tests := []struct {
		name        string
		content     []byte
		filename    string
		expectError bool
	}{
		{
			name:        "Valid PNG",
			content:     []byte{0x89, 0x50, 0x4E, 0x47, 0x0D, 0x0A, 0x1A, 0x0A}, // PNG Magic Number
			filename:    "test.png",
			expectError: false,
		},
		{
			name:        "Valid JPEG",
			content:     []byte{0xFF, 0xD8, 0xFF, 0xE0}, // partial JPEG Magic Number
			filename:    "test.jpg",
			expectError: false,
		},
		{
			name:        "Valid GIF (Should Fail - Not Allowed)",
			content:     []byte("GIF89a"),
			filename:    "test.gif",
			expectError: true,
		},
		{
			name:        "Text File disguised as PNG",
			content:     []byte("This is a text file masquerading as an image."),
			filename:    "malicious.png",
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a buffer to write our multipart form data
			body := new(bytes.Buffer)
			writer := multipart.NewWriter(body)

			// Create form file
			part, err := writer.CreateFormFile("file", tt.filename)
			if err != nil {
				t.Fatal(err)
			}

			// Write content to the part
			// We need enough bytes for detection (at least 512 normally, but DetecContentType handles fewer)
			// But let's pad it just in case
			paddedContent := make([]byte, 512)
			copy(paddedContent, tt.content)

			part.Write(paddedContent)
			writer.Close()

			// Create a request with the body
			req := httptest.NewRequest("POST", "/", body)
			req.Header.Set("Content-Type", writer.FormDataContentType())

			// Parse the multipart form
			err = req.ParseMultipartForm(10 << 20)
			if err != nil {
				t.Fatal(err)
			}

			// Get the file header
			_, header, err := req.FormFile("file")
			if err != nil {
				t.Fatal(err)
			}

			// Test the function
			err = app.ValidateImageFile(header)

			if tt.expectError && err == nil {
				t.Errorf("Expected error but got nil")
			}
			if !tt.expectError && err != nil {
				t.Errorf("Expected no error but got: %v", err)
			}
		})
	}
}
