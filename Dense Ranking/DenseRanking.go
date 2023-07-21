package main

import (
	"fmt"
)

func getRank(skor []int, myPoint int) int {
	urutanSkor := make([]int, 0)
	for _, score := range skor {
		if len(urutanSkor) == 0 || urutanSkor[len(urutanSkor)-1] != score {
			urutanSkor = append(urutanSkor, score)
		}
	}

	rank := 1
	for _, score := range urutanSkor {
		if myPoint >= score {
			return rank
		}
		rank++
	}

	return rank
}

func main() {
	var n, m int
	fmt.Scan(&n)
	scores := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&scores[i])
	}
	fmt.Scan(&m)
	myScore := make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Scan(&myScore[i])
	}

	// Cari rangking untuk setiap skor yang ingin GITS ketahui
	for i := 0; i < m; i++ {
		ranking := getRank(scores, myScore[i])
		fmt.Printf("%d ", ranking)
	}
	fmt.Println()
}
