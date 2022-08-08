package Problem_101_200

type Pos struct {
	R int
	C int
}

func dfs(grid [][]byte, pos Pos, visited map[Pos]bool) {
	rCnt := len(grid)
	cCnt := len(grid[0])
	if visited[pos] || pos.R < 0 || pos.R >= rCnt || pos.C < 0 || pos.C >= cCnt || grid[pos.R][pos.C] == byte(48) {
		return
	}
	visited[pos] = true
	neighbours := []Pos{{pos.R - 1, pos.C}, {pos.R, pos.C + 1}, {pos.R + 1, pos.C}, {pos.R, pos.C - 1}}
	for i := 0; i < 4; i++ {
		dfs(grid, neighbours[i], visited)
	}
}

func numIslands(grid [][]byte) int {
	rCnt := len(grid)
	cCnt := len(grid[0])
	ans := 0
	visited := make(map[Pos]bool)

	for r := 0; r < rCnt; r++ {
		for c := 0; c < cCnt; c++ {
			pos := Pos{r, c}
			if grid[r][c] == byte(49) && !visited[pos] {
				ans += 1
				dfs(grid, pos, visited)
			}
		}
	}
	return ans
}
