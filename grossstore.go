package main

import "fmt"

func main() {

	units := Units()
	fmt.Println(units)
	// Output: map[...] with entries like ("dozen": 12)

	bill := NewBill()
	fmt.Println(bill)
	// Output: map[]

	ok := AddItem(bill, units, "tomato", "half_of_a_dozen")
	fmt.Println("AddItem1", ok)
	// Output: true (since dozen is a valid unit)
	//ok = AddItem(bill, units, "tomato", "quarter_of_a_dozen")
	//fmt.Println("AddItem2", ok)

	ok2 := RemoveItem(bill, units, "tomato", "half_of_a_dozen")
	fmt.Println("RemoveItem", ok2)
	// Output: false (because there are no carrots in the bill)

	bill = map[string]int{"carrot": 12, "grapes": 3}
	qty, ok := GetItem(bill, "carrot")
	fmt.Println(qty)
	// Output: 12
	fmt.Println(ok)
	// Output: true
}

//map[KeyType]ElementType

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	units := map[string]int{
		"quarter_of_a_dozen": 3,
		"half_of_a_dozen":    6,
		"dozen":              12,
		"small_gross":        120,
		"gross":              144,
		"great_gross":        1728,
	}
	return units
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return map[string]int{}
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	//bill = NewBill()
	//units = Units()
	u, ok := units[unit]
	if ok == true {
		//fmt.Println("bill", bill)
		//fmt.Println("units", units)
		//fmt.Println("before add", bill)
		bill[item] += u
		//fmt.Println("after add", bill)

	}
	//fmt.Println("AddItem, final", bill)
	return ok

}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	uni, oku := units[unit]
	curr, okb := bill[item]
	qty := curr - uni
	//fmt.Println("qty", qty)
	if oku == false {
		return false
	}
	if okb == false {
		return false
	}
	//fmt.Println("bill before assign", bill)
	if qty < 0 {
		return false
	} else if qty == 0 {
		delete(bill, item)
		//fmt.Println("bill after assign", bill)
	} else {
		//fmt.Println("qty before assign", qty)
		//fmt.Println("bill before assign", bill)
		bill[item] = qty
		//fmt.Println("bill after assign", bill)
	}
	return true
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	qty, ok := bill[item]
	return qty, ok

}
