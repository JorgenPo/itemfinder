package main

import (
	"findthing/core"
	"findthing/types"
	"fmt"
	"os"
)

const (
	VERSION_MAJOR = 0
	VERSION_MINOR = 2
	VERSION_PATCH = 1
)

func printHelp() {
	fmt.Println("Avito and other sites parser cli interface")
	fmt.Printf("Usage: findthings [query]\n")
	fmt.Println("  query - string query with item keywords or name for search")
}

func printResult(n int, item *types.Item) {
	fmt.Printf("================   Result #%v '%v'   ================\n", n+1, item.Title)
	fmt.Println()
	fmt.Printf("Price %v rubles\n", item.Price)
	fmt.Println()
	fmt.Println(item.Address)
	fmt.Printf("Telephone: %v\n", item.PhoneNumber)

	fmt.Println()

	fmt.Printf("Seller info:\n\n")
	fmt.Println(item.User.Name)

	if item.User.IsCompany {
		fmt.Println("Organization")
	} else {
		fmt.Println("Not organization")
	}

	fmt.Println()
	fmt.Printf("Description:\n")
	fmt.Println(item.Description)

	fmt.Printf("================ End Result #%v '%v' ================\n", n+1, item.Title)
}

func main() {
	fmt.Printf("findthings v%v.%v.%v\n", VERSION_MAJOR, VERSION_MINOR, VERSION_PATCH)


	if len(os.Args) < 2 {
		printHelp()
		os.Exit(1)
	}

	q := types.Query{City: "sankt-peterburg", Query: os.Args[1]}

	finder := core.NewFinder()
	results := finder.Find(q)

	fmt.Println()
	fmt.Println()
	for number, result := range results {
		printResult(number, result)
		fmt.Println()
	}

}