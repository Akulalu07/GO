package main

import (
	"fmt"
	"github.com/gotk3/gotk3/gtk"
	"log"
)

var buttonNumber = 1

func AddNewButton(win *gtk.Window, box *gtk.Box, me *gtk.Button) {
	number := 0
	buttonNumber++
	r:=0
	box.GetChildren().Foreach(func(child any) {
    		// Переменная child связана с интерфейсом, содержащим
    		// величину типа *gtk.Widget

		tmp, _ := child.(*gtk.Widget).Cast()
    		// Переменная tmp связана с интерфейсом, содержащим
    		// элемент правильного типа (например, *gtk.Button)
		
		childButton, ok := tmp.(*gtk.Button) 

		if !ok { return }
		number++
		if childButton.Native() == me.Native(){
			r = number		
		} 

	})
	
	button, err := gtk.ButtonNewWithLabel(fmt.Sprintf("Кнопка%d", buttonNumber))
	if err != nil {
		log.Panic(err)
	}


	button.Connect("clicked", func() {
		AddNewButton(win, box, button)
	})

	box.Add(button)

	box.ReorderChild(button, r+1)

	win.ShowAll()
}

func main() {
	gtk.Init(nil)

	win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		log.Panic(err)
	}
	win.SetTitle("LET'S GO")
	win.Connect("destroy", func() {
		gtk.MainQuit()
	})

	box, err := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 12)
	if err != nil {
		log.Panic(err)
	}
	box.SetMarginTop(24)
	box.SetMarginBottom(24)
	box.SetMarginStart(24)
	box.SetMarginEnd(24)

	win.Add(box)

	label, err := gtk.LabelNew("")
	if err != nil {
		log.Panic(err)
	}
	label.SetMarkup("Ы")

	button, err := gtk.ButtonNewWithLabel(fmt.Sprintf("Кнопка%d", 1))
	if err != nil {
		log.Panic(err)
	}

	button.Connect("clicked", func() {
		AddNewButton(win, box, button)
	})

	box.Add(label)
	box.Add(button)

	win.ShowAll()
	gtk.Main()
}
