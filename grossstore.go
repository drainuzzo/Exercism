package main

import "fmt"

func main() {

	units := Units()
	fmt.Println(units)
	// Output: map[...] with entries like ("dozen": 12)

	bill := NewBill()
	fmt.Println(bill)
	// Output: map[]

	ok := AddItem(bill, units, "carrot", "dozen")
	fmt.Println(ok)
	// Output: true (since dozen is a valid unit)

	ok2 := RemoveItem(bill, units, "carrot", "dozen")
	fmt.Println(ok2)
	// Output: false (because there are no carrots in the bill)

	cill := map[string]int{"carrot": 12, "grapes": 3}
	qty, ok := GetItem(cill, "carrot")
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
	bill = NewBill()
	units = Units()
	_, ok := units[unit]
	if ok == true {
		if bill[item] == units[unit] {
			bill[item] += units[unit]
		} else {
			bill[item] = units[unit]
		}
	}
	//fmt.Println(bill)
	return ok

}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	units = Units()
	_, oku := units[unit]
	_, okb := bill[item]
	qty, _ := GetItem(bill, item)
	//return oku && oku
	if (okb == false || oku == false) && qty < 0 {
		return false
	} else if qty == 0 {
		delete(bill, item)
		return true
	} else {
		bill[item]--
		return true
	}
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	qty, ok := bill[item]
	return qty, ok

}
