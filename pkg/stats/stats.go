package stats

import (
	"github.com/AmirjonQodirov/bank/pkg/types"
)

func Avg(payments []types.Payment) types.Money {
	var avg types.Money
	count := 0
	for _, p := range payments {
		avg += p.Amount
		count++
	}
	return avg / types.Money(count)
}



func TotalInCategory(payments []types.Payment, category types.Category) types.Money{
	var sum types.Money
	for _, p := range payments {
		if(p.Category == category){
			sum += p.Amount
		}
	}
	return sum
}
