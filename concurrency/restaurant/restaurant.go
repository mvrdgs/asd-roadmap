package restaurant

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

var (
	DishList = map[string]time.Duration{
		"pizza":  5 * time.Millisecond,
		"pasta":  3 * time.Millisecond,
		"salad":  2 * time.Millisecond,
		"burger": 4 * time.Millisecond,
		"steak":  6 * time.Millisecond,
	}

	WaiterList = []string{
		"John",
		"Mark",
		"Luke",
		"Matthew",
		"Peter",
	}
)

type Order struct {
	ID     int64  `json:"id"`
	Dish   string `json:"dish"`
	Waiter string `json:"waiter"`
	Ready  bool   `json:"ready"`
}

type Restaurant struct {
	IsOpen        bool
	CookersNumber int
	Menu          []string
	OrderQueue    chan *Order
	Done          chan *Order
	OrdersCounter uint
	DoneCounter   uint

	log bool
}

func NewRestaurant(cookersNumber int, log bool) *Restaurant {
	var menu []string
	for dish := range DishList {
		menu = append(menu, dish)
	}

	return &Restaurant{
		IsOpen:        false,
		CookersNumber: cookersNumber,
		Menu:          menu,
	}
}

func (r *Restaurant) Open() {
	r.IsOpen = true
	r.OrderQueue = make(chan *Order, r.CookersNumber)
	r.Done = make(chan *Order)

	go r.StartPreparingDishes()
	go r.ServeDish()
}

func (r *Restaurant) Close() uint {
	r.IsOpen = false

	for r.OrdersCounter > r.DoneCounter {
		// wait for all orders to be done
	}

	close(r.OrderQueue)
	close(r.Done)

	if r.log {
		log.Println("Total orders:", r.OrdersCounter)
	}

	return r.OrdersCounter
}

func (r *Restaurant) AddNewOrder() {
	if !r.IsOpen {
		log.Println("restaurant is closed")
		return
	}

	waiter := rand.Intn(len(WaiterList))
	dish := rand.Intn(len(r.Menu))
	order := Order{
		ID:     int64(r.OrdersCounter + 1),
		Dish:   r.Menu[dish],
		Waiter: WaiterList[waiter],
		Ready:  false,
	}

	r.OrderQueue <- &order
	r.OrdersCounter++
}

func (r *Restaurant) StartPreparingDishes() {
	for order := range r.OrderQueue {
		go r.prepareDish(order)
	}
}

func (r *Restaurant) prepareDish(order *Order) {
	time.Sleep(DishList[order.Dish])
	order.Ready = true
	r.Done <- order
}

func (r *Restaurant) ServeDish() {
	m := sync.Mutex{}
	for order := range r.Done {
		if order.Ready {
			m.Lock()
			r.DoneCounter++
			m.Unlock()

			if r.log {
				log.Println("order ready", order)
			}
		}
	}
}
