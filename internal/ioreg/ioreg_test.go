package ioreg

import (
	"errors"
	"testing"
	"time"

	"github.com/florian42/me-api/internal/cmd"
)

type mockRunner struct {
	output []byte
	err    error
}

func (m mockRunner) Run(name string, args ...string) ([]byte, error) {
	return m.output, m.err
}

func TestGetIdleTime(t *testing.T) {
	tests := []struct {
		name       string
		runner     cmd.CommandRunner
		want       time.Duration
		wantErr    bool
		errContain string
	}{
		{
			name:   "successful parse",
			runner: mockRunner{output: []byte(`"HIDIdleTime" = 1234567890000`)},
			want:   time.Duration(1234567890000),
		},
		{
			name:       "no match found",
			runner:     mockRunner{output: []byte(`some other output`)},
			wantErr:    true,
			errContain: "HIDIdleTime not found",
		},
		{
			name:       "command fails",
			runner:     mockRunner{err: errors.New("command not found")},
			wantErr:    true,
			errContain: "command not found",
		},
		{
			name:   "with spacing variations",
			runner: mockRunner{output: []byte(`  "HIDIdleTime"  =  5000000000  `)},
			want:   time.Duration(5000000000),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetIdleTime(tt.runner)

			if tt.wantErr {
				if err == nil {
					t.Errorf("GetIdleTime() error = nil, wantErr %v", tt.wantErr)
					return
				}
				if tt.errContain != "" && !contains(err.Error(), tt.errContain) {
					t.Errorf("GetIdleTime() error = %v, should contain %v", err, tt.errContain)
				}
				return
			}

			if err != nil {
				t.Errorf("GetIdleTime() unexpected error = %v", err)
				return
			}

			if got != tt.want {
				t.Errorf("GetIdleTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func contains(s, substr string) bool {
	return len(substr) <= len(s) && (s == substr || len(s) > 0 && containsSubstring(s, substr))
}

func containsSubstring(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
