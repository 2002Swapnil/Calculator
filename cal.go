package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/Knetic/govaluate"
	"strconv"
)

func main() {
	a := app.New()
	w := a.NewWindow("Calculator")
	output:= ""
	input := widget.NewLabel(output)
	history :=widget.NewButton("history",func(){
	})
	back :=widget.NewButton("<=",func(){
		if len(output) >0{
		output=output[:len(output)-1];
		input.SetText(output);
		}
		
	})

	clear :=widget.NewButton("AC",func(){
		output="";
		input.SetText(output);
	})
	openBracket :=widget.NewButton("(",func(){
		output=output+"(";
		input.SetText(output);
	})
	closeBracket :=widget.NewButton(")",func(){
		output=output+")";
		input.SetText(output);
	})
	divide :=widget.NewButton("/",func(){
		output=output+"/";
		input.SetText(output);
	})
	seven :=widget.NewButton("7",func(){
		output+="7";
		input.SetText(output);
	})
	eight :=widget.NewButton("8",func(){
		output+="8"
		input.SetText(output)
	})
	nine :=widget.NewButton("9",func(){
		output+="9"
		input.SetText(output)
	})
	multiply :=widget.NewButton("*",func(){
		output+="*"
		input.SetText(output)
	})
	four :=widget.NewButton("4",func(){
		output+="4"
		input.SetText(output)
	})
	five :=widget.NewButton("5",func(){
		output+="5"
		input.SetText(output)
	})
	six :=widget.NewButton("6",func(){
		output+="6"
		input.SetText(output)
	})
	minus :=widget.NewButton("-",func(){
		output+="-"
		input.SetText(output)
	})
	one :=widget.NewButton("1",func(){
		output+="1"
		input.SetText(output)
	})
	two :=widget.NewButton("2",func(){
		output+="2"
		input.SetText(output)
	})
	three :=widget.NewButton("3",func(){
		output+="3"
		input.SetText(output)
	})
	plus :=widget.NewButton("+",func(){
		output+="+"
		input.SetText(output)
	})
	zero :=widget.NewButton("0",func(){
		output+="0"
		input.SetText(output)
	})
	dot :=widget.NewButton(".",func(){
		output+="."
		input.SetText(output)
	})
	equal :=widget.NewButton("=",func(){
		expression, err := govaluate.NewEvaluableExpression(output);
		if err==nil {
			result, err := expression.Evaluate(nil);
		
			if err==nil{
				output = strconv.FormatFloat(result.(float64), 'f', -1, 64);
			} else{
				output= "error";
			}
		} else{
			output="error";
		}
		input.SetText(output);
	})
	
	
	w.SetContent(container.NewVBox(
		input,
		container.NewGridWithColumns(1,
			container.NewGridWithColumns(2,
			history,
			back,),

			container.NewGridWithColumns(4,
			clear,
			openBracket,
			closeBracket,
			divide,),

			container.NewGridWithColumns(4,
			seven,
			eight,
			nine,
			multiply,),

			container.NewGridWithColumns(4,
			four,
			five,
			six,
			minus,),

			container.NewGridWithColumns(4,
			one,
			two,
			three,
			plus,),

		),
		container.NewGridWithColumns(2,
			container.NewGridWithColumns(2,
			zero,
			dot,
			),
			equal,
		),	

		
	),)

		
	w.ShowAndRun()
}