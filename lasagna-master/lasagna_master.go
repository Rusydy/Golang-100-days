package lasagna

// TODO: define the 'PreparationTime()' function

func PreparationTime(layers []string, minutes int) int {
	if minutes == 0 {
		return len(layers) * 2
	}
	return len(layers) * minutes
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {
	sauce := Counter(layers, "sauce")
	noodles := Counter(layers, "noodles")

	return 50 * noodles, 0.2 * float64(sauce)
}

func Counter(list []string, str string) int {
	r := 0
	for _, v := range list {
		if v == str {
			r++
		}
	}
	return r
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList []string, myList []string) []string {
	// for _, v := range friendsList {
	// 	if Counter(friendsList, v) == 0 {
	// 		myList[len(myList)-1] = v
	// 	}
	// }

	// this is not a good solution, probably a bug from exercism
	myList[len(myList)-1] = friendsList[len(friendsList)-1]
	return myList
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, portions int) (qtyPortions []float64) {
	for _, v := range quantities {
		qtyPortions = append(qtyPortions, v*float64(portions)/2)
	}
	return
}
