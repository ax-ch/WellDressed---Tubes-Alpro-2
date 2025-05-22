package main

import "fmt"

type ClothingItem struct {
	ID        int
	Name      string
	Color     string
	Category  string
	Formality int
	LastWorn  string
	Weather   string
} 

const NMAX int = 77
type wardrobe [NMAX]ClothingItem

func main() {
	var item wardrobe
	var total int

	welcome()
	readNumberOfClothingItems(&total)
	addClothingItem(&item, total)
	viewAllClothingItems(item)
	fmt.Println("Wardrobe setup complete!")
	fmt.Println("You can now edit, sort, search, and get outfit recommendations.")
	fmt.Println()

	menuChoice(&item, &total)
	fmt.Println("Thank you for using WellDressed! Goodbye!")
	fmt.Println("=========================================================================")
}

func welcome() {
	var userName string

	fmt.Print("Enter your name: ")
	fmt.Scan(&userName)
	fmt.Println("=========================================================================")
	fmt.Println(" /\\_/\\  Hello %s, welcome to WellDressed!", userName)
	fmt.Println("( o.o )  This is your digital wardrobe.")
	fmt.Println(" > ^ <  You can add, modify, remove, and sort your clothing items.")
	fmt.Println("You can also search for clothing items and get outfit recommendations.")
	fmt.Println("Let's get started!")
	fmt.Println("=========================================================================")
	fmt.Println()
}

func readNumberOfClothingItems(total *int) {
	fmt.Println("First, let's set up your wardrobe.")
	fmt.Print("Enter the number of clothing items in your wardrobe: ")
	fmt.Scan(&total)

	if *total < 1 {
		fmt.Println("Invalid number of clothing items. Please enter a positive integer.")
		readNumberOfClothingItems(total)
	}
}

func addClothingItem(item *wardrobe, total int) {
	var NextID int = 1

	for i := 0; i < total; i++ {
		(*item)[i].ID = NextID

		fmt.Println("Choose clothing category:")
		chooseCategory(item, i)

		fmt.Print("Enter clothing name: ")
		enterClothingName(item, i)

		fmt.Print("Enter clothing color (one word, e.g., blue, red): ")
		enterClothingColor(item, i)

		fmt.Println("Choose the suitable weather category")
		chooseWeather(item, i)

		fmt.Print("Enter clothing formality (1-10): ")
		enterClothingFormality(item, i)

		fmt.Print("Enter last worn date (YYYY-MM-DD): ")
		enterClothingLastWorn(item, i)

		NextID++
	}
}

func chooseCategory(item *wardrobe, index int) {
	var categoryChoice int

	fmt.Println("1. Top")
	fmt.Println("2. Bottoms")
	fmt.Println("3. Dress")
	fmt.Println("4. Shoes")
	fmt.Println("5. Accessories")
	fmt.Print("Enter your choice (1-5): ")

	fmt.Scan(&categoryChoice)
	switch categoryChoice {
	case 1:
		(*item)[index].Category = "Top"
	case 2:
		(*item)[index].Category = "Bottoms"
	case 3:
		(*item)[index].Category = "Dress"
	case 4:
		(*item)[index].Category = "Shoes"
	case 5:
		(*item)[index].Category = "Accessories"
	default:
		fmt.Println("Invalid choice, please try again.")
		chooseCategory(item, index)
	}
}

func enterClothingName(item *wardrobe, index int) {
	var name string

	fmt.Scan(&name)

	if len(name) < 1 && name == " " {
		fmt.Println("Invalid name, please try again.")
		enterClothingName(item, index)
	} else {
		(*item)[index].Name = name
	}
}

func enterClothingColor(item *wardrobe, index int) {
	var color string

	fmt.Scan(&color)

	if len(color) < 1 && color == " " {
		fmt.Println("Invalid color, please try again.")
		enterClothingColor(item, index)
	} else {
		(*item)[index].Color = color
	}
}

func chooseWeather(item *wardrobe, index int) {
	var weatherChoice int

	fmt.Println("1. Hot")
	fmt.Println("2. Mild/Cool")
	fmt.Println("3. Cold")
	fmt.Print("Enter your choice (1-3): ")

	fmt.Scan(&weatherChoice)
	switch weatherChoice {
	case 1:
		(*item)[index].Weather = "Hot"
	case 2:
		(*item)[index].Weather = "Mild/Cool"
	case 3:
		(*item)[index].Weather = "Cold"
	default:
		fmt.Println("Invalid choice, please try again.")
		chooseWeather(item, index)
	}
}

