package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	units := map[string]int{}
    units["quarter_of_a_dozen"] = 3
    units["half_of_a_dozen"] = 6
    units["dozen"] = 12
    units["small_gross"] = 120
    units["gross"] = 144
    units["great_gross"] = 1728
    return units
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	newBill := map[string]int{}
    return newBill
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	quantity, ok := units[unit]
    if !ok {
        return false
    }
    bill[item] += quantity
    return true
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	if _, exists := bill[item]; !exists{
        return false
    }
    
	removeQty, exists := units[unit]
	if !exists {
		return false
	}

	if removeQty < 0 {
		return false
	}

	currentQty := bill[item]
	newQty := currentQty - removeQty

	if newQty < 0 {
		return false
	}

	if newQty == 0 {
		delete(bill, item)
	} else {
		bill[item] = newQty
	}
	return true
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	quantity, exists := bill[item]
    if !exists {
        return 0, false
    }
    return quantity, true
}
