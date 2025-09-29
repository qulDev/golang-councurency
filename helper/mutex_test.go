package helper

import (
	"log"
	"sync"
	"testing"
	"time"
)

func TestMutex(t *testing.T) {
	var x int = 0
	var mutex sync.Mutex
	for i := 1; i <= 1000; i++ {
		go func() {
			for j := 1; j <= 100; j++ {
				mutex.Lock()
				x = x + 1
				mutex.Unlock()
			}
		}()
	}

	time.Sleep(5 * time.Second)
	t.Logf("Counter: %d", x)
}

type BankAccount struct {
	RWMutex sync.RWMutex
	Balance int
}

func (b *BankAccount) AddBalance(amount int) {
	b.RWMutex.Lock()
	b.Balance = b.Balance + amount
	b.RWMutex.Unlock()
}

func (b *BankAccount) GetBalance() int {
	b.RWMutex.RLock()
	balance := b.Balance
	b.RWMutex.RUnlock()
	return balance
}

func TestRWMutex(t *testing.T) {
	account := BankAccount{}
	for i := 1; i <= 100; i++ {
		go func() {
			for j := 1; j <= 100; j++ {
				account.AddBalance(1)
				t.Logf("Balance: %d", account.GetBalance())
			}
		}()
	}
	time.Sleep(5 * time.Second)
	t.Logf("Final Balance: %d", account.GetBalance())
}

type UserBelance struct {
	sync.Mutex
	Name    string
	Balance int
}

func (user *UserBelance) Lock() {
	user.Mutex.Lock()
}

func (user *UserBelance) Unlock() {
	user.Mutex.Unlock()
}

func (user *UserBelance) Change(amount int) {
	user.Balance = user.Balance + amount
}

func Tranfer(user1, user2 *UserBelance, amount int) {
	user1.Lock()
	log.Println("Lock", user1.Name)
	user1.Change(-amount)

	time.Sleep(1 * time.Second)

	user2.Lock()
	log.Println("Lock", user2.Name)
	user2.Change(amount)

	user1.Unlock()
	user2.Unlock()
	log.Println("Unlock", user1.Name)
	log.Println("Unlock", user2.Name)
}

func TestDeadLock(t *testing.T) {
	userA := UserBelance{
		Name:    "Budi",
		Balance: 100000,
	}

	userB := UserBelance{
		Name:    "Eko",
		Balance: 100000,
	}

	go Tranfer(&userA, &userB, 10000)
	go Tranfer(&userB, &userA, 20000)

	time.Sleep(10 * time.Second)
	log.Println("User A Balance", userA.Balance)
	log.Println("User B Balance", userB.Balance)
}
