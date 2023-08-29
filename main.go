package main

import (
    "fmt"

	"github.com/micmonay/keybd_event"
	"github.com/nsf/termbox-go"
    "gitlab.com/gomidi/midi"
    . "gitlab.com/gomidi/midi/midimessage/channel" // (Channel Messages)
    "gitlab.com/gomidi/midi/reader"
    "gitlab.com/gomidi/rtmididrv"
)

// This example reads from the first input port
func main() {
	err := termbox.Init()
	checkErr(err)
	defer termbox.Close()

    drv, err := rtmididrv.New()
    checkErr(err)

    // make sure to close the driver at the end
    defer drv.Close()

	kb, err := keybd_event.NewKeyBonding()

    ins, err := drv.Ins()
    checkErr(err)

    // takes the first input
    in := ins[1]

    fmt.Printf("opening MIDI Port %v\n", in)
    checkErr(in.Open())

    defer in.Close()


	eventQueue := make(chan termbox.Event)
    go func() {
		for {
			eventQueue <- termbox.PollEvent()
		}
	}()
	// to disable logging, pass mid.NoLogger() as option
    rd := reader.New(
        reader.NoLogger(),
        // print every message
        reader.Each(func(pos *reader.Position, msg midi.Message) {

            shift := false
			kb.HasSHIFT(shift)
            switch v := msg.(type) {
            case NoteOn:
				fmt.Printf("%d\n", v.Key())
				switch v.Key() {
				case 60:
					kb.SetKeys(keybd_event.VK_SPACE)
					kb.Press()
				case 58:
					kb.SetKeys(keybd_event.VK_BACKSPACE)
					kb.Press()
				case 59:
					kb.SetKeys(keybd_event.VK_D)
					kb.Press()
				case 56:
					kb.SetKeys(42)
					kb.Press()
				case 57:
					kb.SetKeys(keybd_event.VK_W)
					kb.Press()
				case 55:
					kb.SetKeys(keybd_event.VK_S)
					kb.Press()
				case 53:
					kb.SetKeys(keybd_event.VK_A)
					kb.Press()
				case 54:
					kb.SetKeys(29)
					kb.Press()
				case 61:
					kb.SetKeys(keybd_event.VK_E)
					kb.Press()
				}

            case NoteOff:
                fmt.Printf("%d\n", v.Key())
				switch v.Key() {
				case 60:
					kb.SetKeys(keybd_event.VK_SPACE)
					kb.Release()
				case 68:
					kb.SetKeys(keybd_event.VK_BACKSPACE)
					kb.Release()
				case 59:
					kb.SetKeys(keybd_event.VK_D)
					kb.Release()
				case 56:
					kb.SetKeys(42)
					kb.Release()
				case 57:
					kb.SetKeys(keybd_event.VK_W)
					kb.Release()
				case 55:
					kb.SetKeys(keybd_event.VK_S)
					kb.Release()
				case 53:
					kb.SetKeys(keybd_event.VK_A)
					kb.Release()
				case 54:
					kb.SetKeys(29)
					kb.Release()
				case 61:
					kb.SetKeys(keybd_event.VK_E)
					kb.Release()
				}
				
            }
        }),
    )

	
	poll := true
    // listen for MIDI
	for poll {
		err = rd.ListenTo(in)
		//checkErr(err)
		ev := <- eventQueue
		if ev.Type == termbox.EventKey {
			if ev.Key == termbox.KeyEsc {
				poll = false
			}
		}
	}

    err = in.StopListening()
    checkErr(err)
    fmt.Printf("closing MIDI Port %v\n", in)
}

func checkErr(err error) {
    if err != nil {
        panic(err.Error())
    }
}
func deferTest(){
	fmt.Println("e")
}