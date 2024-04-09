package main 

import (
	"fmt"
	"github.com/gotk3/gotk3/gtk"
    	"log"
)
var win_flag= 0 
func kill(win *gtk.Window,box *gtk.Box){
	they_dont_know_how_to_live := [](*gtk.Button){}

	box.GetChildren().Foreach(func(child any){
		tmp,_ :=child.(*gtk.Widget).Cast()
		childb, ok := tmp.(*gtk.Button)
		if !ok{return}
		they_dont_know_how_to_live = append(they_dont_know_how_to_live, childb)
	})

	for _,x:= range(they_dont_know_how_to_live){
		box.Remove(x)
	} 

	win.ShowAll()
}

func Menu(win *gtk.Window , box *gtk.Box){
	kill(win,box)
        button_new_game, err := gtk.ButtonNewWithLabel("New Game")
    	button_preferences, err1 := gtk.ButtonNewWithLabel("Preferences")
    	button_exit, err2 := gtk.ButtonNewWithLabel("Exit")
 	
	if err != nil || err1!=nil || err2!= nil {
        	log.Panic(err)
    	}

    	button_new_game.Connect("clicked", func() {
        
    	})

   	 button_preferences.Connect("clicked",func(){
	    preferences(win,box)
	    fmt.Println(win_flag)
   	 })

    	button_exit.Connect("clicked",func(){gtk.MainQuit()})
    	box.Add(button_new_game)
    	box.Add(button_preferences)
    	box.Add(button_exit)
	win.ShowAll()

}

func preferences(win *gtk.Window , box *gtk.Box){
	kill(win,box)
	button_win, err1 := gtk.ButtonNewWithLabel("I want win")
    	button_lose, err2 := gtk.ButtonNewWithLabel("I want lose")
 	
	if err1!=nil || err2!= nil {
        	log.Panic(err1)
    	}
	
	box.Add(button_win)
    	box.Add(button_lose)
	win.ShowAll()
	button_win.Connect("clicked",func(){win_flag=1})
	button_lose.Connect("clicked",func(){win_flag=0})
}

func main() {
    gtk.Init(nil)

    win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)

    if err != nil {
        log.Fatal(err)
    }
    win.Connect("destroy", func() {
        gtk.MainQuit()
    })
    box, err := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 12)

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
    Menu(win,box)
    box.Add(label)
    
    win.ShowAll()

    gtk.Main()
} 
