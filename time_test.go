package go_tilities

import (
	"testing"
	"time"
)

func TestGetLast7DaysTSWithNow(t *testing.T) {
	type args struct {
		timestamp int64
		offset    int64
	}
	tests := []struct {
		name  string
		args  args
		want  int64
		want1 int64
	}{
		{
			name:  "1",
			args:  args{timestamp: time.Now().Unix(), offset: 25200},
			want:  1672592400,
			want1: 1673197200,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := GetLast7DaysTSWithNow(tt.args.timestamp, tt.args.offset)
			if got != tt.want {
				t.Errorf("GetLast7DaysTSWithNow() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("GetLast7DaysTSWithNow() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestDayStartEndDate(t *testing.T) {
	type args struct {
		timestamp int64
		timeZone  *time.Location
	}
	tests := []struct {
		name  string
		args  args
		want  int64
		want1 int64
	}{
		// TODO: Add test cases.
		{
			name:  "1",
			args:  args{timestamp: time.Now().Unix()},
			want:  0,
			want1: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := DayStartEndDate(tt.args.timestamp, tt.args.timeZone)
			if got != tt.want {
				t.Errorf("DayStartEndDate() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("DayStartEndDate() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestGetDayTSWithDate(t *testing.T) {
	type args struct {
		date     string
		timeZone string
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		// TODO: Add test cases.
		{
			name: "1",
			args: args{date: "2023-1-1"},
			want: 1672617600,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetDayTSWithDate(tt.args.date, tt.args.timeZone); got != tt.want {
				t.Errorf("GetDayTSWithDate() = %v, want %v", got, tt.want)
			}
		})
	}
}
