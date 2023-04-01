package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type item struct {
	What    string
	Checked bool
}

type shoppingList struct {
	Name  string
	Items []item

	list        *widget.List
	filterEntry *widget.Entry
}

type appData struct {
	shoppingLists []*shoppingList

	app  fyne.App
	win  fyne.Window
	tabs *container.DocTabs
}

func main() {
	a := app.NewWithID("github.com.ganaeld.shopping")

	myApp := &appData{shoppingLists: []*shoppingList{},
		app: a,
		win: a.NewWindow("Shopping List")}

	myApp.tabs = container.NewDocTabs(nil)
	myApp.tabs.CreateTab = myApp.createTab
	myApp.tabs.OnClosed = func(item *container.TabItem) {
		for index, value := range myApp.shoppingLists {
			if value.Name == item.Text {
				myApp.deleteShoppingList(index, value)
				return
			}
		}
	}
	myApp.tabs.SetTabLocation(container.TabLocationLeading)

	myApp.win.SetContent(container.NewMax(myApp.tabs))
	myApp.win.Resize(fyne.NewSize(800, 600))
	myApp.win.ShowAndRun()
}
