package lsappinfo

import (
	"errors"
	"fmt"
	"strings"
	"testing"

	"github.com/florian42/me-api/internal/cmd"
)

type sequentialMockRunner struct {
	responses []mockResponse
	callCount int
}

type mockResponse struct {
	output []byte
	err    error
}

func (m *sequentialMockRunner) Run(name string, args ...string) ([]byte, error) {
	if m.callCount >= len(m.responses) {
		return nil, fmt.Errorf("unexpected call #%d", m.callCount+1)
	}
	resp := m.responses[m.callCount]
	m.callCount++
	return resp.output, resp.err
}

func TestGetFrontmostAppName(t *testing.T) {
	tests := []struct {
		name       string
		runner     cmd.CommandRunner
		want       string
		wantErr    bool
		errContain string
	}{
		{
			name: "successful parse",
			runner: &sequentialMockRunner{
				responses: []mockResponse{
					{output: []byte("ASN:0x0-0x12345:com.apple.Safari")}, // lsappinfo front
					{output: []byte(`"CFBundleDisplayName"="Safari"`)},   // lsappinfo info -only name ...
				},
			},
			want: "Safari",
		},
		{
			name: "first command fails",
			runner: &sequentialMockRunner{
				responses: []mockResponse{
					{err: errors.New("command not found")},
				},
			},
			wantErr:    true,
			errContain: "failed to get frontmost app",
		},
		{
			name: "second command fails",
			runner: &sequentialMockRunner{
				responses: []mockResponse{
					{output: []byte("ASN:0x0-0x12345:com.apple.Safari")},
					{err: errors.New("permission denied")},
				},
			},
			wantErr:    true,
			errContain: "failed to get app info",
		},
		{
			name: "empty ASN",
			runner: &sequentialMockRunner{
				responses: []mockResponse{
					{output: []byte("")},
				},
			},
			wantErr:    true,
			errContain: "no frontmost app found",
		},
		{
			name: "cannot parse name",
			runner: &sequentialMockRunner{
				responses: []mockResponse{
					{output: []byte("ASN:0x0-0x12345:com.test.app")},
					{output: []byte("invalid output")},
				},
			},
			wantErr:    true,
			errContain: "could not parse app name",
		},
		{
			name: "alternate key name LSDisplayName",
			runner: &sequentialMockRunner{
				responses: []mockResponse{
					{output: []byte("ASN:0x0-0x12345:com.test.app")},
					{output: []byte(`"LSDisplayName"="TestApp"`)},
				},
			},
			want: "TestApp",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetFrontmostAppName(tt.runner)

			if tt.wantErr {
				if err == nil {
					t.Errorf("GetFrontmostAppName() error = nil, wantErr %v", tt.wantErr)
					return
				}
				if tt.errContain != "" && !strings.Contains(err.Error(), tt.errContain) {
					t.Errorf("GetFrontmostAppName() error = %v, should contain %v", err, tt.errContain)
				}
				return
			}

			if err != nil {
				t.Errorf("GetFrontmostAppName() unexpected error = %v", err)
				return
			}

			if got != tt.want {
				t.Errorf("GetFrontmostAppName() = %v, want %v", got, tt.want)
			}
		})
	}
}
