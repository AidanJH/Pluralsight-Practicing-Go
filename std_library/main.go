//Strings
package main

import (
	"fmt"
	"log"
	"strings"
	"time"
)

func main() {
	//byteCode()
	//compareStrings()
	timedExecution()
}

func byteCode() {
	theString := "\x47\x6f\x20"
	otherString := "This is a string"

	for i := 0; i < len(theString); i++ {
		//Print out actual hexcodes
		fmt.Printf("%x ", theString[i])
	}
	//Print out hexcode as actual string
	fmt.Println("\n" + theString)

	//Print out individual characters from hexcode byte sequence using %q
	for i := 0; i < len(theString); i++ {
		//Print out actual hexcodes
		fmt.Printf("%q ", theString[i])
	}

	//Print out a range of characters from a string
	fmt.Println("\n" + otherString[0:2])

	//Print out a solo character as ascii
	fmt.Println(otherString[1])
	//Print out a solo character
	fmt.Println(otherString[1:2])
}

func compareStrings() {
	s1 := "hello"
	s2 := "hello"

	s3 := "Hello"

	//This is a case sensitive comparison
	if s1 == s2 {
		fmt.Println("Strings match")
	}

	//if 0 is returned then the strings match
	//Strings compare is slightly faster in operation
	if strings.Compare(s1, s3) == 0 {
		fmt.Println("Strings also match")
	}

}

func timedExecution() {
	BigString1 := "3333"
	BigString3 := "3334"

	start1 := time.Now()
	if BigString1 == BigString3 {
		fmt.Println("Big strings are equal")
	} else {
		fmt.Println("Big strings are not equal")
	}
	timeTrack(start1, "basic compare")

	start2 := time.Now()
	if strings.Compare(BigString1, BigString3) == 0 {
		fmt.Println("Big strings are equal")
	} else {
		fmt.Println("Big strings are not equal")
	}
	timeTrack(start2, "Strings.Compare")

}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}

func compareCaseInsensitiveStrings(a, b string) bool {
	if len(a) == len(b) {
		if strings.Compare(strings.ToLower(a), strings.ToLower(b)) == 0 {
			return true
		}
	}
	return false
}
