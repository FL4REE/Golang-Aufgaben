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
	"name":    name,    // Verknüpft den Namen "name" mit der Funktion name
}

// Definiert die Funktion hello, die "Hello, World!" ausgibt
func hello() {
	fmt.Println("Hello, World!")
}

// Definiert die Funktion goodbye, die "Goodbye, World!" ausgibt
func goodbye() {
	fmt.Println("Goodbye, World!")
}

func name() {
	var name string // Deklariert eine Variable, um den Namen des Benutzers zu speichern

	// Fordert den Benutzer auf, seinen Namen einzugeben
	fmt.Print("Bitte geben Sie Ihren Namen ein: ")
	fmt.Scanln(&name) // Liest die Eingabe des Benutzers und speichert sie in der Variable name

	// Begrüßt den Benutzer mit seinem Namen
	fmt.Printf("Hallo, %s!\n", name) // Gibt eine Begrüßungsnachricht aus
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
