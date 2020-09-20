package Week_04

//dfs解法
//二维平面的每个点都会访问和处理一次，时间复杂度O(m*n)
//遍历时，最大深度可能全部节点在一次处理中访问，此时栈空间为m*n, 空间复杂度O(m*n)
func numIslands(grid [][]byte) int {
    count := 0
    for i := 0; i < len(grid); i++ {
        for j := 0; j < len(grid[0]); j++ {
            if grid[i][j] == '1' {
                count++
                visitOneLand(grid, i, j)
            }
        }
    }
    return count
}

func visitOneLand(grid [][]byte, x, y int) {
    if x < 0 || x >= len(grid) || y < 0 || y >= len(grid[0]) || grid[x][y] != '1' {
        return
    }
    grid[x][y] = '2'
    visitOneLand(grid, x+1, y)
    visitOneLand(grid, x-1, y)
    visitOneLand(grid, x, y+1)
    visitOneLand(grid, x, y-1)
}

//bfs解法
//二维平面的每个点都会访问和处理一次，时间复杂度O(m*n)
//遍历时，最大宽度可能平面的边缘节点个数，此时需要缓存m+n个节点, 空间复杂度O(m+n)
func numIslands1(grid [][]byte) int {
    count := 0
    for i := 0; i < len(grid); i++ {
        for j := 0; j < len(grid[0]); j++ {
            if grid[i][j] == '1' {
                count++
                visitOneLand(grid, i, j)
            }
        }
    }
    return count
}

type point struct {
    x int
    y int
}

func visitOneLand(grid [][]byte, i, j int) {
    cache := []*point{&point{i, j}}
    grid[i][j] = '2'
    visitOnePoint := func(x, y int) {
        if x < 0 || x >= len(grid) || y < 0 || y >= len(grid[0]) || grid[x][y] != '1' {
            return
        }
        cache = append(cache, &point{x, y})
        grid[x][y] = '2'
    }

    for len(cache) > 0 {
        pt := cache[0]
        cache = cache[1:]
        visitOnePoint(pt.x+1, pt.y)
        visitOnePoint(pt.x-1, pt.y)
        visitOnePoint(pt.x, pt.y+1)
        visitOnePoint(pt.x, pt.y-1)
    }
}
