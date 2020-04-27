import (
"fmt"
)

type cord struct {
    x int
    y int
}

var sum map[int]int 

type degree map[int]int 

func (d degree) incr(key int) {
    if _, ok := d[key]; ok {
        d[key] += 1
    } else {
        d[key] = 1
    }
}
func min (a int, b int) int {
    if a > b {
        return b
    } else {
        return a
    }
}

/*
[[2,1,1],[2,3,1],[3,4,1],[3,6,1]]
g:= map[2:[1 3] 3:[4]]
w:= map[{2 1}:1 {2 3}:1 {3 4}:1]
*/

func bfs(g map[int][]int, w map[cord]int, k int, dd degree){
    
    level:= []int{k}
    trav := map[int]int{k:0}
    for len(level) >0 {
        var next_level []int
        for _, node := range(level){
            for _, child := range(g[node]) {
                val := sum[node] + w[cord{node,child}]
                if _, ok:= sum[child]; !ok{
                    sum[child] = val
                } else {
                    sum[child] = min(sum[child], val)
                }
                if dd[child] == trav[child] {continue}
                trav[child] +=1
                next_level = append(next_level, child)
            }
        }
        level = next_level
    }
    return
}

func networkDelayTime(times [][]int, N int, K int) (mx int) {
    sum = make(map[int]int)
    g:= make(map[int][]int) //graph
    w:= make(map[cord]int)  //weight 
    dd := degree{} //degree
    for i:=0; i<len(times) ;i++ {
        g[times[i][0]] = append(g[times[i][0]], times[i][1])
        var d cord 
        d.x = times[i][0]
        d.y = times[i][1]
        w[d] = times[i][2]
        dd.incr(d.y)
    }
    sum[K] = 0
    bfs(g,w,K,dd)
    if len(sum) != N {
        return -1
    } 
    ans := sum[K]
    for n, _ := range(sum){
        if sum[n] > ans {
            ans = sum[n]
        }
    }
    fmt.Println(g,w,sum, dd)
    return ans
}
