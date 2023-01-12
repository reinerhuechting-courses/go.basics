package listen

func ExampleIntListDemo() {
	IntListDemo()

	// Output:
	// []
	// []
	// [42]
	// [25 33 105 35 44]
	// [42 25 33 105 35 44]
	// 33
	// 0 : 42
	// 1 : 25
	// 2 : 33
	// 3 : 105
	// 4 : 35
	// 5 : 44
	// 0 : 42
	// 1 : 25
	// 2 : 33
	// 3 : 105
	// 4 : 35
	// 5 : 44
	// 284
	// 5602905000
}

func ExampleStringListDemo() {
	StringListDemo()

	// Output:
	// [Hallo ihr vielen Strings.]
	// Hallo
	// vielen
	// Strings.
	// [vielen Strings.]
}

func ExampleSlicingDemo() {
	SlicingDemo()

	// Output:
	// [33 44 55]
	// [11 22 42 44 55 66 77 88 99]
	// [42 44 55]
	// 22
	// 42
	// 44
	// 55
	// 66
	// 77
	// 88
	// [11 44 84 88 110 132 154 176 99]
	// [11 44 84 110 132 154 176 99]
	// [11 44 84 110 3333 132 154 176 99]
}
