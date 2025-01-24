package main

import (
	"fmt"
	"gopass/app"
	"gopass/config"
	"log"
)

func main() {

	cfg := config.DefaultConfig()
	
	for {

		fmt.Println("\nChoose an option:")
		fmt.Println("1. Start the application")
		fmt.Println("2. Change Preference")
		fmt.Println("3. Generate again")
		fmt.Println("4. Exit the application")

		var choice int
		fmt.Print("Enter your choice: ")
		_, err := fmt.Scan(&choice)
		if err != nil {
			fmt.Println("Invalid input, please try again.")
			continue
		}

		switch choice {
		case 1:

			password, err := app.GeneratePassword(cfg)
			if err != nil {
				log.Fatalf("Error generating password: %v", err)
			}
			fmt.Printf("Generated Password: %s\n", password)
		
		case 2:

			fmt.Print("Enter the desired password length: ")
			var length int
			_, err := fmt.Scan(&length)
			if err != nil || length <= 0 {
				fmt.Println("Invalid input. Please enter a positive integer.")
				continue
			}
			cfg.Length = length
			
			fmt.Print("Include lowercase letters? (yes/no): ")
			cfg.IncludeLower = getYesNoInput()
			
			fmt.Print("Include uppercase letters? (yes/no): ")
			cfg.IncludeUpper = getYesNoInput()
			
			fmt.Print("Include numbers? (yes/no): ")
			cfg.IncludeNumbers = getYesNoInput()
			
			fmt.Print("Include symbols? (yes/no): ")
			cfg.IncludeSymbols = getYesNoInput()
			
			fmt.Println("Preferences updated successfully!")
		
		case 3:

			password, err := app.GeneratePassword(cfg)
			if err != nil {
				log.Fatalf("Error generating password: %v", err)
			}
			fmt.Printf("Generated Password: %s\n", password)
		
		case 4:
			
			fmt.Println("Exiting the application. Goodbye!")
			return
		
		default:
			fmt.Println("Invalid choice. Please select a valid option.")
		}
	}
}


func getYesNoInput() bool {
	var input string
	for {
		fmt.Scan(&input)
		switch input {
		case "yes", "y":
			return true
		case "no", "n":
			return false
		default:
			fmt.Print("Invalid input. Please enter 'yes' or 'no': ")
		}
	}
}
