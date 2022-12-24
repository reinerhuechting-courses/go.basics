# Hallo Welt

Dies ist das erste Programm, das typischerweise beim Lernen einer neuen
Programmiersprache geschrieben wird. Es zeigt den grundsätzlichen Aufbau eines
Programms.

Das eigentliche Programm findet sich in der Datei [hello.go](hello.go).
Dies ist eine sogenannte *Quelldatei*, eine Textdatei, in der sich der *Quelltext* des Progamms befindet.
Dabei handelt es sich i.W. um eine oder mehrere *Funktionen*, die das Programm strukturieren.
Jede dieser Funktionen enthält Anweisungen, die der Computer ausführt, wenn das Programm läuft.

## Erläuterungen zum Programm

Die erste Zeile jeder Go-Quelldatei ist eine Angabe, zu welchem *Package* das Programm gehört.
Packages erlauben es, größere Projekte zu strukturieren.
Für's erste verwenden wir das Package `main`.
Dieses Package existiert in jedem Programm, das am Ende vom Benutzer ausgeführt wird.

```go
package main
```

Der zweite Block in jeder Quelldatei sind die *Import-Statements*.
Hiermit gibt man an, welche Packages (s.o.) man vorhat zu verwenden.
Wir verwenden hier das Package `fmt`, das Funktionen zur formatierten Ein- und Ausgabe
von Text bereitstellt. Diese werden wir weiter unten benutzen, um einen Text auszugeben.

```go
import "fmt"
```

Nun folgt die Funktion `main()`. Anders als die beiden oberen Punkte ist dies kein
verpflichtender Teil jeder Quelldatei, sondern es ist der *Einstiegspunkt* des Programms.
D.h. jedes Programm (aber nicht jede Datei), das vom Benutzer ausgeführt werden soll,
muss genau eine solche Funktion haben. Hier startet die Abarbeitung des Programms,
d.h. die Anweisungen in dieser Funktion werden nacheinander ausgeführt.

```go
func main() {
    fmt.Println("Hello World")
}
```

Die einzige Anweisung in dieser `main()`-Funktion (und damit in diesem Programm)
ist der Aufruf der Funktion `Println()` aus dem Package `fmt`.
Diese Funktion bekommt ein *Argument* , in diesem Fall die Zeichenkette `"Hello World`,
und sie gibt diese Zeichenkette auf der Konsole aus, wo der Benutzer sie lesen kann.
