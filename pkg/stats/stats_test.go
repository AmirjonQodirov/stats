package stats

import(
	"testing"
	"reflect"
	"github.com/AmirjonQodirov/bank/v2/pkg/types"
)

func TestCategoriesAvg_nil(t *testing.T) {
	var payments []types.Payment
	result:= CategoriesAvg(payments)
	if len(result)!=0{
		t.Error("result len != 0")
	}
}
	
func TestCategoriesAvg_empty(t *testing.T) {
	payments := []types.Payment{}
	result:= CategoriesAvg(payments)
	if len(result)!=0{
		t.Error("result len != 0")
	}
}

func TestCategoriesAvg(t *testing.T) {
	payments := []types.Payment{
		{
			ID: 1,
			Category: "mobile",
			Amount: 50000,
		},
		{
			ID: 2,
			Category: "mobile",
			Amount: 50000,
		},
		{
			ID: 3,
			Category: "card",
			Amount: 10000,
		},
		{
			ID: 4,
			Category: "card",
			Amount: 10000,
		},
		{
			ID: 5,
			Category: "card",
			Amount: 10000,
		},
	}
	expected := map[types.Category]types.Money{
		"card":10000,
		"mobile":50000,
	}
	result:=CategoriesAvg(payments)
	if !reflect.DeepEqual(expected,result){
		t.Errorf("invalid result, expected: %v, actual: %v",expected,result)
	}
}