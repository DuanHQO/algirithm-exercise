package Ch_2_5_Applications

import "math/rand"

type Stock struct {
	Counter int
	id int
	Price float64
	Count int
}

type Customer interface {
	ReSize(newSize int)

}

func StockGenerator() Stock {
	t := Stock{
		Counter: rand.Intn(3000) + 1000,
		id:      rand.Intn(990) + 10,
		Price:   0,
		Count:   0,
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

type Seller struct {
	Ks []Stock
	Size int
}

func (this *Seller) ReSize(newSize int)  {
	ks := make([]Stock, newSize + 1)
	for i := 1; i <= this.Size; i++ {
		ks[i] = this.Ks[i]
	}
	this.Ks = ks
}

func (this *Seller) IsEmpty() bool {
	return this.Size == 0
}

func (this *Seller) AddStock(t Stock) {
	if this.Size == len(this.Ks) - 1 {
		this.ReSize(this.Size << 1)
	}
	this.Size++
	this.Ks[this.Size] = t
}

func (this *Seller) Compare(a, b Stock) int {
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

func (this *Seller) OfHand() Stock {
	if this.IsEmpty() {
		panic("priority queue underflow")
	}
	minSellPrice := this.Ks[1]
	this.Ks[1] = this.Ks[this.Size]
	this.Size--
	this.Ks[this.Size + 1] = Stock{}
	if this.Size > 0 && this.Size == (len(this.Ks) - 1) >> 2 {
		this.ReSize((len(this.Ks) - 1) >> 1)
	}
	return minSellPrice
}

func (this *Seller) swim(k int) {
	t := this.Ks[k]
	for k > 1 && this.Compare(t, this.Ks[k >> 1]) < 0 {
		this.Ks[k] = this.Ks[k >> 1]
		k >>= 1
	}
	this.Ks[k] = t
}

func (this *Seller) sink(k int) {
	t := this.Ks[k]
	for j := k << 1; j <= this.Size; {
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


//买家
type Buyer struct {
	Ks []Stock
	Size int
}

func (this *Buyer) ReSize(newSize int)  {
	ks := make([]Stock, newSize + 1)
	for i := 1; i <= this.Size; i++ {
		ks[i] = this.Ks[i]
	}
	this.Ks = ks
}

func (this *Buyer) IsEmpty() bool {
	return this.Size == 0
}

func (this *Buyer) AddStock(t Stock) {
	if this.Size == len(this.Ks) - 1 {
		this.ReSize(this.Size << 1)
	}
	this.Size++
	this.Ks[this.Size] = t
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

func (this *Buyer) OfHand() Stock {
	if this.IsEmpty() {
		panic("priority queue underflow")
	}
	minSellPrice := this.Ks[1]
	this.Ks[1] = this.Ks[this.Size]
	this.Size--
	this.Ks[this.Size + 1] = Stock{}
	if this.Size > 0 && this.Size == (len(this.Ks) - 1) >> 2 {
		this.ReSize((len(this.Ks) - 1) >> 1)
	}
	return minSellPrice
}

func (this *Buyer) swim(k int) {
	t := this.Ks[k]
	for k > 1 && this.Compare(t, this.Ks[k >> 1]) < 0 {
		this.Ks[k] = this.Ks[k >> 1]
		k >>= 1
	}
	this.Ks[k] = t
}

func (this *Buyer) sink(k int) {
	t := this.Ks[k]
	for j := k << 1; j <= this.Size; {
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


