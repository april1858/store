package main

import "errors"

var store = make(map[string]string)

var keyNotFound = errors.New("key not found")

func Put(key, value string) error {
	store[key] = value
	return nil
}

func Get(key string) (string, error) {
	v, ok := store[key]
	if !ok {
		return "", keyNotFound
	}
	return v, nil
}

func Delete(key string) error {
	delete(store, key)
	return nil
}
