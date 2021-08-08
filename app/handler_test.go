package app

import (
	"testing"
	"time"
)

func TestAreDatesValid(t *testing.T) {
	type args struct {
		startDate time.Time
		endDate   time.Time
	}
	tests := []struct {
		name  string
		args  args
		want  bool
		want1 string
	}{
		{
			name:  "start_date and end_date valid",
			args:  args{
				startDate: time.Date(2021,8,1,0,0,0,0, time.UTC),
				endDate:   time.Date(2021,8,5,0,0,0,0, time.UTC),
			},
			want:  true,
			want1: "",
		},
		{
			name:  "start_date after end_date",
			args:  args{
				startDate: time.Date(2021,8,5,0,0,0,0, time.UTC),
				endDate:   time.Date(2021,8,1,0,0,0,0, time.UTC),
			},
			want:  false,
			want1: "validation error: start_date should be earlier than end_date",
		},
		{
			name:  "start_date before 2015-01-01",
			args:  args{
				startDate: time.Date(2014,8,5,0,0,0,0, time.UTC),
				endDate:   time.Date(2021,8,1,0,0,0,0, time.UTC),
			},
			want:  false,
			want1: "validation error: start_date should be greater or equal to 2015-01-01",
		},
		{
			name:  "start_date should be before today",
			args:  args{
				startDate: time.Date(2019,8,5,0,0,0,0, time.UTC),
				endDate:   time.Date(2021,8,1,0,0,0,0, time.UTC),
			},
			want:  true,
			want1: "",
		},
		{
			name:  "start_date should be before today",
			args:  args{
				startDate: time.Now().AddDate(20,0,0),
				endDate:   time.Date(3120,8,1,0,0,0,0, time.UTC),
			},
			want:  false,
			want1: "validation error: start_date should be at least today",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := areDatesValid(tt.args.startDate, tt.args.endDate)
			if got != tt.want {
				t.Errorf("areDatesValid() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("areDatesValid() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
