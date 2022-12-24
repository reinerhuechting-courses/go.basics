# Grundlagen-Beispiele zu Go

Dieses Verzeichnis enthält eine Reihe von Programmen, die den Umgang mit grundlegenden
Programmier-Konzepten verdeutlichen sollen. Ausgehend von einem ganz einfachen
Programm, das den Benutzer nur nach einer Zahl fragt und dieser wieder ausgibt,
werden verschiedene Konzepte und Programmiertechniken eingeführt.

## Überblick

### [Einzelne Zahl einlesen](01-einzelne_zahl)

Fragt den Benutzer nach einer Zahl und gibt diese wieder aus.

### [Drei Zahlen einlesen](02-mehrere_zahlen)

Fragt den Benutzer nach drei Zahlen und gibt deren Summe aus.
Verwendet dabei *Variablen*, um die einzelnen Zahlen zu speichern und zu addieren.

### [Funktion zum Einlesen der Zahlen](03-funktion)

Führt die gleichen Schritte aus wie zuvor, verwendet aber eine *Funktion*
für das Einlesen, um doppelten Code zu vermeiden.

### [Einlesen der Zahlen in einer Schleife](04-schleife)

Führt weiterhin die gleichen Schritte aus, vermeidet aber mehr doppelten Code,
indem eine *Schleife* für Anweiseungen verwendet wird, die sich wiederholen.

### [Einlesen einer variablen Anzahl von Zahlen in einer Schleife](05-schleife_variable_anzahl)

Eine minimale Änderung des vorhergehenden Programms, bei der die Anzahl der
Schleifen-Wiederholungen schon vor der Schleife festgelegt wird.

### [Funktion, die das Einlesen einer beliebigen Anzahl an Variablen überimmt](06-schleife_in_funktion)

Hier wird die vorherige Schleife in eine eigene Funktion verpackt, um sie einfacher
wiederverwendbar bzw. konfigurierbar zu machen. Diese Funktion kann nun eine beliebige
wählbare Anzahl an Zahlen vom Benutzer erfragen und deren Summe berechnen.

### [Kompaktere Schreibweise für das Programm](07-kompakter)

Der gleiche Ablauf wie zuvor, aber kompaktere Schreibweise.

### [While-Schleife, die läuft, bis eine Bedingung erfüllt ist](08-while-schleife)

Eine Variante des vorherigen Programms, bei der die Anzahl der Wiederholungen
nicht im Programm festgelegt ist, sondern von den Eingaben des Benutzers abhängt.
