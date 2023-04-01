package main

// newShoppingList adds a new shopping list to the appData object calling this method.
func (a *appData) newShoppingList(name string) (*shoppingList, error) {
	newShoppingList := &shoppingList{Name: name}
	a.shoppingLists = append(a.shoppingLists, newShoppingList)
	return newShoppingList, nil
}

// deleteShoppingList deletes the shopping list at index `index`
func (a *appData) deleteShoppingList(index int, sl *shoppingList) error {
	if index < len(a.shoppingLists)-1 && index >= 0 {
		a.shoppingLists[index] = a.shoppingLists[len(a.shoppingLists)-1]
	}
	a.shoppingLists = a.shoppingLists[:len(a.shoppingLists)-1]

	return nil
}
