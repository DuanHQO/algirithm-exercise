package Ch_2_5_Applications

import (
	"fmt"
	"math/rand"
)

type Stock struct {
	Price int
	Count int
}

func StockGenerator() Stock {
	t := Stock{
		Price: rand.Intn(3000) + 1000,
		Count: rand.Intn(990) + 10,
	}
	return t
}

func SortPrice(a, b Stock) int {
	if a.Price < b.Price {
		return -1
	} else if a.Price > b.Price {
		return 1
	} else {
		return 0
	}
}

func SortCount(a, b Stock) int {
	if a.Count < b.Count {
		return -1
	} else if a.Count > b.Count {
		return 1
	} else {
		return 0
	}
}

type Customer struct {
	Ks   []Stock
	Size int
}

func (this *Customer) ReSize(newSize int) {
	ks := make([]Stock, newSize+1)
	for i := 1; i <= this.Size; i++ {
		ks[i] = this.Ks[i]
	}
	this.Ks = ks
}

func (this *Customer) IsEmpty() bool {
	return this.Size == 0
}

func (this *Customer) AddStock(t Stock) {
	if this.Size == len(this.Ks)-1 {
		this.ReSize(this.Size << 1)
	}
	this.Size++
	this.Ks[this.Size] = t
	this.swim(this.Size)
}

func (this *Customer) ShowStocks() {
	fmt.Println("------------------------------------")
	for i := 0; i < len(this.Ks); i++ {
		if this.Ks[i].Count > 0 {
			fmt.Printf("序号: %d %v\n", i, this.Ks[i])
		}
	}
	fmt.Println("------------------------------------")
}

func (this *Customer) Compare(a, b Stock) int {
	if SortPrice(a, b) < 0 {
		return -1
	} else if SortPrice(a, b) > 0 {
		return 1
	} else if SortCount(a, b) > 0 {
		return -1
	} else if SortCount(a, b) < 0 {
		return 1
	}
	return 0
}

func (this *Customer) swim(k int) {
	t := this.Ks[k]
	for k > 1 && this.Compare(t, this.Ks[k>>1]) < 0 {
		this.Ks[k] = this.Ks[k>>1]
		k >>= 1
	}
	this.Ks[k] = t
}

func (this *Customer) sink(k int) {
	t := this.Ks[k]
	for j := k << 1; j <= this.Size; j = k << 1 {
		if this.Compare(this.Ks[j], this.Ks[j+1]) > 0 {
			j++
		}
		if this.Compare(t, this.Ks[j]) <= 0 {
			break
		}
		this.Ks[k] = this.Ks[j]
		k = j
	}
	this.Ks[k] = t
}

func (this *Customer) OfHand() Stock {
	if this.IsEmpty() {
		panic("priority queue underflow")
	}
	minSellPrice := this.Ks[1]
	this.Ks[1] = this.Ks[this.Size]
	this.Size--
	this.sink(1)
	this.Ks[this.Size+1] = Stock{}
	if this.Size > 0 && this.Size == (len(this.Ks)-1)>>2 {
		this.ReSize((len(this.Ks) - 1) >> 1)
	}
	return minSellPrice
}

type Seller struct {
	Customer
}

//买家
type Buyer struct {
	Customer
}

func (this *Buyer) Compare(a, b Stock) int {
	result1 := SortPrice(a, b)
	if result1 == 0 {
		result2 := SortCount(a, b)
		if result2 == 0 {
			return 0
		}
		return SortCount(a, b)
	}
	return result1
}

func TestTrade() {
	seller := Seller{Customer{
		Ks:   make([]Stock, 2),
		Size: 0,
	}}
	buyer := Buyer{Customer{
		Ks:   make([]Stock, 2),
		Size: 0,
	}}

	for i := 0; i < 100; i++ {
		seller.AddStock(StockGenerator())
	}
	//seller.ShowStocks()
	for i := 0; i < 100; i++ {
		buyer.AddStock(StockGenerator())
	}
	//buyer.ShowStocks()
	fmt.Println(seller.IsEmpty(), buyer.IsEmpty())
	for !seller.IsEmpty() && !buyer.IsEmpty() {
		wantToBuy := buyer.OfHand()
		wantToSell := seller.OfHand()
		if wantToBuy.Price < wantToSell.Price {
			seller.AddStock(wantToSell)
		} else {
			count := wantToBuy.Count
			if wantToBuy.Count > wantToSell.Count {
				count = wantToSell.Count
			}
			fmt.Printf("交易成功 ! 买家买了价格为 %d 的股票 %d 股\n", wantToBuy.Price, count)
			if wantToSell.Count > 0 {
				seller.AddStock(wantToSell)
			}
		}
	}
	if seller.IsEmpty() {
		fmt.Println("股票都卖光了！")
	}
	if buyer.IsEmpty() {
		fmt.Println("想买的股票都买到了！")
	}
}
