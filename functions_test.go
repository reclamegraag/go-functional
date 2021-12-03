package functions

import (
	"fmt"
	"math"
	"reflect"
	. "strings"
	"testing"
)

func TestStringSliceMap(test *testing.T) {
	fmt.Printf("\n\tGIVEN\ta String Slice that contains a, b, c")
	fmt.Printf("\n\tWHEN\tthe String Slice values are mapped with the ToUpper function")
	fmt.Printf("\n\tTHEN\t\tthe new values are A, B, C\n\n")
	stringSlice := StringSlice{"a", "b", "c"}
	mappedStringSlice := stringSlice.Map(func(value string) string {
		return ToUpper(value)
	})
	referenceStringSlice := StringSlice{"A", "B", "C"}
	if !reflect.DeepEqual(mappedStringSlice, referenceStringSlice) {
		test.Errorf("\n\nTEST FAILED:\n\tgot:\t\t\t%s\n\twanted:\t%s\n\n", Join(mappedStringSlice[:], ", "), Join(referenceStringSlice[:], ", "))
	}
}

func TestIntSliceMap(test *testing.T) {
	fmt.Printf("\n\tGIVEN\tan Int Slice that contains 1, 2, 3")
	fmt.Printf("\n\tWHEN\tthe Int Slice values are mapped with the Pow function")
	fmt.Printf("\n\tTHEN\t\tthe new values are 1, 4, 9\n\n")
	intSlice := IntSlice{1, 2, 3}
	mappedIntSlice := intSlice.Map(func(i int) int {
		return int(math.Pow(float64(i), 2))
	})
	referenceIntSlice := IntSlice{1, 4, 9}
	if !reflect.DeepEqual(mappedIntSlice, referenceIntSlice) {
		got := Trim(Replace(fmt.Sprint(mappedIntSlice), " ", ", ", -1), "[]")
		wanted := Trim(Replace(fmt.Sprint(referenceIntSlice), " ", ", ", -1), "[]")
		test.Errorf("\n\nTEST FAILED:\n\tgot:\t\t\t%s\n\twanted:\t%s\n\n", got, wanted)
	}
}

func TestStringSliceMkString(test *testing.T) {
	fmt.Printf("\n\tGIVEN\tan String Slice that contains a, b, c AND --- as the wanted separator")
	fmt.Printf("\n\tWHEN\tthe String Slice values are combined into one string")
	fmt.Printf("\n\tTHEN\t\tthe resulting string is a---b---c\n\n")
	stringSlice := StringSlice{"a", "b", "c"}
	stringFromStringSlice := stringSlice.MkString("---")
	referenceStringFromStringSlice := "a---b---c"
	if stringFromStringSlice != referenceStringFromStringSlice {
		test.Errorf("\n\nTEST FAILED:\n\tgot:\t\t\t%s\n\twanted:\t%s\n\n", stringFromStringSlice, referenceStringFromStringSlice)
	}
}

func TestStringSliceMkStringWithEnds(test *testing.T) {
	fmt.Printf("\n\tGIVEN\ta String Slice that contains a, b, c AND --- as the wanted separator AND [( as the start string AND )] as the end string")
	fmt.Printf("\n\tWHEN\tthe String Slice values are combined into one string")
	fmt.Printf("\n\tTHEN\t\tthe resulting string is [(a---b---c)]\n\n")
	stringSlice := StringSlice{"a", "b", "c"}
	stringFromStringSlice := stringSlice.MkStringWithEnds("[(", "---", ")]")
	referenceStringFromStringSlice := "[(a---b---c)]"
	if stringFromStringSlice != referenceStringFromStringSlice {
		test.Errorf("\n\nTEST FAILED:\n\tgot:\t\t\t%s\n\twanted:\t%s\n\n", stringFromStringSlice, referenceStringFromStringSlice)
	}
}
