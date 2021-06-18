package iota

type SubStore struct {
	Subs []*Subscriber
}

func NewSubStore() SubStore {
	var subs []*Subscriber
	return SubStore{subs }
}

func (store *SubStore) AddSub(subscriber *Subscriber) {
	store.Subs = append(store.Subs, subscriber)
}

func (store *SubStore) DropSubs() {
	for _, sub := range store.Subs {
		sub.Drop()
	}
}
