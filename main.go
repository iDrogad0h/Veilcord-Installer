//go:build !gui

// Console version of the installer. This is the default build and uses only
// the Go standard library, so it compiles to a Windows .exe with zero
// dependencies. The pretty two-button window lives in gui.go (build tag: gui).

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"veilcord-installer/internal/injector"
)

func main() {
	fmt.Println("========================================")
	fmt.Println("   Veilcord  -  Instalador (prueba)")
	fmt.Println("========================================")
	fmt.Println()

	discords := injector.FindDiscords()
	if len(discords) == 0 {
		fmt.Println("No encontre ninguna instalacion de Discord.")
		fmt.Println("Abri Discord al menos una vez y volve a intentar.")
		pause()
		return
	}

	fmt.Println("Discord encontrado:")
	for _, d := range discords {
		status := "normal (sin Veilcord)"
		if d.IsInstalled() {
			status = "Veilcord YA instalado"
		}
		fmt.Printf("   - %-20s %s\n", d.Branch, status)
	}
	fmt.Println()
	fmt.Println("Que queres hacer?")
	fmt.Println("   1) Inyectar   (instalar Veilcord)")
	fmt.Println("   2) Restaurar  (sacar Veilcord, dejar Discord normal)")
	fmt.Println("   3) Salir")
	fmt.Print("\nEscribi 1, 2 o 3 y apreta Enter: ")

	switch strings.TrimSpace(readLine()) {
	case "1":
		for _, d := range discords {
			fmt.Printf("Inyectando en %s ... ", d.Branch)
			if err := injector.Install(d); err != nil {
				fmt.Println("ERROR:", err)
			} else {
				fmt.Println("listo")
			}
		}
		fmt.Println("\nAbri Discord ahora.")
		fmt.Println("Para confirmar que funciono, fijate que exista este archivo:")
		fmt.Println("   %TEMP%\\veilcord-injected.txt")
	case "2":
		for _, d := range discords {
			fmt.Printf("Restaurando %s ... ", d.Branch)
			if err := injector.Uninstall(d); err != nil {
				fmt.Println("ERROR:", err)
			} else {
				fmt.Println("listo")
			}
		}
		fmt.Println("\nDiscord quedo como nuevo.")
	default:
		fmt.Println("Saliendo.")
	}
	pause()
}

func readLine() string {
	s, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	return s
}

func pause() {
	fmt.Print("\nApreta Enter para cerrar...")
	readLine()
}
