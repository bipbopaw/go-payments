package main

import "fmt"

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
		return fmt.Errorf("user is nil")
	}

	if ps.Users == nil {
		ps.Users = make(map[string]*User)
	}

	if _, exists := ps.Users[user.ID]; exists {
		return fmt.Errorf("user already exists")
	}

	ps.Users[user.ID] = user
	return nil
}

func (ps *PaymentSystem) AddTransaction(t Transaction) {
	ps.Transactions = append(ps.Transactions, t)
}

func (ps *PaymentSystem) ProcessingTransactions(t Transaction) error {
	fromUser, ok := ps.Users[t.FromID]
	if !ok {
		return fmt.Errorf("from user not found")
	}

	toUser, ok := ps.Users[t.ToID]
	if !ok {
		return fmt.Errorf("to user not found")
	}

	if err := fromUser.Withdraw(t.Amount); err != nil {
		return err
	}

	toUser.Deposit(t.Amount)
	return nil
}

func (u *User) Deposit(amount float64) {
	u.Balance += amount
}

func (u *User) Withdraw(amount float64) error {
	if u.Balance < amount {
		return fmt.Errorf("insufficient funds")
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
			fmt.Println(err)
		}
	}

	fmt.Println("Итого")
	fmt.Println("User1:", ps.Users["1"].Balance) // 850
	fmt.Println("User1:", ps.Users["2"].Balance) // 650
}
