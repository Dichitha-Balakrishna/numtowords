package numtowords

//MaxNum is the largest number that can be converted to words by this package
import (
	"fmt"
)

const MaxNum = 99

// converts the given number to words
func Convert(num int) (string, error) {
	arr := [20]string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten", "eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen", "seventeen", "eighteen", "ninteen"}
	arr2 := [8]string{"twenty", "thirty", "fourty", "fifty", "sixty", "seventy", "eighty", "ninety"}
	if num < 0 || num > MaxNum {
		return "", fmt.Errorf("number out of range")
	} else if num <= 19 {
		return arr[num], nil
	} else {
		tens := num / 10
		ones := num % 10
		if ones == 0 {
			return arr2[tens-2], nil
		} else {
			return arr2[tens-2] + " " + arr[ones], nil
		}
	}

}
