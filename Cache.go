package main
type Cache interface {
Get(key string)  []byte
Put(key string, content  []byte)
}