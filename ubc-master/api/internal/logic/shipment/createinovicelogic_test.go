package shipment

import "testing"

func TestCalculateInvoiceTotals(t *testing.T) {
	tests := []struct {
		name          string
		subTotal      float64
		tax           float64
		shipping      float64
		discount      float64
		adjustments   float64
		deposit       float64
		wantGrand     float64
		wantDeposit   float64
		wantAmountDue float64
	}{
		{name: "deposit zero", subTotal: 100, tax: 5, shipping: 10, discount: 15, adjustments: 20, deposit: 0, wantGrand: 120, wantDeposit: 0, wantAmountDue: 120},
		{name: "deposit partial", subTotal: 100, tax: 5, shipping: 10, discount: 15, adjustments: 20, deposit: 30, wantGrand: 120, wantDeposit: 30, wantAmountDue: 90},
		{name: "deposit equals total", subTotal: 100, tax: 5, shipping: 10, discount: 15, adjustments: 20, deposit: 120, wantGrand: 120, wantDeposit: 120, wantAmountDue: 0},
		{name: "deposit greater than total", subTotal: 100, tax: 5, shipping: 10, discount: 15, adjustments: 20, deposit: 150, wantGrand: 120, wantDeposit: 150, wantAmountDue: 0},
		{name: "deposit negative", subTotal: 100, tax: 5, shipping: 10, discount: 15, adjustments: 20, deposit: -10, wantGrand: 120, wantDeposit: 0, wantAmountDue: 120},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := calculateInvoiceTotals(tt.subTotal, tt.tax, tt.shipping, tt.discount, tt.adjustments, tt.deposit)
			if got.GrandTotal != tt.wantGrand {
				t.Fatalf("GrandTotal = %v, want %v", got.GrandTotal, tt.wantGrand)
			}
			if got.Deposit != tt.wantDeposit {
				t.Fatalf("Deposit = %v, want %v", got.Deposit, tt.wantDeposit)
			}
			if got.AmountDue != tt.wantAmountDue {
				t.Fatalf("AmountDue = %v, want %v", got.AmountDue, tt.wantAmountDue)
			}
		})
	}
}
