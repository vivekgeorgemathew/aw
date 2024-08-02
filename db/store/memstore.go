package store

import "errors"

type MemoryStore struct {
	riskStore map[string]interface{}
}

func (memStore *MemoryStore) Save(id string, data interface{}) error {
	memStore.riskStore[id] = data
	return nil
}

func (memStore *MemoryStore) GetAll() ([]interface{}, error) {
	if len(memStore.riskStore) < 1 {
		return nil, errors.New("empty risk store")
	}
	v := make([]interface{}, 0, len(memStore.riskStore))
	for _, value := range memStore.riskStore {
		v = append(v, value)
	}
	return v, nil
}

func (memStore *MemoryStore) Get(riskId string) (interface{}, error) {
	risk, ok := memStore.riskStore[riskId]
	if !ok {
		return nil, errors.New("risk not found")
	}
	return risk, nil
}
