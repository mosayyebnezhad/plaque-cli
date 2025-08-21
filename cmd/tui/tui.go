package main

import (
	"errors"
	"fmt"
	"plaque/engine"

	"strconv"

	"github.com/manifoldco/promptui"
)

func main() {

	validate := func(input string, secT int) error {
		if len(input) != secT {
			return errors.New("invalid len")
		}
		_, err := strconv.Atoi(input)
		if err != nil {
			return errors.New("invalid number")
		}
		return nil
	}

	validateNum2 := func(s string) error { return validate(s, 2) }

	// validateNum3 := func(s string) error { return validate(s, 3) }

	validatePersianChar := func(s string) error {

		runes := []rune(s)
		if len(runes) < 1 || len(runes) > 1 {
			return errors.New("must be  1 characters")
		}

		for _, r := range runes {
			if !isPersianRune(r) {
				return errors.New("contains non-Persian characters")
			}
		}

		return nil
	}

	prompth := promptui.Prompt{
		Label:    "enter",
		Validate: validatePersianChar,
	}

	prompt3 := promptui.Prompt{
		Label:    "enter",
		Validate: validateNum2,
	}

	var prh, pr3 string

	fmt.Printf("00 \033[32mA\033[0m 000 00 \n")
	if result, err := prompth.Run(); err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	} else {
		prh = result
	}
	fmt.Printf("00 %s 000 \033[32m00\033[0m \n", prh)
	if result, err := prompt3.Run(); err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	} else {
		pr3 = result
	}

	// prh is equal to y
	// pr3 is equal to w
	//
	reciveData := engine.Engine(prh, pr3)

	fmt.Printf("province is :%s \ncity is :%s \n", reciveData.Province, reciveData.City)

}

func isPersianRune(r rune) bool {
	// فارسی از U+0600 تا U+06FF (برای عربی و فارسی)، و U+FB50 تا U+FDFF و U+FE70 تا U+FEFF هم ممکنه کاربرد داشته باشن.
	return (r >= 0x0600 && r <= 0x06FF) ||
		(r >= 0xFB50 && r <= 0xFDFF) ||
		(r >= 0xFE70 && r <= 0xFEFF)
}
