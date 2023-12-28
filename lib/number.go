package numbers

import "strings"

func ToFrench(numbersSlice []int) []string {
	var frenchNumbers []string

	for i := 0; i < len(numbersSlice); i++ {
		currentNumber := numbersSlice[i]
		var frenchNumber string

		if currentNumber == 0 {
			frenchNumber = tens[0]
		} else if currentNumber < 100 {
			frenchNumber = handleTens(currentNumber)
		} else if currentNumber < 1000 {
			frenchNumber = handleHundreds(currentNumber)
		} else {
			frenchNumber = handleThousands(currentNumber)
		}

		frenchNumber = strings.Trim(frenchNumber, "-")
		frenchNumbers = append(frenchNumbers, frenchNumber)
	}

	return frenchNumbers
}

func handleTens(currentNumber int) string {
	var frenchNumber string

	if currentNumber == 0 {
		return ""
	}

	if currentNumber <= 16 {
		frenchNumber = units[currentNumber]
	} else if currentNumber <= 99 {
		unit := currentNumber % 10
		ten := currentNumber / 10

		var frenchUnit string
		frenchTens := tens[ten]
		if isSoixtanteDixOrQuartreVingtDix(currentNumber) {
			if unit <= 6 {
				// soixante-et-onze
				// use previous tens
				frenchTens = tens[ten-1]

				// soixante-et-onze
				var separator string
				if currentNumber == 71 {
					separator = "-et-"
				} else {
					separator = "-"
				}

				// use tens units
				frenchUnit = separator + units[unit+10]
			} else {
				frenchUnit = "-" + units[unit]
			}
		} else {
			if unit == 1 && currentNumber != 81 {
				frenchUnit = "-" + andOne
			} else {
				frenchUnit = "-" + units[unit]
			}
		}

		if unit == 0 {
			// plural
			if ten == 8 {
				frenchUnit = "s"
			} else {
				frenchUnit = ""
			}
		}

		frenchNumber = frenchTens + frenchUnit
	}

	return frenchNumber
}

func handleHundreds(currentNumber int) string {
	var frenchNumber string

	ten := currentNumber % 100
	hundred := currentNumber / 100

	frenchHundred := handleTens(hundred) + "-" + hundreds + "-"
	// cent-un
	if currentNumber < 200 {
		frenchHundred = hundreds + "-"
	}
	frenchNumber = frenchHundred + handleTens(ten)

	if currentNumber == 100 {
		frenchNumber = hundreds
	}

	return frenchNumber
}

func handleThousands(currentNumber int) string {
	var frenchNumber string

	hundred := currentNumber % 1000
	thousand := currentNumber / 1000

	frenchThousand := handleTens(thousand) + "-" + thousands + "-"
	// mille-un
	if currentNumber < 2000 {
		frenchThousand = thousands + "-"
	}

	frenchHundred := handleHundreds(hundred)
	if hundred < 100 {
		frenchHundred = handleTens(hundred)
	}
	frenchNumber = frenchThousand + frenchHundred

	if currentNumber == 1000 {
		frenchNumber = thousands
	}

	return frenchNumber
}

func isSoixtanteDixOrQuartreVingtDix(currentNumber int) bool {
	return (currentNumber > 70 && currentNumber < 80) || (currentNumber > 90 && currentNumber < 100)
}
