package main

import (
	"fmt"
)

func minCostToFinishWalk(N int, K int, supplyStations []int) int {
	totalCost := 0
	currentFood := K // 小明开始时携带K份食物

	for i := 0; i < N; i++ {
		currentFood-- // 消耗1份食物

		if currentFood < 1 && i < N-1 { // 如果食物不足且不是最后一天，需要购买
			buyFood := K - currentFood               // 计算需要购买的食物数量
			totalCost += buyFood * supplyStations[i] // 购买食物并累加花费
			currentFood = K                          // 更新当前携带的食物数量
		}
	}

	return totalCost
}

func main() {
	N := 5                                 // 总路程需要5天
	K := 3                                 // 最多能携带3份食物
	supplyStations := []int{1, 2, 3, 3, 2} // 每个补给站的食物价格，第0天没有补给站，所以价格为0

	fmt.Println(minCostToFinishWalk(N, K, supplyStations)) // 输出最小花费
}