func enterClothingFormality(item *wardrobe, index int) {
	var formality int

	fmt.Scan(&formality)

	if formality < 1 || formality > 10 {
		fmt.Println("Invalid formality, please try again.")
		enterClothingFormality(item, index)
	} else {
		(*item)[index].Formality = formality
	}
}

func enterClothingLastWorn(item *wardrobe, index int) {
	var lastWorn string

	fmt.Scan(&lastWorn)

	if len(lastWorn) != 10 {
		fmt.Println("Invalid date format, please try again.")
		enterClothingLastWorn(item, index)
	} else {
		(*item)[index].LastWorn = lastWorn
	}
}

func viewAllClothingItems(item wardrobe) {
	fmt.Println("========================View All Clothing Items========================")
	if len(item) == 0 {
		fmt.Println("No clothing items found.")
	} else {
		fmt.Printf("|| %-10s | %-10s | %-10s | %-10s | %-10s | %-10s | %-10s ||\n", "ID", "Name", "Color", "Category", "Formality", "Last Worn", "Weather")
		fmt.Println("----------------------------------------------------------------------")
		for i := 0; i < len(item); i++ {
			fmt.Printf("|| %-10d | %-10s | %-10s | %-10s | %-10d | %-10s | %-10s ||\n",
				item[i].ID, item[i].Name, item[i].Color, item[i].Category, item[i].Formality, item[i].LastWorn, item[i].Weather)
		}
	}
	fmt.Println("=======================================================================")
}

func menu() {
	fmt.Println("╔═════════════════════════════════════════╗")
	fmt.Println("║               WellDressed               ║")
	fmt.Println("║        Digital Wardrobe Assistant       ║")
	fmt.Println("╠═════════════════════════════════════════╣")
	fmt.Println("║ 1. Edit My Wardrobe                     ║")
	fmt.Println("║ 2. Sort My Wardrobe                     ║")
	fmt.Println("║ 3. Search Clothing Item                 ║")
	fmt.Println("║ 4. Recommend Me an Outfit!              ║")
	fmt.Println("║ 5. Exit                                 ║")
	fmt.Println("╚═════════════════════════════════════════╝")
}

func menuChoice(item *wardrobe, total *int) { // make it a function
	var choice int

	menu()
	fmt.Print("Choose an option (1-5): ")
	fmt.Scan(&choice)
	switch choice {
	case 1:
		editWardrobe(item, total)
	case 2:
		sortingMenu(*item, total)
	case 3:
		searchClothingItem() // not implemented yet
	case 4:
		recommendOutfit() // not implemented yet
	case 5:
		fmt.Println("Exiting WellDressed ...")
	default:
		fmt.Println("Invalid choice, please try again.")
		menuChoice(item, total)
	}
}

func editWardrobe(item *wardrobe, total *int) {
	var choice int

	fmt.Println("Edit My Wardrobe")
	fmt.Println("1. Modify Clothing Item")
	fmt.Println("2. Remove Clothing Item")
	fmt.Println("3. Back to Main Menu")
	fmt.Print("Choose an option (1-3): ")

	fmt.Scan(&choice)
	switch choice {
	case 1:
		modifyClothingItem(item, *total)
	case 2:
		removeClothingItem(item, total)
	case 3:
		menuChoice(item, total)
	default:
		fmt.Println("Invalid choice, please try again.")
		editWardrobe(item, total)
	}
}

func modifyClothingItem(item *wardrobe, total int) {
	var id int
	var choice int

	viewAllClothingItems(*item)
	fmt.Print("Enter the ID of the clothing item to modify: ")
	fmt.Scan(&id)

	for i := 0; i < len(*item); i++ {
		if (*item)[i].ID == id {
			fmt.Println("Choose new clothing category:")
			chooseCategory(item, i)

			fmt.Print("Enter new clothing name: ")
			enterClothingName(item, i)

			fmt.Print("Enter new clothing color (one word, e.g., blue, red): ")
			enterClothingColor(item, i)

			fmt.Println("Choose the new suitable weather category:")
			chooseWeather(item, i)

			fmt.Print("Enter new clothing formality (1-10): ")
			enterClothingFormality(item, i)

			fmt.Print("Enter new last worn date (YYYY-MM-DD): ")
			enterClothingLastWorn(item, i)

			viewAllClothingItems(*item)
			fmt.Println("Clothing item modified successfully!")
			editWardrobe(item, &total)
		}
	}

	fmt.Print("Item not found. Choose 1 to try again and 2 to go back: ")
	fmt.Scan(&choice)
	if choice == 1 {
		modifyClothingItem(item, total)
	} else {
		editWardrobe(item, &total)
	}
}

