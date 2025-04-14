package generator

import (
	"strings"
	"testing"
)

func TestGenerateSlug(t *testing.T) {
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
			got, err := generateSlug(tt.projectName)
			if (err != nil) != tt.wantErr {
				t.Errorf("generateSlug() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got == nil {
				t.Errorf("generateSlug() got nil result")
				return
			}

			if strings.Compare(*got, tt.expectedResult) != 0 {
				t.Errorf("generateSlug() = %v, want %v", *got, tt.expectedResult)
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

func TestGenerateNamespace(t *testing.T) {
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
			expectedResult: "MyCompany\\MyWordpressPlugin",
			wantErr:        false,
		},
		{
			name:           "with-company-2",
			projectName:    "my-company/my-wordpress-plugin",
			expectedResult: "MyCompany\\MyWordpressPlugin",
			wantErr:        false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := generateNamespace(tt.projectName)
			if (err != nil) != tt.wantErr {
				t.Errorf("generateNamespace() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got == nil {
				t.Errorf("generateNamespace() got nil result")
				return
			}

			if strings.Compare(*got, tt.expectedResult) != 0 {
				t.Errorf("generateNamespace() = %v, want %v", *got, tt.expectedResult)
				return
			}
		})
	}
}

func TestGenerateConstantPrefix(t *testing.T) {
	tests := []struct {
		name           string
		projectName    string
		expectedResult string
		wantErr        bool
	}{
		{
			name:           "basic-1",
			projectName:    "MyWordpressPlugin",
			expectedResult: "MY_WORDPRESS_PLUGIN",
			wantErr:        false,
		},
		{
			name:           "basic-2",
			projectName:    "my-wordpress-plugin",
			expectedResult: "MY_WORDPRESS_PLUGIN",
			wantErr:        false,
		},
		{
			name:           "with-company-1",
			projectName:    "MyCompany/MyWordpressPlugin",
			expectedResult: "MY_WORDPRESS_PLUGIN",
			wantErr:        false,
		},
		{
			name:           "with-company-2",
			projectName:    "my-company/my-wordpress-plugin",
			expectedResult: "MY_WORDPRESS_PLUGIN",
			wantErr:        false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := generateConstantPrefix(tt.projectName)
			if (err != nil) != tt.wantErr {
				t.Errorf("generateConstantPrefix() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got == nil {
				t.Errorf("generateConstantPrefix() got nil result")
				return
			}

			if strings.Compare(*got, tt.expectedResult) != 0 {
				t.Errorf("generateConstantPrefix() = %v, want %v", *got, tt.expectedResult)
				return
			}
		})
	}
}

func TestGenerateFunctionPostfix(t *testing.T) {
	tests := []struct {
		name           string
		projectName    string
		expectedResult string
		wantErr        bool
	}{
		{
			name:           "basic-1",
			projectName:    "MyWordpressPlugin",
			expectedResult: "my_wordpress_plugin",
			wantErr:        false,
		},
		{
			name:           "basic-2",
			projectName:    "my-wordpress-plugin",
			expectedResult: "my_wordpress_plugin",
			wantErr:        false,
		},
		{
			name:           "with-company-1",
			projectName:    "MyCompany/MyWordpressPlugin",
			expectedResult: "my_wordpress_plugin",
			wantErr:        false,
		},
		{
			name:           "with-company-2",
			projectName:    "my-company/my-wordpress-plugin",
			expectedResult: "my_wordpress_plugin",
			wantErr:        false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := generateFunctionPostfix(tt.projectName)
			if (err != nil) != tt.wantErr {
				t.Errorf("generateFunctionPostfix() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got == nil {
				t.Errorf("generateFunctionPostfix() got nil result")
				return
			}

			if strings.Compare(*got, tt.expectedResult) != 0 {
				t.Errorf("generateFunctionPostfix() = %v, want %v", *got, tt.expectedResult)
				return
			}
		})
	}
}

func TestGeneratePluginName(t *testing.T) {
	tests := []struct {
		name           string
		projectName    string
		expectedResult string
		wantErr        bool
	}{
		{
			name:           "basic-1",
			projectName:    "MyWordpressPlugin",
			expectedResult: "My Wordpress Plugin",
			wantErr:        false,
		},
		{
			name:           "basic-2",
			projectName:    "my-wordpress-plugin",
			expectedResult: "My Wordpress Plugin",
			wantErr:        false,
		},
		{
			name:           "with-company-1",
			projectName:    "MyCompany/MyWordpressPlugin",
			expectedResult: "My Wordpress Plugin",
			wantErr:        false,
		},
		{
			name:           "with-company-2",
			projectName:    "my-company/my-wordpress-plugin",
			expectedResult: "My Wordpress Plugin",
			wantErr:        false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := generatePluginName(tt.projectName)
			if (err != nil) != tt.wantErr {
				t.Errorf("generatePluginName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got == nil {
				t.Errorf("generatePluginName() got nil result")
				return
			}

			if strings.Compare(*got, tt.expectedResult) != 0 {
				t.Errorf("generatePluginName() = %v, want %v", *got, tt.expectedResult)
				return
			}
		})
	}
}
