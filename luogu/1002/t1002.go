package main

import "fmt"

func main() {
	var bx, by, mx, my int

	fx := []int{0, -2, -1, 1, 2, 2, 1, -1, -2}
	fy := []int{0, 1, 2, 2, 1, -1, -2, -2, -1}
	fmt.Scan(&bx, &by, &mx, &my)
	bx += 2
	by += 2
	mx += 2
	my += 2

	var f [40][40]int
	f[2][1] = 1

	var s [40][40]bool
	s[mx][my] = true
	for i := 1; i <= 8; i++ {
		s[mx+fx[i]][my+fy[i]] = true
	}

	for i := 2; i <= bx; i++ {
		for j := 2; j <= by; j++ {
			if s[i][j] {
				continue
			}
			f[i][j] = f[i-1][j] + f[i][j-1]
		}
	}

	fmt.Println(f[bx][by])
}
