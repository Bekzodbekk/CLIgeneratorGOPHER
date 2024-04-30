func CLIgenerator(menu []string, functions []func()) {
	var choice int
	fmt.Println("================ Register ================")
	for {
		for idx, i := range menu {
			fmt.Printf("[%d] %s\n", idx+1, i)
		}
		fmt.Print("Tanlang: ")
		fmt.Scan(&choice)

		if choice < 1 || choice > len(functions) {
			fmt.Println("Bunday qator mavjud emas qayta urining!")
			continue
		}
		functions[choice-1]()
	}
}
