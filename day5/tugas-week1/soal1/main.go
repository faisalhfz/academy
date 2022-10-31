package main

import (
	"fmt"
	"math"
)

func buyMaxItems(menu map[string]int, budget int) []string {
	items := sortFromCheapest(menu)
	total := 0
	buy := make([]string, 0)

	for _, item := range items {
		if total+menu[item] > budget {
			break
		}

		total += menu[item]
		buy = append(buy, item)
	}

	return buy
}

func sortFromCheapest(menu map[string]int) []string {
	items := make([]string, 0)

	for item, price := range menu {
		items = append(items, item)
		_ = price
	}

	for i := 0; i < len(items)-1; i++ {
		for j := 0; j < len(items)-1; j++ {
			if menu[items[j]] > menu[items[j+1]] {
				temp := items[j+1]
				items[j+1] = items[j]
				items[j] = temp
			}
		}
	}

	return items
}

func cheapestItem(menu map[string]int) string {
	min := math.MaxInt64
	var cheapestItem string
	for item, price := range menu {
		if price < min {
			min = price
			cheapestItem = item
		}
	}
	return cheapestItem
}

func highestItem(menu map[string]int) string {
	max := 0
	var highestItem string
	for item, price := range menu {
		if price > max {
			max = price
			highestItem = item
		}
	}
	return highestItem
}

func getItemsWithPrice(value int, menu map[string]int) []string {
	var items []string
	for item, price := range menu {
		if price == value {
			items = append(items, item)
		}
	}
	return items
}

func main() {
	menu := map[string]int{}

	menu["Benih Lele"] = 50000
	menu["Pakan lele cap menara"] = 25000
	menu["Probiotik A"] = 75000
	menu["Probiotik Nila B"] = 10000
	menu["Pakan Nila"] = 20000
	menu["Benih Nila A"] = 20000
	menu["Cupang"] = 5000
	menu["Benih Nila B"] = 30000
	menu["Benih Cupang"] = 10000
	menu["Probiotik B"] = 10000

	fmt.Println("-----------------------------------------------")

	// a
	keranjang := buyMaxItems(menu, 100000)

	fmt.Println("Total produk dengan harga dibawah Rp 100.000 : ", len(keranjang))

	for _, item := range keranjang {
		fmt.Println(item, "-", menu[item])
	}

	fmt.Println("-----------------------------------------------")

	// b
	fmt.Println("Daftar produk termurah: ", cheapestItem(menu), "-", menu[cheapestItem(menu)])
	fmt.Println("Daftar produk termahal: ", highestItem(menu), "-", menu[highestItem(menu)])

	fmt.Println("-----------------------------------------------")

	// c
	fmt.Println("Daftar produk dengan harga Rp 10.000 :	")

	items10k := getItemsWithPrice(10000, menu)

	for _, item := range items10k {
		fmt.Println(item, "-", menu[item])
	}

	fmt.Println("-----------------------------------------------")
}
