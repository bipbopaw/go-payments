package main

import "fmt"

type User struct {
	ID      string
	Name    string
	Balance float64
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
	user1 := &User{ID: "1", Name: "Ben", Balance: 1000}
	user2 := &User{ID: "2", Name: "Tom", Balance: 250}

	err := user1.Withdraw(1000)
	if err != nil {
		fmt.Println("error:", err)
	}
	err = user2.Withdraw(200)
	if err != nil {
		fmt.Println("error:", err)
	}

	user1.Deposit(100)
	user2.Deposit(200)

	fmt.Printf("user: %s - balance: %0.2f\n", user1.Name, user1.Balance)
	fmt.Printf("user: %s - balance: %0.2f\n", user2.Name, user2.Balance)

}
