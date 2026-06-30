package numtowords_test

import (
	"testing"

	numtowords "numtowords_test"
)

func TestInvalidNumber(t *testing.T) {
	_, err := numtowords.Convert(numtowords.MaxNum + 1)
	if err != nil {
		t.Log("expected error")
		t.Fail()

	}
	_, err = numtowords.Convert(-1)
	if err == nil {
		t.Log("expected error")
		t.Fail()
	}

}

// func TestHundreds(t *testing.T) {
// 	testcases := map[int]string{
// 		109: "One hundred and nine",
// 		333: "Three hundred and thirty three",

// 	}
// }
