package main

type Store interface {
	Save(*Workday) error
	LoadAll() ([]*Workday, error)
}

type MemStore struct {
	data []*Workday
}

func (m *MemStore) Save(w *Workday) error {
	m.data = append(m.data, w)
	return nil
}

func (m *MemStore) LoadAll() ([]*Workday, error) {
	return m.data, nil
}
