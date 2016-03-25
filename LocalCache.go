package main

type LocalCache struct {
   mapCache map[string][]byte
}


func NewLocalCache() Cache {
    c := new(LocalCache)
    c.mapCache =  make( map[string][]byte)
    return c
}

func (local LocalCache) Get(key string)  []byte {
   return local.mapCache[key];
}


func (local LocalCache) Put(key string, content []byte)  {
   local.mapCache[key] = content; 
}