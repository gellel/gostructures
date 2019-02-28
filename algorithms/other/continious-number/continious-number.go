package continious

const (
	// R declares the constant numeric sum that is assigned when the iteration match condition does not match. 
	R int = 1
	// C declares the constant numeric sum the the sequence of repeating numbers begins at.
	C int = 1
)

// Number iterates across slice, checking adjacent indexes in the argument slice and whether positions a[i] and a[i+1] are identical.
// Counts the longest running sequence for n number of times this condition is satisfied.
func Number(a []int) (int, int) {
	i := 0
	k := -1
	r := R
	c := C
	for i < len(a)-1 {
		if a[i] == a[i+1] {
			r = r + 1
			k = a[i]
		} else {
			if r > c {
				c = r
			}
			r = R
		}
		i = i + 1
	}
	return k, c
}
