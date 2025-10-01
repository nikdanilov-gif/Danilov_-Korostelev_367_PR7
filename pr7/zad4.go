package main

import "fmt"

type Order struct {
	zid int
	client  Customer
	items  []OrderItem
	status  string
}

type OrderItem struct {
	pid  int
	namep  string
	kolich  int
	price float64
}

type Customer struct {
	id int
	name  string
	mail  string
}

func (o *Order) total() float64 {
	sum := 0.0
	for _, v := range o.items {
		sum += float64(v.kolich) * v.price
	}
	return sum
}

func (o *Order) add(pid int, namep string, kolich int, price float64) {
	item := OrderItem{pid, namep, kolich, price}
	o.items = append(o.items, item)
}

func (o *Order) remove(p int) {
	for idx, v := range o.items {
		if v.pid == p {
			o.items = append(o.items[:idx], o.items[idx+1:]...)
			return
		}
	}
}

func (o *Order) change(s string) {
	o.status = s
}

func main() {
	u := Customer{1, "abc", "a@b.c"}
	ord := Order{1, u, []OrderItem{}, "new"}

	ord.add(52, "Лабубу", 3, 10000.0)
	ord.add(12, "Хагиваги", 4, 500.0)
	ord.add(38, "Букварь", 5, 100.0)
	ord.add(22, "ЛадаГранта", 1, 250000.0)

	fmt.Printf("Общая сумма: %.2f\n", ord.total())

	ord.remove(22)
	ord.remove(38)
	ord.change("done")

	fmt.Println("Сумма после удаления")

	fmt.Printf("Общая сумма: %.2f\n", ord.total())
}
