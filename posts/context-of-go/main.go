package main

import (
	"context"
	"fmt"
	"time"
)

const DB_ADDRESS = "db_address"
const CALCULATE_VALUE = "calculate_value"

func readDB(ctx context.Context, cost time.Duration) {
	fmt.Println("DB address is ", ctx.Value(DB_ADDRESS))

	select {
	case <-time.After(cost):
		fmt.Println("read data from db")
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
}

func calculate(ctx context.Context, cost time.Duration) {
	fmt.Println("calculate value is", ctx.Value(CALCULATE_VALUE))
	select {
	case <-time.After(cost): //  模拟数据计算
		fmt.Println("calculate finish")
	case <-ctx.Done():
		fmt.Println(ctx.Err()) // 任务取消的原因
		// 一些清理工作
	}
}

func main() {
	ctx := context.Background()

	// Add Context info
	ctx = context.WithValue(ctx, DB_ADDRESS, "localhost:3306")
	ctx = context.WithValue(ctx, CALCULATE_VALUE, "123")

	ctx, cancel := context.WithTimeout(ctx, time.Second*2)
	defer cancel()

	go readDB(ctx, time.Second*4)
	go calculate(ctx, time.Second*4)

	time.Sleep(time.Second * 5)
}
