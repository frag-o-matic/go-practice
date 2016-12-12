package main

import (
    "fmt"
)

// Clock type API:
//
// New(hour, minute int) Clock     // a "constructor"
// (Clock) String() string         // a "stringer"
// (Clock) Add(minutes int) Clock

type Clock int

// constructor
func New(hour, minute int) Clock {
    time:=(hour*60 + minute) % (60 * 24)
	if time < 0 {
		time += 60 * 24
	}
	return Clock(time)
}

// stringer
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c/60, c%60)
}

// Add minutes to clock
func (c Clock) Add(minutes int) Clock {
	return New(0, int(c)+minutes)
}

func main() {
    c1 := New(8,0)
    c2 := New(8,0)
    c3 := New(12,15)
    c4 := New(22,30)
    c4 = c4.Add(5)
    fmt.Println(c1)
    fmt.Println(c1==c2)
    fmt.Println(c2==c3)
    fmt.Println(c4)
}
