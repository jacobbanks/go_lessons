package integers

import (
	"testing"
	"fmt"
)


func TestAdder(t *testing.T) {
	sum := Add(2, 4)
	got := 6
	if got != sum {
		t.Errorf("expected '%d', got '%d'", got, sum)
	}
}
// This is called a testable example see docs here: https://blog.golang.org/examples
func ExampleAdd() {
	sum := Add(2, 9)
	fmt.Println(sum)
	// Output: 11
}
