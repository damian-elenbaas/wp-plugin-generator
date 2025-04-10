package generator

import (
	"encoding/json"
	"testing"
)

func TestGenerateComposerJson(t *testing.T) {
	tests := []struct {
		name              string
		projectName       string
		authors           []Author
		expectedNamespace string
		wantErr           bool
	}{
		{
			name:        "basic",
			projectName: "project",
			authors: []Author{
				{Name: "John Doe", Email: "john@example.com"},
			},
			expectedNamespace: "Project\\",
			wantErr:           false,
		},
		{
			name:        "advanced project name",
			projectName: "company/my-project",
			authors: []Author{
				{Name: "John Doe", Email: "john@example.com"},
			},
			expectedNamespace: "Company\\MyProject\\",
			wantErr:           false,
		},
		{
			name:        "multiple authors",
			projectName: "project",
			authors: []Author{
				{Name: "John Doe", Email: "john@example.com"},
				{Name: "Adam Doe", Email: "adam@example.com"},
			},
			expectedNamespace: "Project\\",
			wantErr:           false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := generateComposerJson(tt.projectName, tt.authors)

			if (err != nil) != tt.wantErr {
				t.Errorf("generateComposerJson() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			var result map[string]any
			if err := json.Unmarshal([]byte(got), &result); err != nil {
				t.Errorf("generateComposerJson() produced invalid JSON: %v", err)
				return
			}

			// Check project name
			name, ok := result["name"].(string)
			if !ok || name != tt.projectName {
				t.Errorf("generateComposerJson() name mismatch: expected %s, got %s", tt.projectName, name)
				return
			}

			// Check type
			typ, ok := result["type"].(string)
			if !ok || typ != "project" {
				t.Errorf("generateComposerJson() type mismatch: expected 'project', got %s", typ)
				return
			}

			// Check namespace in autoload psr-4
			autoload, ok := result["autoload"].(map[string]any)
			if !ok {
				t.Error("generateComposerJson() autoload field is missing or invalid")
				return
			}

			psr4, ok := autoload["psr-4"].(map[string]any)
			if !ok {
				t.Error("generateComposerJson() psr-4 field is missing or invalid")
				return
			}

			_, ok = psr4[tt.expectedNamespace]
			if !ok {
				t.Errorf("generateComposerJson() namespace %s is missing", tt.expectedNamespace)
				return
			}

			value, ok := psr4[tt.expectedNamespace].(string)
			if !ok || value != "includes/" {
				t.Errorf("generateComposerJson() namespace %s should map to 'includes/', got %v", tt.expectedNamespace, value)
				return
			}

			// Check authors
			authors, ok := result["authors"].([]any)
			if !ok {
				t.Error("generateComposerJson() authors field is missing or invalid")
				return
			}

			if len(authors) != len(tt.authors) {
				t.Errorf("generateComposerJson() expected %d authors, got %d", len(tt.authors), len(authors))
				return
			}

			for i, author := range authors {
				authorMap, ok := author.(map[string]any)
				if !ok {
					t.Errorf("generateComposerJson() author %d is not a map", i)
					continue
				}

				name, ok := authorMap["name"].(string)
				if !ok || name != tt.authors[i].Name {
					t.Errorf("generateComposerJson() author %d name mismatch: expected %s, got %s", i, tt.authors[i].Name, name)
					continue
				}

				email, ok := authorMap["email"].(string)
				if !ok || email != tt.authors[i].Email {
					t.Errorf("generateComposerJson() author %d email mismatch: expected %s, got %s", i, tt.authors[i].Email, email)
				}
			}

			// Check dev dependencies
			requireDev, ok := result["require-dev"].(map[string]any)
			if !ok {
				t.Error("generateComposerJson() require-dev field is missing or invalid")
				return
			}

			value, ok = requireDev["php-stubs/wordpress-stubs"].(string)
			if !ok {
				t.Error("generateComposerJson() php-stubs/wordpress-stubs should be a dependency")
				return
			}
		})
	}
}
