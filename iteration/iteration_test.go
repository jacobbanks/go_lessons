package iteration

import (
	"fmt"
	"testing"
)


func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 2999)
	}
}


func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 5)

	expected := "aaaaa"

	if repeated != expected {
	  t.Errorf("got %q, wanted %q", repeated, expected)
	}

}

func ExampleRepeat() {
	res := Repeat("b", 5)
	fmt.Println(res)
	// Output: bbbbb
}
