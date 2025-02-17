package main

import (
	"errors"
	"log"

	"github.com/charmbracelet/huh"
)

// To start, we'll import the library and define a few variables where'll we
// answers.
var (
	burger string
	name   string
)

// `huh?` separates forms into groups (groups can be seen as pages). Groups are
// made of fields (e.g. Select, Input, Text). We will set three groups (pages)
// for customers to fill out.
func main() {
	// NewForm returns a form with the given groups and default themes and
	// keybindings.
	form := huh.NewForm(
		// Ask the user for a base burger and toppings.
		huh.NewGroup(
			// A select field is a field that allows the user to select from a
			// list of options.
			huh.NewSelect[string]().
				Options(
					huh.NewOption("Charmburger Classic", "classic"),
					huh.NewOption("ChickWich", "chihcwich"),
					huh.NewOption("FishBurger", "fishburger"),
					huh.NewOption("CharmPossible Burger™", "charmpossible"),
				).
				Value(&burger),
		),

		// Gather some final details about the order.
		huh.NewGroup(
			// The input field is a field that allows the user to enter text.
			huh.NewInput().
				Title("What's your name?").
				Value(&name).
				Validate(func(str string) error {
					if str == "Frank" {
						return errors.New("sorry, we dont server customers named Frank")
					}
					return nil
				}),
		),
	)

	err := form.Run()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Burger: %s  -  Name: %s \n", burger, name)
}
