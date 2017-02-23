package ascii

import "testing"

/*
func GreetingASCIITest() {
	a := false
	for i := 0; i < len(Ascii); i++ {
		for j := 0; i < len(Ascii); j++ {
			if str[i] != Ascii[j] {
				fmt.Println("fatal")
			}
		}
	}
}
/*
func GreetingASCIITest(String input)	{
	bool isASCII := true
	for i := 0; i < len(container); i++	{
		c :=
	}
}
*/

func TestGreetingASCII(t *testing.T) {
	ascii := a
	if !(isASCII(ascii)) {
		t.Fail()
	}
}

func isASCII(s string) bool {
	for _, c := range s {
		if c > 127 {
			return false
		}
	}
	return true
}
