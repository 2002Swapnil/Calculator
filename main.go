// package main

// import "fmt"

// func main() {
// 	fmt.Println("Hello, world.")
// }

// go mod init myapp
// go run main.go
// go get fyne.io/fyne/v2


package main

import (
	"log"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello")

	userName := widget.NewLabel("Username : ")
	password := widget.NewLabel("Password : ")
	userEntry := widget.NewEntry()
	userEntry.SetPlaceHolder("Enter username")
	passEntry := widget.NewEntry()
	passEntry.SetPlaceHolder("Enter password")

	userBox := container.NewHBox(
		userName,
		userEntry,
	)
	passBox :=container.NewHBox(
		password,
		passEntry,
	)

	loginBox := container.NewVBox(
		widget.NewButton("Login", func ()  {
			res := container.NewHBox(
				widget.NewLabel("LOGIN SUCCESSFUL"),
			)
			w.SetContent(res)
			log.Println("LOGIN SUCCESSFUL, You used username : ", userEntry.Text, "and password", passEntry.Text)
		}),
	)

	formBox := container.NewVBox(
		userBox,
		passBox,
		loginBox,
	)

	w.SetContent(formBox);

	w.ShowAndRun()
}