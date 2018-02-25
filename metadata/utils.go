package metadata

import "fmt"

func gogo(v Slice, sss [][]int, max int) {
	coor1 := []int{v.R1, v.C1}
	coor2 := []int{v.R2, v.C2}
	a := coor2[0] - coor1[0] + 1 //висота
	b := coor2[1] - coor1[1] + 1 //ширина
	if max > a*b {
		// if height is bigger then width than
		// try to add to the left or right sides first, so we get bigger slice
		if a > b {
			if max-a*b >= a {
				if check(left, v.ID, coor1, coor2, sss) {
					return
				}
				if check(right, v.ID, coor1, coor2, sss) {
					return
				}
			}
			if max-a*b >= b {
				if check(bottom, v.ID, coor1, coor2, sss) {
					return
				}
				if check(top, v.ID, coor1, coor2, sss) {
					return
				}
			}
		} else {
			if max-a*b >= b {
				if check(bottom, v.ID, coor1, coor2, sss) {
					return
				}
				if check(top, v.ID, coor1, coor2, sss) {
					return
				}
			}
			if max-a*b >= a {
				if check(left, v.ID, coor1, coor2, sss) {
					return
				}
				if check(right, v.ID, coor1, coor2, sss) {
					return
				}
			}
		}
	}
}

const (
	left = iota
	right
	top
	bottom
)

func check(side, ID int, coor1, coor2 []int, sss [][]int) bool {
	switch side {
	case left:
		if coor1[1]-1 >= 0 {
			fmt.Print("left ")
			//fmt.Println(sss[coor1[0]][coor1[1]-1])
			//fmt.Println(sss[coor2[0]][coor1[1]-1])
			return isSubstitued(ID, []int{coor1[0], coor1[1] - 1}, []int{coor2[0], coor1[1] - 1}, sss)
		}
		return false
	case right:
		if len(sss[0]) > coor2[1]+1 {
			fmt.Print("right ")
			//fmt.Println(sss[coor1[0]][coor2[1]+1])
			//fmt.Println(sss[coor2[0]][coor2[1]+1])
			return isSubstitued(ID, []int{coor1[0], coor2[1] + 1}, []int{coor2[0], coor2[1] + 1}, sss)
		}
		return false
	case top:
		if coor1[0]-1 >= 0 {
			fmt.Print("top ")
			//fmt.Println(sss[coor1[0]-1][coor1[1]])
			//fmt.Println(sss[coor1[0]-1][coor2[1]])
			return isSubstitued(ID, []int{coor1[0] - 1, coor1[1]}, []int{coor1[0] - 1, coor2[1]}, sss)
		}
		return false
	case bottom:
		var height = len(sss)
		if height > coor2[0]+1 {
			fmt.Print("bottom ")
			//fmt.Println(sss[coor2[0]+1][coor1[1]])
			//fmt.Println(sss[coor2[0]+1][coor2[1]])
			return isSubstitued(ID, []int{coor2[0] + 1, coor1[1]}, []int{coor2[0] + 1, coor2[1]}, sss)
		}
		return false
	}
	return false
}

func isSubstitued(ID int, a []int, b []int, sss [][]int) bool {
	f := true
	for r := a[0]; r <= b[0]; r++ {
		for c := a[1]; c <= b[1]; c++ {
			if sss[r][c] != -1 {
				f = false
			}
		}
	}
	if f {
		for r := a[0]; r <= b[0]; r++ {
			for c := a[1]; c <= b[1]; c++ {
				sss[r][c] = ID
			}
		}
	}
	return f
}
