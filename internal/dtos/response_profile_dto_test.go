package dtos

import (
	"reflect"
	"testing"
	"time"
)

var createdAt = time.Now()

func TestNewResponseProfileDTO(t *testing.T) {
	type args struct {
		id        string
		name      string
		email     string
		avatar    string
		createdAt time.Time
		updatedAt time.Time
	}
	tests := []struct {
		name string
		args args
		want *ResponseProfileDTO
	}{
		{
			name: `should return ResponseProfileDTO`,
			args: args{
				id:        "b665ed26-8ae4-407f-a667-9bb093431caf",
				name:      "Lindalva",
				email:     "Lindalva@gmail.com",
				createdAt: createdAt,
			},
			want: &ResponseProfileDTO{
				ID:        "b665ed26-8ae4-407f-a667-9bb093431caf",
				Name:      "Lindalva",
				Email:     "Lindalva@gmail.com",
				CreatedAt: createdAt,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewResponseProfileDTO(tt.args.id, tt.args.name, tt.args.email, tt.args.avatar, tt.args.createdAt, tt.args.updatedAt); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewResponseProfileDTO() = %v, want %v", got, tt.want)
			}
		})
	}
}
