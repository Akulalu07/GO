package main

import (
    "github.com/gotk3/gotk3/gtk"
    "log"
    "fmt"
    "math/big"
)

func main() {
    gtk.Init(nil)

    builder, err := gtk.BuilderNewFromFile("ui/ui_for_app.glade")
    
    var opInput bool = true
    var lastOp func(*big.Int, *big.Int)*big.Int 
    lastRes := big.NewInt(0)
    var digits string
    if err != nil {
        log.Fatal(err)
    }
    
    obj, _ := builder.GetObject("Display")
    obj.(*gtk.Label).SetMarkup("0")
    //strconv.Atoi()
    swap := func(){
        if !opInput{
            if string(digits[0])!="-"{
                digits = "-" + digits
            }else{
                digits = digits[1:]
            }
            obj.(*gtk.Label).SetMarkup(digits)
        }else{
            lastRes.Neg(lastRes)
            obj.(*gtk.Label).SetMarkup(lastRes.String())
        }
    }
    genDigitHandler := func(digit string) func() {
        return func() {
            if !opInput {
                if digits=="0"{
                    digits = digit
                }else{
                    if digits !="0"{
                        digits += digit
                    }
                }
                obj.(*gtk.Label).SetMarkup(digits)
            } else {
                opInput = !opInput
                digits = digit
                obj.(*gtk.Label).SetMarkup(digits)
            }
        }
    }
    genOpHandler := func(newOp func(*big.Int, *big.Int)*big.Int) func() {
        return func() {
            if !opInput {
    		c := big.NewInt(0)
                c.SetString(digits, 10)
		fmt.Println("lastRes", lastRes, "c", c)
                lastRes = lastOp(lastRes, c)
		fmt.Println("new lastRes", lastRes)
                lastOp = newOp
                opInput = !opInput
                obj.(*gtk.Label).SetMarkup(lastRes.String())
            } else {
                lastOp = newOp
            }
        }
    }

    sub := func(x *big.Int, y *big.Int)*big.Int {
        y.Neg(y)
        n := big.NewInt(0)
        n.Add(x,y)
        return n
    }
    mul := func(x *big.Int, y *big.Int)*big.Int {
        n := big.NewInt(0)
        n.Mul(x,y)
        return n
    }
    add := func(x *big.Int, y *big.Int)*big.Int {
        n := big.NewInt(0)
        n.Add(x,y)
        return n
    }
    div := func(x *big.Int, y *big.Int)*big.Int {
        if y==big.NewInt(0){
            return big.NewInt(0)
        }
        n := big.NewInt(0)
        n.Div(x,y)
        return n
    }
    del := func(){
        if !opInput{
            digits = "0"
            obj.(*gtk.Label).SetMarkup(digits)
        } else {
            opInput = !opInput
            digits = "0"
            obj.(*gtk.Label).SetMarkup(digits)
        }
    }

    exe := func(x *big.Int, y *big.Int)*big.Int{
        return y
    } 
    
    delLast := func() {
        lastOp = exe
        digits = "0"
        opInput = true
    	lastRes = big.NewInt(0)
        obj.(*gtk.Label).SetMarkup(digits)
    }


    lastOp = exe

    builder.ConnectSignals(map[string]any{
        "del": del,
	"del_last": delLast,
        "swap": swap,
        "null": genDigitHandler("0"),
        "one": genDigitHandler("1"),
        "two": genDigitHandler("2"),
        "three": genDigitHandler("3"),
        "four": genDigitHandler("4"),
        "five": genDigitHandler("5"),
        "six": genDigitHandler("6"),
        "seven": genDigitHandler("7"),
        "eight": genDigitHandler("8"),
        "nine": genDigitHandler("9"),
        "sub": genOpHandler(sub),
        "mul": genOpHandler(mul),
        "add": genOpHandler(add),
        "div": genOpHandler(div),
        "exe":  genOpHandler(exe),
    })

    obj1, _ := builder.GetObject("TopLevel")

    win := obj1.(*gtk.Window)

    win.SetDefaultSize(400, 400)

    win.ShowAll()

    // TODO win.Connect("destroy", ...)

    gtk.Main()
}
