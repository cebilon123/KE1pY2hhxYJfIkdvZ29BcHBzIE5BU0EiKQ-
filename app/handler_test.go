package app

import (
	"testing"
	"time"
)

func TestAreDatesValid(t *testing.T) {
	type args struct {
		startDate []time.Time
		endDate   []time.Time
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
				startDate: []time.Time{time.Date(2021,8,1,0,0,0,0, time.UTC)},
				endDate:   []time.Time{time.Date(2021,8,5,0,0,0,0, time.UTC)},
			},
			want:  true,
			want1: "",
		},
		{
			name:  "start_date after end_date",
			args:  args{
				startDate: []time.Time{time.Date(2021,8,5,0,0,0,0, time.UTC)},
				endDate:   []time.Time{time.Date(2021,8,1,0,0,0,0, time.UTC)},
			},
			want:  false,
			want1: "validation error: start_date should be earlier than end_date",
		},
		{
			name:  "start_date and end_date empty",
			args:  args{
				startDate: []time.Time{},
				endDate:   []time.Time{},
			},
			want:  false,
			want1: "start_date or end_date need to passed as query parameters",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := AreDatesValid(tt.args.startDate, tt.args.endDate)
			if got != tt.want {
				t.Errorf("AreDatesValid() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("AreDatesValid() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
