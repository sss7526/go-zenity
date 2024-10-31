package main

import (
    "fmt"
    "github.com/ncruces/zenity"
)

func main() {
    // err := zenity.Question(
    //     "Do you want to continue?",  // Dialog message
    //     zenity.OKLabel("Yes"),        // Label the OK button as "Yes"
    //     zenity.CancelLabel("Cancel")) // Label the Cancel button as "Cancel"

    // if err == zenity.ErrCanceled {
    //     fmt.Println("User canceled")
    // } else if err == nil {
    //     fmt.Println("User confirmed.")
    // }

	// err := zenity.Question(
    //     "Do you want to proceed?",
    //     zenity.OKLabel("Yes"),
    //     zenity.CancelLabel("No"),
    //     zenity.ExtraButton("Don't show again"), // Add the extra checkbox button
    // )

    // if err == zenity.ErrCanceled {
    //     fmt.Println("User cancelled")
    // } else if err == nil {
    //     if checked {
    //         fmt.Println("User checked 'Don't show again'")
    //     }
    //     fmt.Printf("User confirmed with: %v\n", result)
    // }
	// filePath, err := zenity.SelectFile()
	// 	if err != nil {
	// 		if err == zenity.ErrCanceled {
	// 			fmt.Println("File selection was canceled")
	// 		} else {
	// 			fmt.Println("Error selecting file:", err)
	// 		}
	// 	} else {
	// 		fmt.Printf("File selected: %s\n", filePath)
	// 	}

	// filePath, err := zenity.SelectFileSave()

    // if err != nil {
    //     if err == zenity.ErrCanceled {
    //         fmt.Println("Save canceled")
    //     } else {
    //         fmt.Println("Error in save dialog:", err)
    //     }
    // } else {
    //     fmt.Printf("File selected to save: %s\n", filePath)
    // }


	// Present the user with a directory selection dialog
    // dirPath, err := zenity.SelectFile(zenity.Directory())
    // if err != nil {
    //     if err == zenity.ErrCanceled {
    //         fmt.Println("Directory selection canceled")
    //     } else {
    //         fmt.Println("Error selecting directory:", err)
    //     }
    // } else {
    //     fmt.Printf("Directory selected: %s\n", dirPath)
    // }


	// Display an input dialog to collect input from the user
    // input, err := zenity.Entry("Please enter your name:")

    // if err == zenity.ErrCanceled {
    //     fmt.Println("Input canceled")
    // } else if err == nil {
    //     fmt.Printf("User entered: %s\n", input)
    // } else {
    //     fmt.Println("Error:", err)
    // }

	// Dialog with "Don't show again" as an extra option
    // err := zenity.Question(
    //     "Do you want to continue?",
    //     zenity.ExtraButton("Don't show again"),
    //     zenity.OKLabel("Yes"),
    //     zenity.CancelLabel("No"))

    // if err == zenity.ErrExtraButton {
    //     fmt.Println("User selected 'Don't show again'.")
    // } else if err == zenity.ErrCanceled {
    //     fmt.Println("User selected 'No' or closed the dialog.")
    // } else if err == nil {
    //     fmt.Println("User selected 'Yes'.")
    // }

	// Show a question dialog with "Yes", "No", and an extra "Don't show again" option
    err := zenity.Question(
        "Do you want to proceed?",
        zenity.OKLabel("Yes"),
        zenity.CancelLabel("No"),
        zenity.ExtraButton("Don't show again"))

    // Variable to track user's choice on "Don't show again"
    dontShowAgain := false

    if err == zenity.ErrExtraButton {
        // User clicked "Don't show again", process accordingly
        fmt.Println("User clicked 'Don't show again'.")
        dontShowAgain = true
    } else if err == zenity.ErrCanceled {
        // User clicked "No" or closed the dialog
        fmt.Println("User clicked 'No' or canceled.")
    } else if err == nil {
        // User clicked "Yes"
        fmt.Println("User clicked 'Yes'.")
    }

    // Process 'Don't show again' flag based on user's choice
    if dontShowAgain {
        // Save "don't show again" state in your program configuration.
        fmt.Println("Setting 'Don't show again' flag.")
    }

    // Perform your actions based on the results
    if err == nil {
        fmt.Println("Proceeding with the 'Yes' action.")
    } else if err == zenity.ErrCanceled {
        fmt.Println("Stopping due to 'No' selection.")
    }

}