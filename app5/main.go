package main

import (
	"fmt"
	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"
	"io"
	"os"

	//"io/ioutil"
	"log"
)

var lastProvider gtk.IStyleProvider

func take_style(buffer *gtk.TextBuffer) {
	start, end := buffer.GetBounds()
	text, err := buffer.GetText(start, end, false)
	if err != nil {
		fmt.Println("GetText Error:", err)
	}
	file, err := os.OpenFile("test.css", os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		log.Panic(err)
	}
	//io.WriteString(f, "Test string")
	//fContent, err := ioutil.ReadFile("test.css")
	//if err != nil {
	//,. panic(err)
	//}
	_, _ = file.Seek(0, io.SeekStart)
	text = "button label {\n " + text + " ;\n}"
	_ = file.Truncate(0)
	io.WriteString(file, text)
	defer file.Close()
}

func style_button(label *gtk.Label, screen *gdk.Screen) {
	//fmt.Println(button.GetStyleContext())
	if lastProvider != nil {
		gtk.RemoveProviderForScreen(screen, lastProvider)
	}
	style_screen, _ := gtk.CssProviderNew()

	err := style_screen.LoadFromPath("test.css")
	gtk.AddProviderForScreen(screen, style_screen, 0)
	//_, err := button.GetStyleContext()
	if err != nil {
		label.SetLabel(fmt.Sprint(err))
	} else {
		label.SetLabel("GOOD!")
	}

	fmt.Println(err)
	lastProvider = style_screen

}
func main() {
	gtk.Init(nil)

	win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)

	if err != nil {
		log.Panic(err)
	}

	win.SetTitle("Hello, world!")

	win.Connect("destroy", func() {
		gtk.MainQuit()
	})
	win.SetSizeRequest(300, 150)

	box, err := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 12)

	if err != nil {
		log.Panic(err)
	}
	screen, _ := gdk.ScreenGetDefault()
	lastProvider = nil
	box.SetMarginTop(24)
	box.SetMarginBottom(24)
	box.SetMarginStart(24)
	box.SetMarginEnd(24)

	win.Add(box)

	label, err := gtk.LabelNew("")

	if err != nil {
		log.Panic(err)
	}
	//entry, err := gtk.EntryNew()
	//if err != nil {
	//   log.Fatal(err)
	//}
	//entry.SetMaxLength(0)
	textview, err := gtk.TextViewNew()
	if err != nil {
		log.Fatal(err)
	}

	scrolled, err := gtk.ScrolledWindowNew(nil, nil) //textview,nil)
	if err != nil {
		log.Fatal(err)
	}

	textview.SetWrapMode(gtk.WRAP_WORD)

	// Добавляем текстовое поле в окно

	box.PackStart(textview, true, false, 5)
	//text, _ := entry.GetText()
	label.SetMarkup("<b>Нажми</b> эту прекрасную кнопку!")

	button, err := gtk.ButtonNewWithLabel("Жми меня")
	//text, err := gtk.TextBufferNew("F")
	if err != nil {
		log.Panic(err)
	}

	button.Connect("clicked", func(b *gtk.Button) {
		text, _ := textview.GetBuffer()
		take_style(text)
		style_button(label, screen)
		box.ShowAll()
	})
	box.Add(textview)
	box.Add(label)
	box.Add(button)
	box.Add(scrolled)

	win.ShowAll()

	gtk.Main()
}
