package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, average int) int {
    if average == 0 {
        average = 2
    }
    return len(layers) * average
}

// Quantities calculates the amount of noodles and sauce needed.
func Quantities(layers []string) (noodles int, sauce float64) {
    for _, layer := range layers {
        switch layer {
        case "noodles":
            noodles += 50
        case "sauce":
            sauce += 0.2
        }
    }
    return
}

// AddSecretIngredient adds the secret ingredient from a friend's list to my list.
func AddSecretIngredient(friendsList, myList []string) []string {
    if len(friendsList) > 0 {
        myList[len(myList)-1] = friendsList[len(friendsList)-1]
    }
    return myList
}
// ScaleRecipe scales the recipe for a given number of portions.
func ScaleRecipe(quantities []float64, portions int) []float64 {
    scaled := make([]float64, len(quantities))
    factor := float64(portions) / 2.0

    for i, qty := range quantities {
        scaled[i] = qty * factor
    }
    return scaled
}
// TODO: define the 'Quantities()' function

// TODO: define the 'AddSecretIngredient()' function

// TODO: define the 'ScaleRecipe()' function

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
