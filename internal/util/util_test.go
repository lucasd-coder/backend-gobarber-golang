package util

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"

	"github.com/lucasd-coder/backend-gobarber-golang/internal/dtos"
)

var (
	layout    = "2006-01-02 15:04:05"
	dateUTC   = time.Date(2022, time.December, 20, 10, 0, 0, 0, time.UTC)
	dateLocal = time.Date(2022, time.December, 20, 10, 0, 676763, 988220, time.Local)
	body      = bytes.NewBufferString("{\"date\":\"2022-12-20T10:00:00Z\",\"provider_id\":\"dd56ad45-9076-4772-b4d3-d987561f7095\"}")
)

var responseRecorder = &httptest.ResponseRecorder{
	Code: 200,
	Body: body,
}

func TestJsonLog(t *testing.T) {
	type args struct {
		payload interface{}
	}
	tests := []struct {
		name string
		args args
		want *bytes.Buffer
	}{
		{
			name: `should return bytes.Buffer`,
			args: args{
				payload: dtos.AppointmentDTO{
					Date:       dateUTC,
					ProviderID: "dd56ad45-9076-4772-b4d3-d987561f7095",
				},
			},
			want: body,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := JsonLog(tt.args.payload); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("JsonLog() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsValidUUID(t *testing.T) {
	type args struct {
		u string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: `should return false when UUID is invalid `,
			args: args{
				u: "Invalid UUID",
			},
			want: false,
		},
		{
			name: `should return true when UUID is valid`,
			args: args{
				u: "dd56ad45-9076-4772-b4d3-d987561f7095",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValidUUID(tt.args.u); got != tt.want {
				t.Errorf("IsValidUUID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDateUtils(t *testing.T) {
	resultParse, _ := time.Parse(layout, dateUTC.Format(layout))

	type args struct {
		aux    time.Time
		layOut string
	}
	tests := []struct {
		name    string
		args    args
		want    *time.Time
		wantErr bool
	}{
		{
			name: `should return time parses a formatted`,
			args: args{
				aux:    dateUTC,
				layOut: layout,
			},
			want:    &resultParse,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DateUtils(tt.args.aux, tt.args.layOut)
			if (err != nil) != tt.wantErr {
				t.Errorf("DateUtils() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DateUtils() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDateFormat(t *testing.T) {
	type args struct {
		date   time.Time
		layout string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "should return date formatted",
			args: args{
				date:   dateUTC,
				layout: layout,
			},
			want: dateUTC.Format(layout),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DateFormat(tt.args.date, tt.args.layout); got != tt.want {
				t.Errorf("DateFormat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHashPassword(t *testing.T) {
	pass := "mypassword"
	hp, err := HashPassword(pass)
	if err != nil {
		t.Fatalf("Generate HashPassword error: %s", err)
	}

	if !CheckPasswordHash(pass, hp) {
		t.Errorf("%v should hash %s correctly", pass, hp)
	}
}

func TestCheckPasswordHash(t *testing.T) {
	type args struct {
		password string
		hash     string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "should return true when password for corresponding hash",
			args: args{
				password: "12345678",
				hash:     "$2a$08$suhsGAQD9aFzS7Ur5wAv0OoQNNKaQL2dUe7xn3AqwS2m9sI2wK646",
			},
			want: true,
		},
		{
			name: "should return false when password for not corresponding hash",
			args: args{
				password: "1234567891011",
				hash:     "$2a$08$suhsGAQD9aFzS7Ur5wAv0OoQNNKaQL2dUe7xn3AqwS2m9sI2wK646",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckPasswordHash(tt.args.password, tt.args.hash); got != tt.want {
				t.Errorf("CheckPasswordHash() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsAfter(t *testing.T) {
	type args struct {
		compareDate time.Time
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "should return true when date is after date now",
			args: args{
				compareDate: time.Now().Truncate(time.Hour),
			},
			want: true,
		},
		{
			name: "should return false when date not after is date now",
			args: args{
				compareDate: time.Now().Add(time.Hour),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsAfter(tt.args.compareDate); got != tt.want {
				t.Errorf("IsAfter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseFromHttpResponse(t *testing.T) {
	type args struct {
		resp  *http.Response
		model interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "",
			args: args{
				resp:  responseRecorder.Result(),
				model: &dtos.AppointmentDTO{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ParseFromHttpResponse(tt.args.resp, tt.args.model); (err != nil) != tt.wantErr {
				t.Errorf("ParseFromHttpResponse() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
