package sol

import "testing"

func BenchmarkTest(b *testing.B) {
	intervals := []*Interval{{0, 30}, {5, 10}, {15, 20}}
	for idx := 0; idx < b.N; idx++ {
		MinMeetingRooms(intervals)
	}
}
func TestMinMeetingRooms(t *testing.T) {
	type args struct {
		intervals []*Interval
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "intervals = [(0,30),(5,10),(15,20)]",
			args: args{intervals: []*Interval{{0, 30}, {5, 10}, {15, 20}}},
			want: 2,
		},
		{
			name: "intervals = [(2,7)]",
			args: args{intervals: []*Interval{{2, 7}}},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MinMeetingRooms(tt.args.intervals); got != tt.want {
				t.Errorf("MinMeetingRooms() = %v, want %v", got, tt.want)
			}
		})
	}
}
