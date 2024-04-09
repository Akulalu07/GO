package main 
import (
	"github.com/gotk3/gotk3/gtk"
    	"log"
)
type Menu struct {
    entries []MenuEntry
}

func (m *Menu) AddEntryWithAction(label string, next *Menu, action func()) {
    if next == nil { next = m }

    m.entries = append(m.entries, MenuEntry{
        Label: label,
        Next: next,
        Action: action,
    })
}

func (m *Menu) AddEntry(label string, next *Menu) {
    m.AddEntryWithAction(label, next, nil)
}

type MenuEntry struct {
    Label string
    Next *Menu
    Action func()
}

func (e MenuEntry) Use() *Menu {
    if e.Action != nil { e.Action() }

    return e.Next
}
func kill(box *gtk.Box){
	they_dont_know_how_to_live := [](*gtk.Button){}
	box.GetChildren().Foreach(func(child any){
		tmp,_ :=child.(*gtk.Widget).Cast()
		childb, ok := tmp.(*gtk.Button)
		if !ok{return;}
		they_dont_know_how_to_live = append(they_dont_know_how_to_live, childb)
	})

	for _,x:= range(they_dont_know_how_to_live){
		box.Remove(x)
	} 
}


func (m Menu) GtkWidgetRec(parentBox *gtk.Box) *gtk.Widget{
	 if (parentBox == nil) {
		 box, err := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 12)

		 if err != nil {
			log.Panic(err)
		}

		box.SetMarginTop(24)
		box.SetMarginBottom(24)
		box.SetMarginStart(24)
		box.SetMarginEnd(24)
		parentBox = box
	}

	kill(parentBox)

	for _, x :=  range(m.entries){
		function := x.Action
		label := x.Label 
		next := x.Next
		button, err := gtk.ButtonNewWithLabel(label)
   		 
    		if err != nil {
        		log.Panic(err)
	    	}	

		parentBox.Add(button)
		button.Connect("clicked", func(b *gtk.Button) {
			if function != nil { function() }

			if next == nil { return }
			next.GtkWidgetRec(parentBox)
			parentBox.ShowAll()
    		})
	}
	return &parentBox.Widget
}

func (m Menu) GtkWidget() *gtk.Widget {
	return m.GtkWidgetRec(nil)
}

func makeMainMenu(info *gtk.Label) *Menu {
    var mainMenu, newGameMenu, optionsMenu Menu
    
    var gameResult bool
    
    playGame := func() {
        if gameResult {
            info.SetText("Вы выиграли!")
        } else {
            info.SetText("Вы проиграли!")
        }
    }

    clearInfo := func() { info.SetText("") }
    
    mainMenu.AddEntryWithAction("Новая игра", &newGameMenu, playGame)
    
    mainMenu.AddEntry("Настройки", &optionsMenu)
    mainMenu.AddEntryWithAction("Выйти", nil, gtk.MainQuit)
    
    newGameMenu.AddEntryWithAction("Начать заново", nil, playGame)
    newGameMenu.AddEntryWithAction("Выйти в главное меню", &mainMenu, clearInfo)
    
    optionsMenu.AddEntryWithAction("Хочу всегда выигрывать", &mainMenu, func() {
        gameResult = true
    })
    
    optionsMenu.AddEntryWithAction("Хочу всегда проигрывать", &mainMenu, func() {
        gameResult = false
    })
    
    return &mainMenu
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
    
    box.Add(label)
    box.Add(makeMainMenu(label).GtkWidget())

    win.ShowAll()

    gtk.Main()
}
