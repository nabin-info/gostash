package main

import "fmt"

var timeZone = map[string]int{
	"UTC":  0*60*60,
	"EST": -5*60*60,
	"CST": -6*60*60,
	"MST": -7*60*60,
	"PST": -8*60*60,
}

func offset(tz string) int {
	if seconds, ok := timeZone[tz]; ok {
		return seconds
	}
	fmt.Println("unknown time zone:", tz)
	return 0
}

func main() {
	fmt.Printf("https://golang.org/doc/effective_go.html\n")

	// short declarations example (i is local to for{})
	sum := 0
	for i := 0; i < 10; i++ {
	    sum += i
	}

	// range loop is great: []array [:]slice `"'string'"` [key]map <-channel
	for k, v:= range timeZone {
		fmt.Printf("timeZone[%s] = %d\n", k, v); 
	}

	fmt.Println("California timezone has an offset of:", offset("PST"))
	delete(timeZone, "PDT")  // builtin deletes a key from a map
	fmt.Println("Daylight Savings sucks ...", offset("PDT"))
}
