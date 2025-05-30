# WellDressed: Digital Wardrobe Assistant
Welcome to WellDressed, a command-line application built in Go that helps you manage your clothing wardrobe and get outfit recommendations!

## **FEATURES**
WellDressed provides a suite of features to keep your wardrobe organized:
- Add Clothing Items: Easily add new tops, bottoms, dresses, shoes, or accessories to your collection.
- Modify Clothing Items: Update details like name, color, category, formality, weather suitability, and last worn date for any item.
- Remove Clothing Items: Declutter your wardrobe by deleting items you no longer own.
- Sort Functionality: Sort by formality and last worn; view your items sorted from most to least formal, vice versa and items you've worn recently to help rotate your outfit.
- Search Functionality: Search by color and category; guickly find items of a specific color or category (e.g., all your "Tops" or "Shoes").
- Outfit Recommendations: Get personalized outfit suggestions based on your preferred weather conditions, formality level, and whether you prefer a top-and-bottom combination or a dress.

## **HOW TO USE**
To run WellDressed, you'll need Go installed on your system.

1. Save the code: Save the provided Go code as main.go in a directory of your choice.
2. Navigate to the directory: Open your terminal or command prompt and navigate to the directory where you saved main.go.
3. Run the application: Execute the program using the go run command: _go run main.go_
4. Follow the prompts: The application will greet you and guide you through setting up your wardrobe. You can then select options from the main menu to manage your clothes or get recommendations.

## **ALGORITHMS USED**
- Searching: Used sequential search for colors and binary search for category.
- Sorting: Used selection sort for sort by formality and insertion sort for sort by last worn (string comparison for date).

## **STRUCTURE**
- The ClothingItem struct defines the properties of each clothing item in your wardrobe.
- Your wardrobe is stored as a fixed-size array (wardrobe [NMAX]ClothingItem), allowing for up to 99 clothing items.
- Used 33 subprograms (function and procedures) as follows:
> 1. main(): The entry point of your program.
> 2. welcome(): Displays a welcome message and prompts for user name.
> 3. setupWardrobe(): Initializes the wardrobe, now including initial data.
> 4. addClothingItem(): Handles adding new clothing items to the wardrobe.
> 5. enterClothingName(): Prompts for and validates clothing item name.
> 6. chooseClothingCategory(): Prompts for and validates clothing category choice.
> 7. category(): Converts a numeric category choice to its string representation.
> 8. chooseClothingWeather(): Prompts for and validates weather suitability choice.
> 9. weather(): Converts a numeric weather choice to its string representation.
> 10. chooseClothingFormality(): Prompts for and validates formality choice.
> 11. formality(): Converts a numeric formality choice to its string representation.
> 12. enterClothingColor(): Prompts for and validates clothing color.
> 13. enterClothingLastWorn(): Prompts for and validates last worn date.
> 14. viewAllClothingItems(): Displays all items currently in the wardrobe.
> 15. menu(): Displays the main menu options.
> 16. menuChoiceValidity(): Validates main menu input.
> 17. editWardrobe(): Presents options for modifying, removing, or adding items.
> 18. modifyClothingItem(): Guides the user through modifying an existing item.
> 19. modifyByField(): Allows the user to choose and modify a specific field of an item.
> 20. removeClothingItem(): Handles the removal of an item from the wardrobe.
> 21. sorting(): Presents sorting options and calls respective sorting functions.
> 23. selectionSortDescendingFormal(): Sorts items by formality in descending order.
> 24. selectionSortAscendingFormal(): Sorts items by formality in ascending order.
> 26. insertionSortByLastWorn(): Sorts items by their last worn date (most recent first).
> 27. printSorted(): Prints the sorted wardrobe items.
> 28. searchClothingItem(): Presents search options (by color or category).
> 29. sequentialSearchColor(): Performs a linear search for items by color.
> 30. selectionSortAscendingCategory(): Sorts items by category in ascending order (used before binary search).
> 31. binarySearchCategory(): Performs a binary search for items by category.
> 32. recommendOutfit(): Orchestrates the outfit recommendation process.
> 33. userPreferences(): Gathers user preferences for outfit recommendation.
> 34. findRecommendation(): Finds and displays the recommended outfit.
> 35. sequentialSearchRecommendation(): Finds a suitable item for recommendation based on weather and formality.

## **Control Flow**
- **Loops**: These are extensively used for repetition. You can find loops in places like:
  - Iterating through the wardrobe array (e.g., in viewAllClothingItems, sequentialSearchColor, selectionSortDescendingFormal, insertionSortByLastWorn).
  - Handling user input validation (e.g., menuChoiceValidity, addClothingItem, enterClothingName, chooseClothingCategory, chooseClothingWeather, chooseClothingFormality, enterClothingColor, enterClothingLastWorn).
  - **Repeat Until**: The main application loop in main() to keep the program running until the user chooses to exit.
- **If Statements**: if and else if statements are crucial for conditional logic, allowing your program to make decisions. They are used for:
  - Checking if a clothing item is found during search or modification operations (found boolean variable).
  - Validating user input ranges and formats.
  - Determining the correct category, weather, or formality string based on numeric input.
  - Providing specific outfit recommendations based on user preferences (top-bottom vs. dress combo).
  - Handling cases where no matching item is found during recommendations.
- **Switch Case**: These provide a cleaner way to handle multiple conditional branches, especially when dealing with menu choices or category/formality/weather conversions.
  - The main menu's action dispatch (main()).
  - Editing wardrobe choices (editWardrobe).
  - Modifying specific fields of an item (modifyByField).
  - Converting numeric category, weather, and formality choices to their string representations (e.g., in category, weather, formality functions).
  - Choosing between search options (searchClothingItem).
