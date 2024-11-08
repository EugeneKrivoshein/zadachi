package main

import (
	"reflect"
	"testing"
)

func TestAdd(t *testing.T) {
	m := NewStringIntMap()
	m.Add("five", 5)

	if value, exists := m.Get("five"); !exists || value != 5 {
		t.Errorf("ожидаем %d, получили %d", 5, value)
	}
}

func TestRemove(t *testing.T) {
	m := NewStringIntMap()
	m.Add("five", 5)
	m.Remove("five")

	if _, exists := m.Get("five"); exists {
		t.Errorf("после удаления 'five' все еще в мапе")
	}
}

func TestCopy(t *testing.T) {
	m := NewStringIntMap()
	m.Add("five", 5)
	m.Add("six", 6)
	copyMap := m.Copy()

	// проверка, что копия содержит те же значения
	if !reflect.DeepEqual(copyMap, m.data) {
		t.Errorf("ожидаем %v, получили %v", m.data, copyMap)
	}

	// изменение копии не должно менять ориг карту
	copyMap["five"] = 7
	if value, _ := m.Get("five"); value != 5 {
		t.Errorf("Ошибка в Copy(): изменение копии не должно влиять на оригинальную карту")
	}
}

func TestExists(t *testing.T) {
	m := NewStringIntMap()
	m.Add("five", 5)

	if !m.Exists("five") {
		t.Errorf("ожидалось true для ключа 'five'")
	}

	if m.Exists("six") {
		t.Errorf("ожидалось значение false для несуществующего ключа 'six'")
	}
}

func TestGet(t *testing.T) {
	m := NewStringIntMap()
	m.Add("five", 5)

	value, exists := m.Get("five")
	if !exists || value != 5 {
		t.Errorf("ожидаем %d, получили %d", 5, value)
	}

	if _, exists := m.Get("six"); exists {
		t.Errorf("ожидалось значение false для несуществующего ключа 'six'")
	}
}
