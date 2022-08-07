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

	crestPlane := make(map[Pos]bool)
	crestPlane[Pos{sr, sc}] = true
	visited := make(map[Pos]bool)
	visited[Pos{sr, sc}] = true

	for len(crestPlane) > 0 {
		newCrestPlane := make(map[Pos]bool)
		for pos := range crestPlane {
			neighbours := []Pos{{pos.R - 1, pos.C}, {pos.R, pos.C + 1}, {pos.R + 1, pos.C},
				{pos.R, pos.C - 1}}
			for i := 0; i < len(neighbours); i++ {
				p := neighbours[i]
				if p.R < 0 || p.R >= rCount || p.C < 0 || p.C >= cCount {
					continue
				}
				if _, e := visited[p]; e || image[p.R][p.C] != tColor {
					continue
				}

				visited[p] = true
				newCrestPlane[p] = true
			}
		}
		crestPlane = newCrestPlane
	}

	for p := range visited {
		image[p.R][p.C] = color
	}
	return image
}
