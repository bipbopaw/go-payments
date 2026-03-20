package main

import "fmt"

var (
	ErrUserNil           = fmt.Errorf("user is nil")
	ErrUserExists        = fmt.Errorf("user already exists")
	ErrUserNotFound      = fmt.Errorf("user not found")
	ErrInsufficientFunds = fmt.Errorf("insufficient funds")
	ErrInvalidAmount     = fmt.Errorf("invalid amount")
)

type User struct {
	ID      string
	Name    string
	Balance float64
}

type Transaction struct {
	FromID string
	ToID   string
	Amount float64
}

type PaymentSystem struct {
	Users        map[string]*User
	Transactions []Transaction
}

func (ps *PaymentSystem) AddUser(user *User) error {
	if user == nil {
		return ErrUserNil
	}

	if ps.Users == nil {
		ps.Users = make(map[string]*User)
	}

	if _, exists := ps.Users[user.ID]; exists {
		return ErrUserExists
	}

	ps.Users[user.ID] = user
	return nil
}

func (ps *PaymentSystem) AddTransaction(t Transaction) {
	ps.Transactions = append(ps.Transactions, t)
}

func (ps *PaymentSystem) ProcessingTransactions(t Transaction) error {
	if t.Amount <= 0 {
		return ErrInvalidAmount
	}

	fromUser, ok := ps.Users[t.FromID]
	if !ok {
		return fmt.Errorf("from: %w", ErrUserNotFound)
	}

	toUser, ok := ps.Users[t.ToID]
	if !ok {
		return fmt.Errorf("to: %w", ErrUserNotFound)
	}

	if err := fromUser.Withdraw(t.Amount); err != nil {
		return err
	}

	toUser.Deposit(t.Amount)
	return nil
}

func (u *User) Deposit(amount float64) error {
	if amount <= 0 {
		return ErrInvalidAmount
	}
	u.Balance += amount
	return nil
}

func (u *User) Withdraw(amount float64) error {
	if amount <= 0 {
		return ErrInvalidAmount
	}

	if u.Balance < amount {
		return ErrInsufficientFunds
	}

	u.Balance -= amount
	return nil
}

func main() {
	ps := &PaymentSystem{Users: make(map[string]*User)}

	fmt.Println("Создаю UserID: 1 с балансом 1000")
	fmt.Println("Создаю UserID: 2 с балансом 500")

	user1 := &User{ID: "1", Name: "Ben", Balance: 1000}
	user2 := &User{ID: "2", Name: "Tom", Balance: 500}

	if err := ps.AddUser(user1); err != nil {
		fmt.Println(err)
	}
	if err := ps.AddUser(user2); err != nil {
		fmt.Println(err)
	}

	fmt.Println("Перевожу с UserID: 1 на UserID: 2 сумму в размере 200")
	fmt.Println("Перевожу с UserID: 2 на UserID: 1 сумму в размере 50")

	ps.AddTransaction(Transaction{FromID: "1", ToID: "2", Amount: 200})
	ps.AddTransaction(Transaction{FromID: "2", ToID: "1", Amount: 50})

	for _, t := range ps.Transactions {
		if err := ps.ProcessingTransactions(t); err != nil {
			fmt.Println("transaction error:", err)
		}
	}

	fmt.Println("Итого")
	fmt.Println("User1:", ps.Users["1"].Balance) // 850
	fmt.Println("User2:", ps.Users["2"].Balance) // 650
}
