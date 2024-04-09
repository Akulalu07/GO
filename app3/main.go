package main
import (
	"log"
	"github.com/gotk3/gotk3/gtk"
)
var open = 0
var windows   = [](*gtk.Window){}
func killWindows(){
	for _,x := range( windows){
		x.Destroy()
	}
	windows   = [](*gtk.Window){}
	
}
func END(){
	win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	open++

	if err != nil {
	log.Panic(err)
	}
	win.SetTitle("Купи Слона!")

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

	label.SetMarkup("<b>К сожалению, сейчас у нас нет слона, так что мы не сможем Вам его продать.</b>!")
	box.Add(label)
	win.ShowAll()

	killWindows()

}
func main_menu(){
	win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	open++
	if err != nil {
	log.Panic(err)
	}
	windows = append(windows, win)
	win.SetTitle("Купи Слона!")

	win.Connect("destroy", func() {
	open --
	if open==0{
		gtk.MainQuit()
	}
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

	label.SetMarkup("<b>Купи слона!</b>!")
	box.Add(label)
	button_yes, err := gtk.ButtonNewWithLabel("Да!")

	if err != nil {
		log.Panic(err)
	}
	button_no, err := gtk.ButtonNewWithLabel("Нет!")

	if err != nil {
		log.Panic(err)
	}



    	button_yes.Connect("clicked", func(b *gtk.Button) {
    		END()
	})
	button_no.Connect("clicked", func(b *gtk.Button) {
    		main_menu()
	})

	box.Add(button_yes)
	box.Add(button_no)
	win.ShowAll()
}

func main(){
	gtk.Init(nil)
	main_menu()
    	gtk.Main()
}
