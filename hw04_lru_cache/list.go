package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	len   int
	first *ListItem
	last  *ListItem
}

// NewList Создать новый список
func NewList() List {
	return new(list)
}

// Len Получить длину списка
func (l list) Len() int {
	return l.len
}

// Front Получить первый элемент списка
func (l list) Front() *ListItem {
	return l.first
}

// Back Получить последний элемент списка
func (l list) Back() *ListItem {
	return l.last
}

// PushFront Добавить элемент в начало списка
func (l *list) PushFront(v interface{}) *ListItem {
	item := ListItem{Value: v}
	temp := l.first
	item.Next = temp
	l.first = &item

	if l.len == 0 {
		l.last = l.first
	} else {
		temp.Prev = &item
	}

	l.len++

	return l.first
}

// PushBack Добавить элемент в конец списка
func (l *list) PushBack(v interface{}) *ListItem {
	item := ListItem{Value: v}
	temp := l.last
	item.Prev = temp
	l.last = &item

	if l.len == 0 {
		l.first = l.last
	} else {
		temp.Next = &item
	}

	l.len++

	return l.last
}

// Remove Удалить элемент из списка
func (l *list) Remove(i *ListItem) {
	prev := i.Prev
	next := i.Next

	if prev != nil {
		prev.Next = next
	} else {
		l.first = next
	}

	if next != nil {
		next.Prev = prev
	} else {
		l.last = prev
	}

	l.len--

	i.Prev = nil
	i.Next = nil
}

// MoveToFront Переместить элемент в начало списка
func (l *list) MoveToFront(i *ListItem) {
	if i == l.Front() {
		return
	}

	prev := i.Prev // Предыдущий перемещаемого
	next := i.Next // Следующий перемещаемого

	// Выстраиваем связь предыдущий перемещаемого -> следующий перемещаемого
	if prev != nil {
		prev.Next = next
	}

	// Выстраиваем связь следующий перемещаемого <- предыдущий перемещаемого
	if next != nil {
		next.Prev = prev
	}

	// Перемещаем последний элемент
	if prev != nil && next == nil {
		l.last = prev
	}

	i.Next = nil
	i.Prev = nil

	first := l.first // Первый элемент списка
	first.Prev = i
	i.Next = first
	l.first = i
}
