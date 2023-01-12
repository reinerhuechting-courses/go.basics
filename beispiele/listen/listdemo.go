package listen

import "fmt"

// Erzeugt Listen (Slices) vom Typ `int` und demonstiert, wie man damit umgehen kann.
func IntListDemo() {

	// Erzeuge zwei leere Listen für Zahlen auf verschiedene Art:
	list1 := make([]int, 0)
	list2 := []int{}

	// Beide Listen auf die Konsole ausgeben:
	fmt.Println(list1)
	fmt.Println(list2)

	// Eine Liste mit einem Element erzeugen und ausgeben:
	list3 := []int{42}
	fmt.Println(list3)

	// Einige Elemente an die bisher leere Liste
	// `list1` anängen und diese wieder ausgeben:
	list1 = append(list1, 25)
	list1 = append(list1, 33)
	list1 = append(list1, 105, 35, 44)
	fmt.Println(list1)

	// Alle Elemente aus `list1` an `list3` anhängen:
	list3 = append(list3, list1...)

	// Die so veränderte `list3` ausgeben:
	fmt.Println(list3)

	// Ein ein einzelnes Element von `list3` ausgeben
	// (hier das dritte Element, d.h. das Element an Stelle 2):
	fmt.Println(list3[2])

	// In einer Schleife alle Elemente von `list3` jeweils mit ihrer Position ausgeben:
	for i := 0; i < len(list3); i++ {
		fmt.Println(i, ":", list3[i])
	}

	// Eine andere Art der Schleife, um alle Elemente auszugeben:
	for i, v := range list3 {
		fmt.Println(i, ":", v)
	}

	// Alle Elemente aus `list3` aufsummieren:
	sum_list3 := 0
	for _, v := range list3 {
		sum_list3 += v
	}
	// Das Ergebnis ausgeben:
	fmt.Println(sum_list3)

	// Alle Elemente aus `list3` aufmultiplizieren:
	product_list3 := 1
	for _, v := range list3 {
		product_list3 *= v
	}
	// Das Ergebnis ausgeben:
	fmt.Println(product_list3)
}

// Erzeugt Listen (Slices) vom Typ `string` und demonstiert, wie man damit umgehen kann.
func StringListDemo() {
	// Erzeuge eine leere String-Liste:
	list1 := []string{}

	// Hänge Elemente an:
	list1 = append(list1, "Hallo")
	list1 = append(list1, "ihr")
	list1 = append(list1, "vielen")
	list1 = append(list1, "Strings.")

	// Liste ausgeben:
	fmt.Println(list1)

	// Alle Strings aus `list1` ausgeben, deren Länge größer als 4 ist:
	for _, element := range list1 {
		if len(element) > 4 {
			fmt.Println(element)
		}
	}

	// Alle Strings aus `list1` in eine neue Liste kopieren, deren Länge größer als 5 ist:
	list2 := []string{}
	for _, element := range list1 {
		if len(element) > 5 {
			list2 = append(list2, element)
		}
	}
	// Die neue Liste ausgeben:
	fmt.Println(list2)
}

// Demonstriert, wie man Ausschnitte aus anderen Listen holt und damit arbeitet.
func SlicingDemo() {
	// Eine Liste von Zahlen erstellen:
	list1 := []int{11, 22, 33, 44, 55, 66, 77, 88, 99}

	// Einen Ausschnitt aus `list1` wählen und als `list2` speichern:
	list2 := list1[2:5]

	// `list2` ausgeben:
	fmt.Println(list2)

	// Ein Element von `list2` verändern:
	list2[0] = 42

	// Beide Listen ausgeben.
	// Beachte: Es haben sich beide geändert!
	fmt.Println(list1)
	fmt.Println(list2)

	// Alle Elemente aus `list1` ausgeben, außer dem ersten und dem letzten Element:
	for _, v := range list1[1 : len(list1)-1] {
		fmt.Println(v)
	}

	// Alle Elemente aus `list1` verdoppeln, außer dem ersten und dem letzten Element:
	for i := range list1[1 : len(list1)-1] {
		list1[i+1] *= 2
	}
	// Die Liste ausgeben:
	fmt.Println(list1)

	// Das Element an Stelle 3 aus `list1` löschen:
	list1 = append(list1[:3], list1[4:]...)
	// Die Liste ausgeben:
	fmt.Println(list1)

	// Ein Element in List an Stelle 4 einfügen:
	end := []int{}
	end = append(end, list1[4:]...)
	list1 = append(list1[:4], 3333)
	list1 = append(list1, end...)
	// Die Liste ausgeben:
	fmt.Println(list1)
}
