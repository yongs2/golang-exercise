package main

type state interface {
	addItem(count int) error
	requestItem() error
	insertMoney(money int) error
	dispenseItem() error
}
