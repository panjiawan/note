package main

import (
	"fmt"
	"sync"
)

type Inventory struct {
	stock   int          // 库存数量
	rwMutex sync.RWMutex // 保护库存
}

// 查询库存（使用读写锁的读锁）
func (inv *Inventory) getStock() int {
	inv.rwMutex.RLock()
	defer inv.rwMutex.RUnlock()
	return inv.stock
}

// 扣减库存（使用互斥锁和读写锁的写锁）
func (inv *Inventory) deductStock(userID int, quantity int) bool {
	// 先检查库存（写锁）
	inv.rwMutex.Lock()
	defer inv.rwMutex.Unlock()

	// 检查，防止超卖
	if inv.stock < quantity {
		fmt.Printf("用户%d: 库存不足，扣减失败\n", userID)
		return false
	}

	inv.stock -= quantity
	fmt.Printf("用户%d: 成功购买%d件，剩余库存%d\n", userID, quantity, inv.stock)
	return true
}

func main() {
	inventory := &Inventory{stock: 10} // 初始库存10件

	var wg sync.WaitGroup

	// 模拟100个用户同时抢购
	for i := 1; i <= 100; i++ {
		wg.Add(1)
		go func(userID int) {
			defer wg.Done()
			inventory.deductStock(userID, 1)
		}(i)
	}

	wg.Wait()
	fmt.Printf("最终库存: %d\n", inventory.getStock())
}
