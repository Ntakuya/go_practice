package evenorodd

import (
	"testing"
)

func TestEvenOrOdd(t *testing.T) {
	tests := []struct {
		name string
		arg int
		want string
	}{
		{ name: "Even", arg: 2 , want: "Even" },
		{ name: "Odd", arg: 1, want: "Odd" },
		{ name: "Zero", arg: 0, want: "Even" },
		{ name: "minus Odd", arg: -1, want: "Odd" },
		{ name: "minus Even", arg: -2, want: "Even" },
	}

	for _, tt := range tests {
		t.Run(tt.name, func (t *testing.T) {
			t.Cleanup(func() {
				t.Log("clean up !")
			})

			defer t.Log("defer!")

			if got := EvenOrOdd(tt.arg); got != tt.want {
				t.Fatalf("EvenOrOdd() = %v, want %v", got, tt.want)
			}
			t.Log("after EvenOrOdd() ...")
		})
	}
}