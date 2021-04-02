package stats

import (
	"github.com/AmirjonQodirov/bank/v2/pkg/types"
)

func Avg(payments []types.Payment) types.Money {
	var avg types.Money
	count := 0
	for _, p := range payments {
		if p.Status != types.StatusFail {
			avg += p.Amount
			count++
		}
	}
	return avg / types.Money(count)
}

func TotalInCategory(payments []types.Payment, category types.Category) types.Money {
	var sum types.Money
	for _, p := range payments {
		if p.Category == category && p.Status != types.StatusFail {
			sum += p.Amount
		}
	}
	return sum
}

func CategoriesAvg(payments []types.Payment) map[types.Category]types.Money {
	categories := map[types.Category]types.Money{}
	categories_count := map[types.Category]types.Money{}
	for _, payment := range payments {
		categories[payment.Category] += payment.Amount
		categories_count[payment.Category]++		
	}
	for key:= range categories{
		categories[key]=categories[key]/categories_count[key]
	}		
	return categories
}
