package main

import "fmt"

func main() {
	prices := map[string]int{}

	prices["Benih Lele"] = 50000
	prices["Pakan lele cap menara"] = 25000
	prices["Probiotik A"] = 75000
	prices["Probiotik Nila B"] = 10000
	prices["Pakan Nila"] = 20000
	prices["Benih Nila A"] = 20000
	prices["Cupang"] = 5000
	prices["Benih Nila B"] = 30000
	prices["Benih Cupang"] = 10000
	prices["Probiotik B"] = 10000

	fmt.Println("-----------------------------------------------")

	// a
	fmt.Println("Total produk dengan harga dibawah  Rp 100.000 :")

	items := make([]string, len(prices))

	i := 0
	for item, price := range prices {
		items[i] = item
		i++
		_ = price
	}

	for i := 0; i < len(items)-1; i++ {
		for j := 0; j < len(items)-1; j++ {
			if prices[items[j]] > prices[items[j+1]] {
				temp := items[j+1]
				items[j+1] = items[j]
				items[j] = temp
			}
		}
	}

	total := 0
	for _, item := range items {
		total += prices[item]
		if total > 100000 {
			break
		}
		fmt.Println(item, "-", prices[item])
	}

	fmt.Println("-----------------------------------------------")

	// b
	fmt.Println("Daftar produk termurah: ", items[0], "-", prices[items[0]])
	fmt.Println("Daftar produk termahal: ", items[len(items)-1], "-", prices[items[len(items)-1]])

	fmt.Println("-----------------------------------------------")

	// c
	fmt.Println("Daftar produk dengan harga Rp 10.000 :	")
	for item, price := range prices {
		if price == 10000 {
			fmt.Println(item, "-", price)
		}
	}

	fmt.Println("-----------------------------------------------")
}
