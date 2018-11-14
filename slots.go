package slots

// Slots is used to select a slot (index) in a list.
type Slots chan int

// New creates a new Slots object, able to return "n" items (0 to n-1)
func New(n int) Slots {
	s := make(chan int, n)
	for i := 0; i < n; i++ {
		s <- i
	}

	return s
}

// Take returns the first empty slot in the list.
// If there are no empty slots, it will wait for one to become available
func (s Slots) Take() int {
	return <-s
}

// Release return a "used" slot to the list and makes it available again.
func (s Slots) Release(v int) {
	s <- v
}
