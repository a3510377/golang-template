package log

import (
	"fmt"
	"math"
)

func Test() {
	const MAX = 6
	Format := fmt.Sprintf("%%d: %%#0%db %%d\n", MAX)

	for i := 0; i < MAX; i++ {
		d := (int)(math.Pow(2, float64(i)))
		fmt.Printf(Format, i, d, d)
	}
}
