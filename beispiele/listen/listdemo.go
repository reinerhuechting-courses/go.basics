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
