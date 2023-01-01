package cache

type Cache map[string]interface{}

func New() Cache {
	return make(map[string]interface{})
}

func (c Cache) Set(key string, val interface{}) {
	c[key] = val
}

func (c Cache) Get(key string) interface{} {
	return c[key]
}

func (c Cache) Delete(key string) {
	delete(c, key)
}

// func (c Cache) Show() {
// 	for k, v := range c {
// 		fmt.Println("key: ", k, "value: ", v)
// 	}
// }
