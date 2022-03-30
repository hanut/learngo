package main

import (
	"fmt"
	"time"
)

func main() {
	st := time.Now()
	var h, m, dh, dm, fh, fm uint8
	fmt.Scanf("%d %d\n%d %d\n", &h, &m, &dh, &dm)

	fh = h + dh
	if fh >= 24 {
		fh = fh % 24
	}
	fm = m + dm
	if fm >= 60 {
		fh += fm / 60
		fm = fm % 60
	}
	r := fmt.Sprintf("%02d %02d", fh, fm)
	fmt.Println(r)
	ts := time.Since(st)
	fmt.Println("Exec time", ts)
}
