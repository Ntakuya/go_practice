package evenorodd

func EvenOrOdd(number int) string {
	if result := number % 2; result == 0 {
		return "Even"
	}
	return "Odd"
}

// func EvenOrOdd(number int) string {
// 	return []string{"Even", "Odd"}[number & 1]
// }

// func EvenOrOdd(number int) string {
// 	return []string{"Even", "Odd"}[int(math.Abs(float64(number))) % 2]
// }

// func EvenOrOdd(number int) string {
// 	dict := map[int]string{0:"Even", 1:"Odd"}
// 	return dict[int(math.Abs(float64(number))) % 2]
// }

// func EvenOrOdd(number int) string {
//     dict := map[int]string{0:"Even", 1:"Odd"}
// 	return dict[number & 1]
// }