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

func (u *User) WithDraw(amount float64) error {
	if u.Balance < amount {
		return fmt.Errorf("insufficient funds: balance %.2f, requested %.2f", u.Balance, amount)
	}
	u.Balance -= amount
	return nil
}

func main() {
	user1 := &User{ID: "1", Name: "User1", Balance: 1000}
	user2 := &User{ID: "2", Name: "User2", Balance: 0}

	user1.Deposit(120)
	if err := user1.WithDraw(100); err != nil {
		fmt.Println(err)
	}
	user2.Deposit(100)
	if err := user2.WithDraw(200); err != nil {
		fmt.Println(err)
	}

	fmt.Printf("ID: %s, Name: %s, Balance: %.2f\n", user1.ID, user1.Name, user1.Balance)
	fmt.Printf("ID: %s, Name: %s, Balance: %.2f\n", user2.ID, user2.Name, user2.Balance)
}
