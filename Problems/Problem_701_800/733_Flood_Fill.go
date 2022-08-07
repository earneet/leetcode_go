package Problem_701_800

type Pos struct {
	R int
	C int
}

func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	if color == 0 {
		return image
	}

	rCount := len(image)
	cCount := len(image[0])
	tColor := image[sr][sc]

	crestal := make(map[Pos]bool)
	crestal[Pos{sr, sc}] = true
	visited := make(map[Pos]bool)
	visited[Pos{sr, sc}] = true

	for len(crestal) > 0 {
		newCrestal := make(map[Pos]bool)
		for pos, _ := range crestal {
			neibours := []Pos{Pos{pos.R - 1, pos.C}, Pos{pos.R, pos.C + 1}, Pos{pos.R + 1, pos.C}, Pos{pos.R, pos.C - 1}}
			for i := 0; i < len(neibours); i++ {
				p := neibours[i]
				if p.R < 0 || p.R >= rCount || p.C < 0 || p.C >= cCount {
					continue
				}
				if _, e := visited[p]; e || image[p.R][p.C] != tColor {
					continue
				}

				visited[p] = true
				newCrestal[p] = true
			}
		}
		crestal = newCrestal
	}

	for p, _ := range visited {
		image[p.R][p.C] = color
	}
	return image
}
