//go:build gui

// Windowed version: the portable two-button installer (Inyectar / Restaurar).
// Uses lxn/walk (a pure-Go Win32 GUI, no CGO). Built in CI / on Windows with:
//     go get github.com/lxn/walk github.com/lxn/win
//     go build -tags gui -ldflags "-H windowsgui -s -w" -o Veilcord.exe .

package main

import (
	"fmt"

	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"

	"veilcord-installer/internal/injector"
)

func main() {
	var logBox *walk.TextEdit

	logln := func(s string) {
		if logBox != nil {
			logBox.AppendText(s + "\r\n")
		}
	}

	run := func(install bool) {
		ds := injector.FindDiscords()
		if len(ds) == 0 {
			logln("No encontre Discord. Abrilo una vez y reintenta.")
			return
		}
		for _, d := range ds {
			var err error
			if install {
				err = injector.Install(d)
			} else {
				err = injector.Uninstall(d)
			}
			if err != nil {
				logln(fmt.Sprintf("%s: ERROR: %v", d.Branch, err))
				continue
			}
			if install {
				logln(d.Branch + ": Veilcord inyectado")
			} else {
				logln(d.Branch + ": restaurado")
			}
		}
		if install {
			logln("Abri Discord. Para confirmar: %TEMP%\\veilcord-injected.txt")
		}
	}

	_, _ = MainWindow{
		Title:   "Veilcord",
		MinSize: Size{Width: 440, Height: 340},
		Size:    Size{Width: 440, Height: 340},
		Layout:  VBox{MarginsZero: false},
		Children: []Widget{
			Label{Text: "Veilcord", Font: Font{PointSize: 18, Bold: true}},
			Label{Text: "Modifica tu Discord. Se cierra solo al inyectar."},
			Composite{
				Layout: HBox{},
				Children: []Widget{
					PushButton{Text: "Inyectar", OnClicked: func() { run(true) }},
					PushButton{Text: "Restaurar", OnClicked: func() { run(false) }},
				},
			},
			TextEdit{AssignTo: &logBox, ReadOnly: true, VScroll: true},
		},
	}.Run()
}
