package main

import (
	"fmt"

	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
)

func addEvent() {

	eventHook := robotgo.Start()

	var e hook.Event
	var key []string

	fmt.Println("--- Please press ctrl + shift + q to stop hook ---")
	fmt.Println("--- Please press w---")

	for e = range eventHook {
		if e.Kind == hook.KeyDown {
			fmt.Println(key)

			switch key {
			case "W":
				fmt.Println("--- Please press w---")
			case "w":
				fmt.Println("--- Please press ctrl + shift + q to stop hook ---")
			default:
				fmt.Printf("pressed %s \n", key)
			}

		}
	}

	// robotgo.EventHook(hook.KeyDown, []string{"q", "ctrl", "shift"}, func(e hook.Event) {
	// 	fmt.Println("ctrl-shift-q")
	// 	robotgo.EventEnd()
	// })

	// fmt.Println("--- Please press w---")
	// robotgo.EventHook(hook.KeyDown, []string{"w"}, func(e hook.Event) {
	// 	robotgo.TypeStrDelay("test", 2)
	// })

	// s := robotgo.EventStart()
	// <-robotgo.EventProcess(s)
}

// func addMouse() {
// 	fmt.Println("--- Please press left mouse button to see it's position and the right mouse button to exit ---")
// 	robotgo.EventHook(hook.MouseDown, []string{}, func(e hook.Event) {
// 		if e.Button == hook.MouseMap["left"] {
// 			fmt.Printf("mouse left @ %v - %v\n", e.X, e.Y)
// 		} else if e.Button == hook.MouseMap["right"] {
// 			robotgo.EventEnd()
// 		}
// 	})

// 	s := robotgo.EventStart()
// 	<-robotgo.EventProcess(s)
// }

func lowLevel() {
	////////////////////////////////////////////////////////////////////////////////
	// Global event listener
	////////////////////////////////////////////////////////////////////////////////
	fmt.Println("Press q to stop event gathering")
	evChan := robotgo.EventStart()
	for e := range evChan {
		fmt.Println(e)
		if e.Keychar == 'q' {
			robotgo.EventEnd()
			// break
		}
	}
}

func main() {
	fmt.Println("test begin...")
	addEvent()

	// addMouse()

	lowLevel()
}
