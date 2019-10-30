package UnionFind

type WeightedQuickUnionUF struct {
	id []int
	sz []int
	count int
}

func (this *WeightedQuickUnionUF) New(n int)  {
	this.id = make([]int, n)
	this.sz = make([]int, n)
	for key := range this.sz {
		this.sz[key] = 1
	}
	for key := range this.id {
		this.id[key] = key
	}
	this.count = n
}

func (this *WeightedQuickUnionUF) Count() int {
	return this.count
}

func (this *WeightedQuickUnionUF) Find(p int) int {
	for  {
		if this.id[p] != p {
			p = this.id[p]
		} else {
			break
		}
	}
	return p
}

func (this *WeightedQuickUnionUF) Connected(p, q int) bool {
	return this.Find(p) == this.Find(q)
}

func (this *WeightedQuickUnionUF) Union(p, q int) {
	i := this.Find(p)
	j := this.Find(q)
	if i == j {
		return
	}

	if this.sz[i] < this.sz[j] {
		this.id[i] = this.id[j]
		this.sz[j] += this.sz[i]
	} else {
		this.id[j] = this.id[i]
		this.sz[i] += this.sz[j]
	}
	this.count--

}
