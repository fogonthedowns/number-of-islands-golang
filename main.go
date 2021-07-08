package main

func numIslands(grid [][]byte) int {
	nr := len(grid)
	if grid == nil || nr == 0 {
		return 0
	}

	nc := len(grid[0])
	num_islands := 0

	for r := 0; r < nr; r++ {
		for c := 0; c < nc; c++ {
			// look for, mark and score a new island
			if grid[r][c] == '1' {
				num_islands++
				grid[r][c] = '0'
			} else {
				continue
			}

			// clean up work, to mark anything around that island as found
			queue := make([][]int, 0)
			queue = append(queue, []int{r, c})

			for len(queue) > 0 {
				cur := queue[0]
				queue = queue[1:]

				row := cur[0]
				col := cur[1]

				if row-1 >= 0 {
					if grid[row-1][col] == '1' {
						queue = append(queue, []int{row - 1, col})
						grid[row-1][col] = '0'
					}
				}

				if row+1 < nr {
					if grid[row+1][col] == '1' {
						queue = append(queue, []int{row + 1, col})
						grid[row+1][col] = '0'
					}
				}

				if col-1 >= 0 {
					if grid[row][col-1] == '1' {
						queue = append(queue, []int{row, col - 1})
						grid[row][col-1] = '0'
					}
				}

				if col+1 < nc {
					if grid[row][col+1] == '1' {
						queue = append(queue, []int{row, col + 1})
						grid[row][col+1] = '0'
					}
				}

			}

		}
	}
	return num_islands
}
