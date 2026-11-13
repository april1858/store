package main

import "errors"

var ErrorNoSuchKey = errors.New("no such key")

type KeyValueStore struct {
	m map[string]string
}

func NewKeyValueStore() *KeyValueStore {
	return &KeyValueStore{m: make(map[string]string)}
}

func (store *KeyValueStore) Put(key, value string) error {
	store.m[key] = value
	return nil
}

func (store *KeyValueStore) Get(key string) (string, error) {
	value, ok := store.m[key]
	if !ok {
		return "", ErrorNoSuchKey
	}
	return value, nil
}

func (store *KeyValueStore) Delete(key string) error {
	delete(store.m, key)
	return nil
}
