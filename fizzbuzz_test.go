package fizzbuzz

import (
	"fmt"
	"testing"
)

func TestFizzBuzz01(t *testing.T) {
	want := "1\n2\nFizz\n4\nBuzz\nFizz\n7\n8\nFizz\nBuzz\n11\nFizz\n13\n14\nFizzBuzz\n16\n17\nFizz\n19\nBuzz\n"
	got := ""
	for s := range FizzBuzz(20) {
		got += fmt.Sprintf("%v%v", s, "\n")
	}
	if got != want {
		t.Errorf("FizzBuzz(15) \n got: \n%v \n want: \n%v", got, want)
	}
}

func TestFizzBuzz02(t *testing.T) {
	Fizz = "foo"
	Buzz = "bar"

	want := "1\n2\nfoo\n4\nbar\nfoo\n7\n8\nfoo\nbar\n11\nfoo\n13\n14\nfoobar\n16\n17\nfoo\n19\nbar\n"
	got := ""
	for s := range FizzBuzz(20) {
		got += fmt.Sprintf("%v%v", s, "\n")
	}

	if got != want {
		t.Errorf("FizzBuzz(15) \n got: \n%v \n want: \n%v", got, want)
	}
}
