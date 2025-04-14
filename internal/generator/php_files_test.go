package generator

import (
	"strings"
	"testing"
)

func TestGenerateTextDomain(t *testing.T) {
	tests := []struct {
		name           string
		projectName    string
		expectedResult string
		wantErr        bool
	}{
		{
			name:           "basic-1",
			projectName:    "MyWordpressPlugin",
			expectedResult: "my-wordpress-plugin",
			wantErr:        false,
		},
		{
			name:           "basic-2",
			projectName:    "my-wordpress-plugin",
			expectedResult: "my-wordpress-plugin",
			wantErr:        false,
		},
		{
			name:           "with-company-1",
			projectName:    "MyCompany/MyWordpressPlugin",
			expectedResult: "my-wordpress-plugin",
			wantErr:        false,
		},
		{
			name:           "with-company-2",
			projectName:    "my-company/my-wordpress-plugin",
			expectedResult: "my-wordpress-plugin",
			wantErr:        false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := generateTextDomain(tt.projectName)
			if (err != nil) != tt.wantErr {
				t.Errorf("generateTextDomain() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got == nil {
				t.Errorf("generateTextDomain() got nil result")
				return
			}

			if strings.Compare(*got, tt.expectedResult) != 0 {
				t.Errorf("generateTextDomain() = %v, want %v", *got, tt.expectedResult)
				return
			}
		})
	}
}

func TestGenerateClassPrefix(t *testing.T) {
	tests := []struct {
		name           string
		projectName    string
		expectedResult string
		wantErr        bool
	}{
		{
			name:           "basic-1",
			projectName:    "MyWordpressPlugin",
			expectedResult: "MyWordpressPlugin",
			wantErr:        false,
		},
		{
			name:           "basic-2",
			projectName:    "my-wordpress-plugin",
			expectedResult: "MyWordpressPlugin",
			wantErr:        false,
		},
		{
			name:           "with-company-1",
			projectName:    "MyCompany/MyWordpressPlugin",
			expectedResult: "MyWordpressPlugin",
			wantErr:        false,
		},
		{
			name:           "with-company-2",
			projectName:    "my-company/my-wordpress-plugin",
			expectedResult: "MyWordpressPlugin",
			wantErr:        false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := generateClassPrefix(tt.projectName)
			if (err != nil) != tt.wantErr {
				t.Errorf("generateClassPrefix() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got == nil {
				t.Errorf("generateClassPrefix() got nil result")
				return
			}

			if strings.Compare(*got, tt.expectedResult) != 0 {
				t.Errorf("generateClassPrefix() = %v, want %v", *got, tt.expectedResult)
				return
			}
		})
	}
}
