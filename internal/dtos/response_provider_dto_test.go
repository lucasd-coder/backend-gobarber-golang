package dtos

import (
	"reflect"
	"testing"
)

func TestNewResponseAllInDayFromProviderDTO(t *testing.T) {
	type args struct {
		hour      int
		available bool
	}
	tests := []struct {
		name string
		args args
		want *ResponseAllInDayFromProviderDTO
	}{
		{
			name: `should return ResponseAllInDayFromProviderDTO`,
			args: args{
				hour:      10,
				available: true,
			},
			want: &ResponseAllInDayFromProviderDTO{
				Hour:      10,
				Available: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewResponseAllInDayFromProviderDTO(tt.args.hour, tt.args.available); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewResponseAllInDayFromProviderDTO() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewResponseAllInMonthFromProviderDTO(t *testing.T) {
	type args struct {
		day       int
		available bool
	}
	tests := []struct {
		name string
		args args
		want *ResponseAllInMonthFromProviderDTO
	}{
		{
			name: `should return ResponseAllInMonthFromProviderDTO`,
			args: args{
				day:       22,
				available: false,
			},
			want: &ResponseAllInMonthFromProviderDTO{
				Day:       22,
				Available: false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewResponseAllInMonthFromProviderDTO(tt.args.day, tt.args.available); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewResponseAllInMonthFromProviderDTO() = %v, want %v", got, tt.want)
			}
		})
	}
}
