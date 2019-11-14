package _6_排序算法

//简单选择排序
//特点:O(n^2)不稳定
func SimplySelectSort(arr []int) {
	for i := 0;i < len(arr); i ++ {
		x := i //默认第一个元素就是最小值
		for j := i+1; j < len(arr); j ++ {
			if arr[j] < arr[x] { //如果有比x下标还小的则重新赋值
				x = j
			}
		}
		if x != i {//如果判断x,i不相等则交换位置.
			arr[i], arr[x] = arr[x], arr[i]
		}
	}
}
