package countpositivessumnegatives

import (
	"reflect"
	"testing"
)

func TestCountPositivesSumNegatives(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want []int
	} {
		{ name: "Only Positives", args: []int{1, 2}, want: []int{2,0} },
		{ name: "Only One Test Case", args: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, -11, -12, -13, -14, -15}, want: []int{10, -65}},
		{ name: "Only Negative", args: []int{-1, -2}, want: []int{0, -3} },
		{ name: "Only Zero", args: []int{0}, want: []int{1,0} },
	}

	for _, tt := range tests {
		t.Run(tt.name, func (t *testing.T) {
			t.Cleanup(func() {
				t.Log("clean up !")
			})

			defer t.Log("defer!")

			if data := CountPositivesSumNegatives(tt.args); !reflect.DeepEqual(data, tt.want) {
				t.Fatalf("CountPositivesSumNegatives() = %v, want %v", data, tt.want)
			}
			t.Log("after CountPositivesSumNegatives() ...")
		})
	}
}