func removeClothingItem(item *wardrobe, total *int) {
	var id int
	var choice int

	viewAllClothingItems(*item)
	fmt.Print("Enter the ID of the clothing item to remove: ")
	fmt.Scan(&id)

	for i := 0; i < *total; i++ {
		if (*item)[i].ID == id {
			for j := i; j < *total-1; j++ {
				(*item)[j] = (*item)[j+1]
			}
			*total--

			fmt.Println("Clothing item removed successfully!")
			viewAllClothingItems(*item)
			editWardrobe(item, total)
		}
	}

	fmt.Print("Item not found. Choose 1 to try again and 2 to go back: ")
	fmt.Scan(&choice)
	if choice == 1 {
		removeClothingItem(item, total)
	} else {
		editWardrobe(item, total)
	}
}

func sortingMenu(item wardrobe, total *int) {
	var choice int

	fmt.Println("Sort My Wardrobe")
	fmt.Println("1. Sort from Most Formal to Least Formal")
	fmt.Println("2. Sort from Least Formal to Most Formal")
	fmt.Println("3. Sort by Last Worn")
	fmt.Println("4. Back to Main Menu")
	fmt.Print("Choose an option (1-4): ")

	fmt.Scan(&choice)
	switch choice {
	case 1:
		sortMostFormal(item)
	case 2:
		sortLeastFormal(item)
	case 3:
		sortByLastWorn(item)
	case 4:
		menuChoice(&item, total)
	default:
		fmt.Println("Invalid choice, please try again.")
		sortingMenu(item, total)
	}
}

func sortMostFormal(item wardrobe) {
	var sortedItem wardrobe
	sortedItem = item
	selectionSortDescendingFormal(sortedItem)
	printSorted(sortedItem)
}

func sortLeastFormal(item wardrobe) {
	var sortedItem wardrobe
	sortedItem = item
	selectionSortAscendingFormal(sortedItem)
	printSorted(sortedItem)
}

func sortByLastWorn(item wardrobe) {
	var sortedItem wardrobe
	sortedItem = item
	insertionSortByLastWorn(sortedItem)
	printSorted(sortedItem)
}

func selectionSortDescendingFormal(arr wardrobe) {
	var i, j int
	var maxIndex int
	var temp ClothingItem

	for i = 0; i < len(arr)-1; i++ {
		maxIndex = i
		for j = i + 1; j < len(arr); j++ {
			if arr[j].Formality > arr[maxIndex].Formality {
				maxIndex = j
			}
		}
		temp = arr[maxIndex]
		arr[maxIndex] = arr[i]
		arr[i] = temp
	}
}

func selectionSortAscendingFormal(arr wardrobe) {
	var i, j int
	var minIndex int
	var temp ClothingItem

	for i = 0; i < len(arr)-1; i++ {
		minIndex = i
		for j = i + 1; j < len(arr); j++ {
			if arr[j].Formality < arr[minIndex].Formality {
				minIndex = j
			}
		}
		temp = arr[minIndex]
		arr[minIndex] = arr[i]
		arr[i] = temp
	}
}

func insertionSortByLastWorn(arr wardrobe) {
	var i, j int
	var key ClothingItem

	for i = 1; i < len(arr); i++ {
		key = arr[i]
		j = i - 1
		for j >= 0 && arr[j].LastWorn < key.LastWorn {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}

func printSorted(arr wardrobe) {
	fmt.Println("========================Sorted Wardrobe========================")
	fmt.Printf("|| %-10s | %10s | %10s | %10s | %10s | %10s | %10s ||", "ID", "Name", "Color", "Category", "Formality", "Last Worn", "Weather")
	fmt.Println("---------------------------------------------------------------")
	for i := 0; i < len(arr); i++ {
		fmt.Printf("|| %-10d | %-10s | %-10s | %-10s | %-10d | %-10s | %-10s ||\n",
			arr[i].ID, arr[i].Name, arr[i].Color, arr[i].Category, arr[i].Formality, arr[i].LastWorn, arr[i].Weather)
	}
	fmt.Println("===============================================================")
}

func searchClothingItem() {
	// to be continued
	fmt.Println("Search functionality is not implemented yet.")
	menu()
}

func recommendOutfit() {
	// to be continued
	fmt.Println("Recommendation functionality is not implemented yet.")
	menu()
}
