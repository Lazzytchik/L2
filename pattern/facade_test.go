package pattern

import "testing"

func TestInit(t *testing.T) {
	facade := newWalletFacade("test")
	if facade != nil {
		return
	}

	t.Fatalf("Facade has not been initialized.")
}

func TestMoneyAdd(t *testing.T) {
	facade := newWalletFacade("test")

	var value float64 = 67

	if err := facade.AddMoney("test", value); err != nil {
		t.Fatalf("Error on money Add: %s", err)
	}

	if facade.wallet.balance != value {
		t.Fatalf("Wallet balance is wrong: want: %f, got: %f", value, facade.wallet.balance)
	}
}

func TestMoneyDeduct(t *testing.T) {
	facade := newWalletFacade("test")

	var value float64 = 67

	if err := facade.AddMoney("test", value); err != nil {
		t.Fatalf("Error on money Add: %s", err)
	}

	if err := facade.DeductMoney("test", value); err != nil {
		t.Fatalf("Error on money Deduct: %s", err)
	}

	if facade.wallet.balance != 0 {
		t.Fatalf("Wallet balance is wrong: want: %d, got: %f", 0, facade.wallet.balance)
	}
}
