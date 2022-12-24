package dtos

import (
	"reflect"
	"testing"
	"time"
)

var resp = ResponseProfileDTO{
	ID:        "b665ed26-8ae4-407f-a667-9bb093431caf",
	Name:      "Lindalva",
	Email:     "Lindalva@gmail.com",
	CreatedAt: time.Now(),
}

var token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"

func TestNewResponseUserAuthenticatedSuccessDTO(t *testing.T) {
	type args struct {
		response ResponseProfileDTO
		token    string
	}
	tests := []struct {
		name string
		args args
		want *ResponseUserAuthenticatedSuccessDTO
	}{
		{
			name: `should return ResponseUserAuthenticatedSuccessDTO`,
			args: args{
				response: resp,
				token:    token,
			},
			want: &ResponseUserAuthenticatedSuccessDTO{
				Response: resp,
				Token:    token,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewResponseUserAuthenticatedSuccessDTO(tt.args.response, tt.args.token); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewResponseUserAuthenticatedSuccessDTO() = %v, want %v", got, tt.want)
			}
		})
	}
}
