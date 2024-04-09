package main

import (
    "fmt"
    "runtime"
    //"time"

    "github.com/gotk3/gotk3/glib"
    "github.com/gotk3/gotk3/gtk"
)

func animatedText(frame int) string {
    var sym rune

    switch frame % 4 {
    case 0: sym = '-'
    case 1: sym = '/'
    case 2: sym = '-'
    case 3: sym = '\\'
    }

    text :=            " Очень сложная анимация \n\n\n"

    text += fmt.Sprintf("                  %c\n\n\n", sym)

    return text
}

// Указание типа chan<- вместо chan запрещает операцию чтения из канала
// Этим типом мы показываем, что функция будет использовать этот канал
// только для записи в него.
func makeAnimatedWindow(stopped chan<- struct{}) {
    win, _ := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)

    stop := make(chan struct{}, 1)
    win.Connect("destroy", func() {
        stop <- struct{}{}
    })

    label, _ := gtk.LabelNew(animatedText(0))

    win.Add(label)

    win.ShowAll()

    go func() {
        frame := 1

        for {
            select {
            case <-stop: 
                stopped <- struct{}{}
                return
            default:  
                // эта ветка выбирается, если все остальные ветки
                // в настоящий момент заблокированы
            }

	    //time.Sleep(time.Second)
          
            currentFrame := frame

            // Эта функция откладывает исполнение функции-входа
            // до цикла обработки событий. Напоминаем, что 
            // взаимодействие с элементами Gtk нужно производить
            // одним и тем же исполнителем.
            glib.IdleAdd(func() {
                label.SetLabel(animatedText(currentFrame))
            })

            frame += 1
        }
    }()
}


func main() {
    // Привязываем текущего исполнителя к этой задаче:
    // теперь она не может быть передана ни одному другому исполнителю.
    // Но при этом её исполнитель может временно отвлекаться на 
    // другие задачи.
    runtime.LockOSThread()

    gtk.Init(nil)

    win, _ := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
    win.Connect("destroy", func() {
        gtk.MainQuit()
    })

    win.SetDefaultSize(400, 400)

    button, _ := gtk.ButtonNewWithLabel("Show animated window")
    win.Add(button)
    win.ShowAll()

    stopped := make(chan struct{}, 1)
    stopped <- struct{}{}

    button.Connect("clicked", func() {
        select {
        case <-stopped: makeAnimatedWindow(stopped)
        default:
        }
    })

    gtk.Main()
}
