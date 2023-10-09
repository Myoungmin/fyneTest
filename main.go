// package main

// import (
// 	"fmt"

// 	"fyne.io/fyne/v2/app"
// 	"fyne.io/fyne/v2/container"
// 	"fyne.io/fyne/v2/widget"
// 	"github.com/Myoungmin/fyneTest/goLib"
// )

// func main() {
// 	myApp := app.New()
// 	myWindow := myApp.NewWindow("Hello")

// 	result := goLib.Add(3, 4) // mygolib 패키지에서 Add 함수 호출
// 	hello := widget.NewLabel(fmt.Sprintf("Result from C++: %d", result))
// 	myWindow.SetContent(container.NewVBox(hello, widget.NewButton("Quit", func() {
// 		myApp.Quit()
// 	})))

// 	myWindow.ShowAndRun()
// }

package main

import (
	"fmt"

	"github.com/Myoungmin/fyneTest/cppExample"
)

func main() {
	result := cppExample.Add(3, 4)
	fmt.Println(result)
}
