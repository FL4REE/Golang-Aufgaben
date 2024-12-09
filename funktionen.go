package main

import (
	"fmt" // Importiert das fmt-Paket für die Ein- und Ausgabe
	"os"  // Importiert das os-Paket für den Zugriff auf Betriebssystemfunktionen
)

// Definiert einen Typ Function, der eine Funktion ohne Parameter und Rückgabewert darstellt
type Function func()

// Erstellt eine Map, die Funktionsnamen mit den entsprechenden Funktionszeigern verknüpft
var functions = map[string]Function{
	"hello":   hello,   // Verknüpft den Namen "hello" mit der Funktion hello
	"goodbye": goodbye, // Verknüpft den Namen "goodbye" mit der Funktion goodbye
	"input":   input,   // Verknüpft den Namen "name" mit der Funktion input
	"zahl":    zahl,    // Verknüpft den Namen "name" mit der Funktion zahl

}

// Definiert die Funktion hello, die "Hello, World!" ausgibt
func hello() {
	fmt.Println("Hello, World!")
}

// Definiert die Funktion goodbye, die "Goodbye, World!" ausgibt
func goodbye() {
	fmt.Println("Goodbye, World!")
}

func input() {
	var name string // Deklariert eine Variable, um den Namen des Benutzers zu speichern

	// Fordert den Benutzer auf, seinen Namen einzugeben
	fmt.Print("Bitte geben Sie Ihren Namen ein: ")
	fmt.Scanln(&name) // Liest die Eingabe des Benutzers und speichert sie in der Variable name

	// Begrüßt den Benutzer mit seinem Namen
	fmt.Printf("Hallo, %s!\n", name) // Gibt eine Begrüßungsnachricht aus
}

func zahl() {
	var number int // Deklariert eine Variable, um die eingegebene Zahl zu speichern

	// Fordert den Benutzer auf, eine Zahl einzugeben
	fmt.Print("Bitte geben Sie eine Zahl ein: ")
	fmt.Scanln(&number) // Liest die Eingabe des Benutzers und speichert sie in der Variable number

	// Überprüft, ob die Zahl zwischen 1 und 100 liegt
	if number >= 1 && number <= 100 {
		fmt.Println("Die Zahl liegt zwischen 1 und 100.") // Gibt eine Bestätigung aus
	} else {
		fmt.Println("Die Zahl liegt nicht zwischen 1 und 100.") // Gibt eine Fehlermeldung aus
	}
}

// Die Hauptfunktion des Programms
func main() {
	// Gibt die verfügbaren Funktionen aus
	fmt.Println("Verfügbare Funktionen:")
	for name := range functions {
		fmt.Println(name) // Listet jeden Funktionsnamen auf
	}

	// Fordert den Benutzer auf, den Namen einer Funktion einzugeben
	fmt.Print("Geben Sie den Namen der Funktion ein, die Sie ausführen möchten: ")
	var input string
	fmt.Scanln(&input) // Liest die Benutzereingabe

	// Überprüft, ob die eingegebene Funktion existiert
	if function, exists := functions[input]; exists {
		function() // Führt die Funktion aus, wenn sie existiert
	} else {
		fmt.Println("Funktion nicht gefunden.") // Gibt eine Fehlermeldung aus, wenn die Funktion nicht existiert
		os.Exit(1)                              // Beendet das Programm mit einem Fehlercode
	}
}
