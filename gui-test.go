package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

var (
	nameEntry       *widget.Entry
	positionEntry   *widget.Entry
	departmentEntry *widget.Entry
	packageACheck   *widget.Check
	packageBCheck   *widget.Check
	packageCCheck   *widget.Check
	resultLabel     *widget.Label
	prices          = map[string]int{"A": 100, "B": 200, "C": 300}
	myWindow        fyne.Window
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Form Request Hardware")

	// Widgets
	nameEntry = widget.NewEntry()
	positionEntry = widget.NewEntry()
	departmentEntry = widget.NewEntry()

	packageACheck = widget.NewCheck("Package A ($100)", nil)
	packageBCheck = widget.NewCheck("Package B ($200)", nil)
	packageCCheck = widget.NewCheck("Package C ($300)", nil)

	submitBtn := widget.NewButton("Submit", handleSubmit)

	resultLabel = widget.NewLabel("Results will appear here...")
	resultLabel.Wrapping = fyne.TextWrapWord
	// resultLabel.Wrapping = fyne.TextWrapWord  // Or any other mode you prefer

	// Layouts
	formContainer := container.NewGridWithColumns(2,
		widget.NewLabel("Name:"), nameEntry,
		widget.NewLabel("Position:"), positionEntry,
		widget.NewLabel("Department:"), departmentEntry,
		widget.NewLabel("Which software do you require?"), widget.NewLabel(""),
		packageACheck, packageBCheck, packageCCheck, widget.NewLabel(""),
	)

	mainContainer := container.NewVBox(
		formContainer,
		widget.NewSeparator(),
		container.NewStack(submitBtn),
		resultLabel,
	)

	myWindow.SetContent(mainContainer)
	myWindow.ShowAndRun()
}

func handleSubmit() {
	totalPrice := 0
	if packageACheck.Checked {
		totalPrice += prices["A"]
	}
	if packageBCheck.Checked {
		totalPrice += prices["B"]
	}
	if packageCCheck.Checked {
		totalPrice += prices["C"]
	}

	results := fmt.Sprintf(
		"Name: %s\nPosition: %s\nDepartment: %s\nTotal Price: $%d",
		nameEntry.Text, positionEntry.Text, departmentEntry.Text, totalPrice,
	)

	resultLabel.SetText(results)

	dialog.ShowInformation("Price", fmt.Sprintf("Total Price: $%d", totalPrice), myWindow)
}
