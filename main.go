package main

import (
	"fmt"

	"thebashshell/govento/event"
)

func hello(evt event.Event) {
	fmt.Println("hello", evt.Type)
}

// TestTarget tests EventTarget
// func TestTarget(t *testing.T) {

// }

func main() {
	target := event.NewEventTarget()

	// if !reflect.DeepEqual(target, Target{}) {
	// 	t.FailNow()
	// }

	target.AddEventListener("click", hello, event.ListernerOptions{})
	target.DispatchEvent(event.NewEvent("click", event.InitOptions{}))
}
