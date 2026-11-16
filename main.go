package main

import "fmt"

// TODO: Создайте структуру User с полями ID, Name и Balance
type User struct {
	ID      string
	Name    string
	Balance float64
}

// TODO: Реализуйте метод Deposit
func (u *User) Deposit(amount float64) {
	u.Balance += amount
}

// TODO: Реализуйте метод Withdraw с проверкой на недостаток средств
func (u *User) WithDraw(amount float64) error {
	if u.Balance < amount {
		return fmt.Errorf("insufficient funds: balance %.2f, requested %.2f", u.Balance, amount)
	}

	u.Balance -= amount
	return nil
}

func main() {
	// TODO: Создайте несколько пользователей user1, user2
	user1 := &User{ID: "1", Name: "User1", Balance: 1000}
	user2 := &User{ID: "2", Name: "User2", Balance: 0}

	// TODO: Проверьте работу методов Deposit и Withdraw

	user1.Deposit(120.0)
	err := user1.WithDraw(100.0)

	if err != nil {
		fmt.Println(err)
	}

	user2.Deposit(100)

	err2 := user2.WithDraw(200.0)

	if err2 != nil {
		fmt.Println(err2)
	}

	// TODO: Выведите информацию о пользователях
	fmt.Printf("ID: %s, Name: %s, Balance: %.2f\n", user1.ID, user1.Name, user1.Balance)
	fmt.Printf("ID: %s, Name: %s, Balance: %.2f\n", user2.ID, user2.Name, user2.Balance)
}
