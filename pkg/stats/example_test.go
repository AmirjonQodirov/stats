package stats

import (
	"fmt"

	"github.com/AmirjonQodirov/bank/pkg/types"
)

func ExampleAvg() {
	payments := []types.Payment{
		{
			Amount:   10000,
			ID:       2,
			Category: "mobile",
		},
		{
			Amount:   11000,
			ID:       22,
			Category: "mobile",
		},
		{
			Amount:   10000,
			ID:       222,
			Category: "mobile",
		},
		{
			Amount:   11000,
			ID:       2222,
			Category: "mobile",
		}}
	avg := Avg(payments)
	fmt.Println(avg)
	//Output: 10500
}

func ExampleTotalInCategory() {
	payments := []types.Payment{
		{
			Amount:   10000,
			ID:       2,
			Category: "mobile",
		},
		{
			Amount:   11000,
			ID:       22,
			Category: "mobile1",
		},
		{
			Amount:   10000,
			ID:       222,
			Category: "mobile1",
		},
		{
			Amount:   11000,
			ID:       2222,
			Category: "mobile",
		}}
	sum := TotalInCategory(payments, "mobile")
	fmt.Println(sum)
	//Output: 21000
}
