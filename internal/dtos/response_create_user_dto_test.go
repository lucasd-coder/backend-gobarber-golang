package dtos

import (
	"reflect"
	"testing"
)

func TestNewResponseCreateUserDTO(t *testing.T) {
	type args struct {
		name  string
		email string
	}
	tests := []struct {
		name string
		args args
		want *ResponseCreateUserDTO
	}{
		{
			name: `should return ResponseCreateUserDTO`,
			args: args{
				name:  "Lindalva",
				email: "Lindalva@gmail.com",
			},
			want: &ResponseCreateUserDTO{
				Name:  "Lindalva",
				Email: "Lindalva@gmail.com",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewResponseCreateUserDTO(tt.args.name, tt.args.email); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewResponseCreateUserDTO() = %v, want %v", got, tt.want)
			}
		})
	}
}
