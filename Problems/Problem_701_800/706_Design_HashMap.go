package Problem_701_800

type MyHashMap struct {
	table []*Entity
	cap   int
}

type Entity struct {
	Key   int
	Value int
}

func Constructor() MyHashMap {
	table := make([]*Entity, 10240)
	for i := 0; i < 10240; i++ {
		table[i] = &Entity{-1, -1}
	}
	hashMap := MyHashMap{table, len(table)}
	return hashMap
}

func (this *MyHashMap) Put(key int, value int) {
	hashCode := key % this.cap
	for hc := hashCode; hc < this.cap; hc++ {
		if this.table[hc].Key == key {
			this.table[hc].Value = value
			return
		} else if this.table[hc].Key == -1 {
			this.table[hc].Key = key
			this.table[hc].Value = value
			return
		}
	}

	for hc := 0; hc < hashCode; hc++ {
		if this.table[hc].Key == key {
			this.table[hc].Key = value
			return
		} else if this.table[hc].Key == -1 {
			this.table[hc].Key = key
			this.table[hc].Value = value
			return
		}
	}
}

func (this *MyHashMap) Get(key int) int {
	hashCode := key % this.cap
	for hc := hashCode; hc < this.cap; hc++ {
		if this.table[hc].Key == key {
			return this.table[hc].Value
		}
	}

	for hc := 0; hc < hashCode; hc++ {
		if this.table[hc].Key == key {
			return this.table[hc].Value
		}
	}
	return -1
}

func (this *MyHashMap) Remove(key int) {
	hashCode := key % this.cap
	for hc := hashCode; hc < this.cap; hc++ {
		if this.table[hc].Key == key {
			this.table[hc].Key = -1
			return
		}
	}

	for hc := 0; hc < hashCode; hc++ {
		if this.table[hc].Key == key {
			this.table[hc].Key = -1
			return
		}
	}
}
