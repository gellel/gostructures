package continious

const (
	R int = 1
	C int = 1
)

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
			r = 0
		}
		i = i + 1
	}
	return k, c
}
