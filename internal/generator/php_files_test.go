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
			name:           "basic",
			projectName:    "MyWordpressPlugin",
			expectedResult: "my-wordpress-plugin",
			wantErr:        false,
		},
		{
			name:           "with-company",
			projectName:    "MyCompany/MyWordpressPlugin",
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
