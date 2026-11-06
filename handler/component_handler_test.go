package handler

import (
	"testing"
)

func TestValidateComponentStock(t *testing.T) {
	tests := []struct {
		name          string
		stock         int
		quantityPerPC int
		totalPC       int
		expected      bool
	}{
		{"Stok cukup", 10, 2, 5, true},
		{"Stok tidak cukup", 10, 2, 6, false},
		{"Stok negatif", -1, 2, 5, false},
		{"Jumlah PC nol", 10, 2, 0, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ValidateComponentStock(
				tt.stock,
				tt.quantityPerPC,
				tt.totalPC,
			)

			if result != tt.expected {
				t.Errorf(
					"Dapat %v, seharusnya %v",
					result,
					tt.expected,
				)
			}
		})
	}
}
