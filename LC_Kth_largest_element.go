import (
    "container/heap"
    "fmt"
)

type List []int 

func (L List) Len () int { return len(L)}

func (L List) Less(i,j int) bool {
    return L[i] > L[j]
}

func (h List) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *List) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *List) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func findKthLargest(nums []int, k int) int {
    h := make(List, len(nums))
    for i:=0; i<len(nums);i++ {
        //heap.Push(&h, nums[i])
        h[i] = nums[i]
    }
    
    heap.Init(&h)
    fmt.Println(h)
    for (k>1){
        heap.Pop(&h)
        k-=1
    }
    x:= heap.Pop(&h).(int)
    return x
}
