package tosec_test

import (
	"github.com/climbus/retro-romkit/pkg/tosec"
	"testing"
)

func TestGetPlatform(t *testing.T) {
	tests := []struct {
		name     string
		platform string
		want     tosec.Platform
		error    bool
	}{
		{
			name:     "get existing platform",
			platform: "nes",
			want: tosec.Platform{
				Name:        "nes",
				Description: "Nintendo Entertainment System",
				FileTypes:   []string{".nes", ".fds"},
			},
			error: false,
		},
		{
			name:     "get non-existing platform",
			platform: "unknown",
			want:     tosec.Platform{},
			error:    true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tosec.GetPlatform(tt.platform)

			if err != tt.error {
				t.Errorf("GetPlatform() error = %v, wantErr %v", err, tt.error)
				return
			}

			if got.Description != tt.want.Description || got.Name != tt.want.Name || len(got.FileTypes) != len(tt.want.FileTypes) {
				t.Errorf("GetPlatform() got = %v, want %v", got, tt.want)
			}

		})
	}
}
