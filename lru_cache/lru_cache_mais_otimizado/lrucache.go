package main

import "fmt"

type Item struct {
	key   string
	value any
	next  *Item
	prev  *Item
}

type Cache struct {
	capacity int
	dict     map[string]*Item
	items    *Item
	head     *Item
}

//função construtora
func New(capacity int) *Cache {

	return &Cache{
		capacity: capacity,
		dict:     make(map[string]*Item),
	}
}

func (c *Cache) Get(key string) {
}

func main() {
	// Usando a função New para criar uma instância de Cache
	cache := New(10)

	// Verificando o tipo de retorno e capacidade do cache
	fmt.Printf("Tipo de retorno: %T\n", cache)
	fmt.Printf("Capacidade do cache: %d\n", cache.capacity)
	fmt.Printf("Dicionário do cache: %d\n", cache.dict)

}
