package main

// Abs returns the unsigned value of any input
func Abs(in int) uint {
	if in < 0 {
		return uint(in * -1)
	}
	return uint(in)
}
