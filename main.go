package main

import "fmt"

type ClothingItem struct {
	ID        int
	Name      string
	Color     string
	Category  string
	Formality int
	LastWorn  string // e.g., "2024-05-01"
	Weather   string // e.g., "Rainy", "Hot"
}

type wardrobe []ClothingItem

func main() {
	var item wardrobe
	var choice int

	menu()
	fmt.Scan(&choice)
	switch choice {
	case 1:
		editWardrobe()
	case 2:
		sortingMenu(item)
	case 3:
		searchClothingItem()
	case 4:
		recommendOutfit()
	case 5:
		fmt.Println("Exiting WellDressed ...")
	default:
		fmt.Println("Invalid choice, please try again.")
		menu()
	}
}

func menu() {
	fmt.Println("=======WellDressed: Digital Wardrobe=======")
	fmt.Println("1. Edit My Wardrobe")
	fmt.Println("2. Sort My Wardrobe")
	fmt.Println("3. Search Clothing Item")
	fmt.Println("4. Reccommend Me an Outfit!")
	fmt.Println("5. Exit")
	fmt.Println("===========================================")
	fmt.Print("Choose an option (1-5): ")
}

func editWardrobe() {
	var item wardrobe
	var choice int

	fmt.Println("=======Edit My Wardrobe=======")
	fmt.Println("1. Add Clothing Item")
	fmt.Println("2. Modify Clothing Item")
	fmt.Println("3. Remove Clothing Item")
	fmt.Println("4. View All Clothing Items")
	fmt.Println("5. Back to Main Menu")
	fmt.Println("===============================")
	fmt.Print("Choose an option (1-5): ")

	fmt.Scan(&choice)
	switch choice {
	case 1:
		addClothingItem(&item)
	case 2:
		modifyClothingItem(&item)
	case 3:
		removeClothingItem(&item)
	case 4:
		viewAllClothingItems(item) // not made
	case 5:
		menu()
	default:
		fmt.Println("Invalid choice, please try again.")
		editWardrobe()
	}
}

func menuCategory(item *wardrobe, index int) {
	var categoryChoice int

	fmt.Println("1. Top")
	fmt.Println("2. Bottoms")
	fmt.Println("3. Dress")
	fmt.Println("4. Shoes")
	fmt.Println("5. Accessories")
	fmt.Println("6. Others")
	fmt.Print("Enter your choice (1-6): ")

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
	case 6:
		(*item)[index].Category = "Others"
	default:
		fmt.Println("Invalid choice, please try again.")
		menuCategory(item, index)
	}
}

func menuWeather(item *wardrobe, index int) {
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
		menuWeather(item, index)
	}
}

// belum ada error handling
func addClothingItem(item *wardrobe) {
	var total, NextID int
	NextID = 1

	fmt.Print("Enter the number of clothing items to add: ")
	fmt.Scan(&total)
	for i := 0; i < total; i++ {
		(*item)[i].ID = NextID

		fmt.Println("Choose clothing category:")
		menuCategory(item, i)

		fmt.Print("Enter clothing name: ")
		fmt.Scan(&(*item)[i].Name)

		fmt.Print("Enter clothing color (one word, e.g., blue, red): ")
		fmt.Scan((*item)[i].Color)

		fmt.Println("Choose the suitable weather category:")
		menuWeather(item, i)

		fmt.Print("Enter clothing formality (1-10): ")
		fmt.Scan((*item)[i].Formality)

		fmt.Print("Enter last worn date (YYYY-MM-DD): ")
		fmt.Scan((*item)[i].LastWorn)

		NextID++
	}
	fmt.Println("Clothing item added successfully!")
	editWardrobe()
}

func modifyClothingItem(item *wardrobe) {
	var id int
	var choice int

	fmt.Print("Enter the ID of the clothing item to modify: ")
	fmt.Scan(&id)

	for i := 0; i < len(*item); i++ {
		if (*item)[i].ID == id {
			fmt.Println("Choose new clothing category:")
			menuCategory(item, i)

			fmt.Print("Enter new clothing name: ")
			fmt.Scan(&(*item)[i].Name)

			fmt.Print("Enter new clothing color (one word, e.g., blue, red): ")
			fmt.Scan(&(*item)[i].Color)

			fmt.Println("Choose the new suitable weather category:")
			menuWeather(item, i)

			fmt.Print("Enter new clothing formality (1-10): ")
			fmt.Scan(&(*item)[i].Formality)

			fmt.Print("Enter new last worn date (YYYY-MM-DD): ")
			fmt.Scan(&(*item)[i].LastWorn)

			fmt.Println("Clothing item modified successfully!")
			return
		}
	}
	fmt.Print("Item not found. Choose 1 to try again and 2 to go back: ")
	fmt.Scan(&choice)
	if choice == 1 {
		modifyClothingItem(item)
	} else {
		editWardrobe()
	}
}

func removeClothingItem(item *wardrobe) {
	var id, count int
	var choice int
	count = len(*item)

	fmt.Print("Enter the ID of the clothing item to remove: ")
	fmt.Scan(&id)

	for i := 0; i < count; i++ {
		if (*item)[i].ID == id {
			for j := i; j < count-1; j++ {
				(*item)[j] = (*item)[j+1]
			}
			*item = (*item)[:count-1]
			fmt.Println("Clothing item removed successfully!")
			return
		}
	}
	fmt.Print("Item not found. Choose 1 to try again and 2 to go back: ")
	fmt.Scan(&choice)
	if choice == 1 {
		removeClothingItem(item)
	} else {
		editWardrobe()
	}
}

func viewAllClothingItems(item wardrobe) {
	fmt.Println("========================View All Clothing Items========================")
	if len(item) == 0 {
		fmt.Println("No clothing items found.")
	} else {
		fmt.Printf("|| %-10s | %-10s | %-10s | %-10s | %-10s | %-10s | %-10s ||\n", "ID", "Name", "Color", "Category", "Formality", "Last Worn", "Weather")
		fmt.Println("---------------------------------------------------------------------")
		for i := 0; i < len(item); i++ {
			fmt.Printf("|| %-10d | %-10s | %-10s | %-10s | %-10d | %-10s | %-10s ||\n",
				item[i].ID, item[i].Name, item[i].Color, item[i].Category, item[i].Formality, item[i].LastWorn, item[i].Weather)
		}
	}
	fmt.Println("=======================================================================")
	editWardrobe()
}

func sortingMenu(item wardrobe) {
	var choice int

	fmt.Println("==============Sort My Wardrobe=============")
	fmt.Println("1. Sort from Most Formal to Least Formal")
	fmt.Println("2. Sort from Least Formal to Most Formal")
	fmt.Println("3. Sort by Last Worn")
	fmt.Println("4. Back")
	fmt.Println("===========================================")
	fmt.Print("Choose an option (1-4): ")

	fmt.Scan(&choice)
	switch choice {
	case 1:
		sortMostFormal(item) // not made
	case 2:
		sortLeastFormal(item) // not made
	case 3:
		sortByLastWorn(item) // not made
	case 4:
		menu()
	default:
		fmt.Println("Invalid choice, please try again.")
		sortingMenu(item)
	}
}

func sortMostFormal(item wardrobe) {
	var arr wardrobe
	arr = item
	selectionSortDescendingFormal(arr)
	printSorted(arr)
}

func sortLeastFormal(item wardrobe) {
	var arr wardrobe
	arr = item
	selectionSortAscendingFormal(arr)
	printSorted(arr)
}

func sortByLastWorn(item wardrobe) {
	var arr wardrobe
	arr = item
	insertionSortByLastWorn(arr)
	printSorted(arr)
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
