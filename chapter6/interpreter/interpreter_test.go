package interpreter

import (
	"testing"
)

func TestCalculate_3_4_sum_2_sub(t *testing.T) {
	//􏰖􏰖􏰖􏱀Operation: (3 + 4) - 2
	tempOperation := "3 4 sum 2 sub"
	expected := 5
	res, err := Calculate(tempOperation)
	if err != nil {
		t.Error(err)
	}

	if res != expected {
		t.Errorf("Expected result not found: %d != %d\n", expected, res)
	}
}

func TestCalculate_5_3_sub_8_mul_4_sum_5_div(t *testing.T) {
	//􏰖􏰖􏰖􏱀Operation: (((5 - 3) * 8) + 4) / 5
	tempOperation := "5 3 sub 8 mul 4 sum 5 div"

	expected := 4
	res, err := Calculate(tempOperation)
	if err != nil {
		t.Error(err)
	}

	if res != expected {
		t.Errorf("Expected result not found: %d != %d\n", expected, res)
	}
}
