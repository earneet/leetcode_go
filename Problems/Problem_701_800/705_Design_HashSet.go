package Problem_701_800

type Entry struct {
	Key  int
	Next *Entry
}

type MyHashSet struct {
	entries []*Entry
}

func Constructor_() MyHashSet {
	return MyHashSet{make([]*Entry, 1024, 1024)}
}

func (this *MyHashSet) Add(key int) {
	hash := key % 1024
	root := this.entries[hash]
	if root == nil {
		this.entries[hash] = &Entry{key, nil}
	} else {
		pre := root
		for root != nil {
			if root.Key == key {
				return
			}
			pre = root
			root = root.Next
		}
		pre.Next = &Entry{key, nil}
	}
}

func (this *MyHashSet) Remove(key int) {
	hash := key % 1024
	root := this.entries[hash]
	if root == nil {
		return
	}
	if root.Key == key {
		this.entries[hash] = root.Next
		return
	}
	pre := root
	root = root.Next
	for root != nil {
		if root.Key == key {
			pre.Next = root.Next
			return
		}
		pre = root
		root = root.Next
	}
}

func (this *MyHashSet) Contains(key int) bool {
	hash := key % 1024
	root := this.entries[hash]
	if root == nil {
		return false
	}
	for ; root != nil; root = root.Next {
		if root.Key == key {
			return true
		}
	}
	return false
}
