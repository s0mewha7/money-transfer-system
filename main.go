package main

import "fmt"

// TODO: Создайте структуру User с полями ID, Name и Balance
type User struct {
	ID      string
	Name    string
	Balance float64
}

// TODO: Создайте структуру Transaction с полями FromID, ToID и Amount
type Transaction struct {
	FromID string
	ToID   string
	Amount float64
}

// TODO: Создайте структуру PaymentSystem с полями Users (мапа типа string с указателем на User), TransactionQueue (слайс Transaction)
type PaymentSystem struct {
	Users            map[string]*User
	TransactionQueue []Transaction
}

func NewPaymentSystem() *PaymentSystem {
	return &PaymentSystem{
		Users:            make(map[string]*User),
		TransactionQueue: []Transaction{},
	}
}

// TODO: Реализуйте метод AddUser
func (p *PaymentSystem) AddUser(ID string) {
	p.Users[ID] = &User{ID: ID}
}

// TODO: Реализуйте метод AddTransaction
func (p *PaymentSystem) AddTransaction(tx Transaction) {
	p.TransactionQueue = append(p.TransactionQueue, tx)
}

// TODO: Реализуйте метод ProcessingTransactions. Метод должен возвращать ошибки если:
// 1. Пользователь ToID не найден
// 2. Пользователь FromID не найден
// 3. У пользователя баланс меньше чем перевод (получить ошибку из Withdraw)

func (p *PaymentSystem) ProcessingTransactions() error {
	for _, tx := range p.TransactionQueue {
		to, ok := p.Users[tx.ToID]
		if !ok {
			return fmt.Errorf("recipient %q not found", tx.ToID)
		}

		from, ok := p.Users[tx.FromID]
		if !ok {
			return fmt.Errorf("sender %q not found", tx.FromID)
		}

		if err := from.WithDraw(tx.Amount); err != nil {
			return err
		}

		to.Deposit(tx.Amount)
	}
	return nil
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
	//// TODO: Создайте несколько пользователей user1, user2
	//user1 := &User{ID: "1", Name: "User1", Balance: 1000}
	//user2 := &User{ID: "2", Name: "User2", Balance: 0}
	//
	//// TODO: Проверьте работу методов Deposit и Withdraw
	//
	//user1.Deposit(120.0)
	//err := user1.WithDraw(100.0)
	//
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//user2.Deposit(100)
	//
	//err2 := user2.WithDraw(200.0)
	//
	//if err2 != nil {
	//	fmt.Println(err2)
	//}
	//
	//// TODO: Выведите информацию о пользователях
	//fmt.Printf("ID: %s, Name: %s, Balance: %.2f\n", user1.ID, user1.Name, user1.Balance)
	//fmt.Printf("ID: %s, Name: %s, Balance: %.2f\n", user2.ID, user2.Name, user2.Balance)

	// TODO: Инициализируем структуру PaymentSystem через указатель, чтобы работать с ней напрямую и сохранять изменения
	NewPayment := NewPaymentSystem()

	// TODO: Добавить пользователей через AddUser
	fmt.Println("Создаю UserID: 1 с балансом 1000")
	NewPayment.AddUser("1")
	NewPayment.Users["1"].Deposit(1000)

	fmt.Println("Создаю UserID: 2 с балансом 500")
	NewPayment.AddUser("2")
	NewPayment.Users["2"].Deposit(500)

	// TODO: Добавляем транзакции
	fmt.Println("Перевожу с UserID: 1 на UserID: 2 сумму в размере 200")
	NewPayment.AddTransaction(
		Transaction{
			FromID: "1",
			ToID:   "2",
			Amount: 200,
		})

	fmt.Println("Перевожу с UserID: 2 на UserID: 1 сумму в размере 50")
	NewPayment.AddTransaction(
		Transaction{
			FromID: "2",
			ToID:   "1",
			Amount: 50,
		})

	// TODO: Обрабатываем транзакции по очереди
	if err := NewPayment.ProcessingTransactions(); err != nil {
		fmt.Println("Ошибка обработки:", err)
	}

	// Итог
	fmt.Println("Итого")
	fmt.Printf("User 1: %.2f (ожидалось 850)\n", NewPayment.Users["1"].Balance)
	fmt.Printf("User 2: %.2f (ожидалось 650)\n", NewPayment.Users["2"].Balance)
}
