package main

import (
	"fmt"
	"sync"
)

// TODO: Создайте структуру User с полями ID, Name и Balance
type User struct {
	ID      string
	Name    string
	Balance float64
	mu      sync.Mutex
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

// TODO: Создать функцию Worker, которая будет обрабатывать транзакции
// В эту функцию перенесется так же логика, которую выполняли в Части 2 - "Обрабатываем транзакции по очереди"
// Только массив по которому будем идти не TransactionQueue, а наш канал ch

func (ps *PaymentSystem) Worker(wg *sync.WaitGroup, ch <-chan Transaction) error {
	defer wg.Done()
	for transaction := range ch {
		err := ps.ProcessingTransactions(transaction)
		if err != nil {
			fmt.Printf("worker: failed to process transaction from %q to %q (amount %.2f): %w",
				transaction.FromID, transaction.ToID, transaction.Amount, err)
			continue
		}
	}
	return nil
}

// TODO: Реализуйте метод ProcessingTransactions. Метод должен возвращать ошибки если:
// 1. Пользователь ToID не найден
// 2. Пользователь FromID не найден
// 3. У пользователя баланс меньше чем перевод (получить ошибку из Withdraw)

func (p *PaymentSystem) ProcessingTransactions(tx Transaction) error {
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
	return nil
}

// TODO: Реализуйте метод Deposit
func (u *User) Deposit(amount float64) {
	u.mu.Lock()
	defer u.mu.Unlock()
	u.Balance += amount
}

// TODO: Реализуйте метод Withdraw с проверкой на недостаток средств
func (u *User) WithDraw(amount float64) error {
	u.mu.Lock()
	defer u.mu.Unlock()
	if u.Balance < amount {
		return fmt.Errorf("insufficient funds: balance %.2f, requested %.2f", u.Balance, amount)
	}

	u.Balance -= amount
	return nil
}

func main() {
	// DONE: Инициализируем структуру PaymentSystem через указатель
	ps := NewPaymentSystem()

	fmt.Println("Создаю UserID: 1 с балансом 1000")
	fmt.Println("Создаю UserID: 2 с балансом 500")

	// DONE: Добавить пользователей через AddUser
	ps.AddUser("1")
	ps.Users["1"].Deposit(1000)
	ps.AddUser("2")
	ps.Users["2"].Deposit(500)

	fmt.Println("Перевожу с UserID: 1 на UserID: 2 сумму в размере 200")
	fmt.Println("Перевожу с UserID: 2 на UserID: 1 сумму в размере 50")

	// DONE: Добавляем транзакции
	ps.AddTransaction(Transaction{FromID: "1", ToID: "2", Amount: 200})
	ps.AddTransaction(Transaction{FromID: "2", ToID: "1", Amount: 50})

	// TODO: Создать канал ch := ...
	ch := make(chan Transaction, len(ps.TransactionQueue))

	// TODO: Инициализируем sync.WaitGroup
	var wg sync.WaitGroup

	// TODO: Создаем несколько горутин через цикл for i := 0 ...
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go ps.Worker(&wg, ch)
	}

	// TODO: Передаем все транзакции в канал и завершаем/закрываем наш канал
	for _, tx := range ps.TransactionQueue {
		ch <- tx
	}
	close(ch)

	// TODO: Ждём завершения всех горутин
	wg.Wait()

	fmt.Println("Итого")
	fmt.Printf("У первого пользователя должно получиться 850, а получилось %f\n", ps.Users["1"].Balance)
	fmt.Printf("У второго пользователя должно получиться 650, а получилось %f", ps.Users["2"].Balance)
}
