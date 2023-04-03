package pattern

import (
	"errors"
)

/*
	Паттерн фасад (Facade).
	Призван создать новый, упрощённый и лимитированный интерфейс к уже существующей системе.

	В примере ниже у нас есть WalletFacade который предоставляет интерфейс для взаимодействия с кошельком.
	Однако взаимодействие с кошельком не ограничивается только им.
	Зачастую оно включает в себя проверку аккаунта, проверки безопасности, оповещения и прочее.
	В примере я использовал только две сущности участвующие в устройстве фасада: Wallet и Account.
	Однако их может больше.

	+ Давая упрощённый доступ к подсистеме мы можем изолировать сложные части подсистемы.
	- Если уйти в крайность, фасад может стать "объектом божественного класса" который будет связывать всю систему приложения.

	Распространённые сценарии использования Фасада:
		- Упрощённый доступ к легаси
		- Упрощение сложных API
		- Упрощение доступа к сторонним сервисам
*/

type WalletFacade struct {
	wallet  *Wallet
	account *Account
}

func newWalletFacade(accountName string) *WalletFacade {
	return &WalletFacade{
		account: newAccount(accountName),
		wallet:  newWallet(),
	}
}

func (wf *WalletFacade) AddMoney(accountName string, value float64) error {
	if err := wf.account.Verify(accountName); err != nil {
		return errors.New("account hasn't been verified")
	}

	wf.wallet.Add(value)

	return nil
}

func (wf *WalletFacade) DeductMoney(accountName string, value float64) error {
	if err := wf.account.Verify(accountName); err != nil {
		return errors.New("account hasn't been verified")
	}

	if err := wf.wallet.Spend(value); err != nil {
		return errors.New("not enough money")
	}

	return nil
}

func newWallet() *Wallet {
	return &Wallet{
		balance: 0,
	}
}

type Wallet struct {
	balance float64
}

func (w *Wallet) Add(value float64) {
	w.balance += value
}

func (w *Wallet) Spend(value float64) error {
	if value > w.balance {
		return errors.New("not enough money in the wallet")
	}
	w.balance -= value
	return nil
}

func newAccount(name string) *Account {
	return &Account{
		name: name,
	}
}

type Account struct {
	name string
}

func (a *Account) Verify(username string) error {
	if username != a.name {
		return errors.New("wrong account")
	}
	return nil
}
