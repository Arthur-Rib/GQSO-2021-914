package main

import "testing"

func Test_soma(t *testing.T) {
	type args struct {
		x float64
		y float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := soma(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("soma() = %v, want %v", got, tt.want)
			}
		})
	}
}
