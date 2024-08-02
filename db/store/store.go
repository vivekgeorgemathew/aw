package store

type Store interface {
	Save(string, interface{}) error
	Get(string) (interface{}, error)
	GetAll() ([]interface{}, error)
}

func NewStore() Store {
	memStore := &MemoryStore{}
	memStore.riskStore = make(map[string]interface{})
	return memStore
}
