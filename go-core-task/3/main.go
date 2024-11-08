package main

import "fmt"

type StringIntMap struct {
	data map[string]int
}

func NewStringIntMap() *StringIntMap {
	return &StringIntMap{data: make(map[string]int)}
}

func (m *StringIntMap) Add(key string, value int) {
	m.data[key] = value
}

func (m *StringIntMap) Remove(key string) {
	delete(m.data, key)
}

func (m *StringIntMap) Copy() map[string]int {
	copyMap := make(map[string]int)
	for key, value := range m.data {
		copyMap[key] = value
	}
	return copyMap
}

func (m *StringIntMap) Exists(key string) bool {
	_, exists := m.data[key]
	return exists
}

func (m *StringIntMap) Get(key string) (int, bool) {
	value, ok := m.data[key]
	return value, ok
}

func main() {
	str := NewStringIntMap()
	str.Add("one", 1)
	str.Add("two", 2)
	str.Add("three", 3)
	fmt.Println("мапа после добавления в неё значений:", str.data)
	str.Remove("one")
	fmt.Println("мапа после удаления ключа 'one':", str.data)
	fmt.Println("копия мапы:", str.Copy())
	fmt.Println("проверка наличия ключа 'two'", str.Exists("two"))
	if value, found := str.Get("two"); found {
		fmt.Printf("значение по ключу 'two' получено: %d\n", value)
	} else {
		fmt.Println("ключ 'two' не найден")
	}
}
