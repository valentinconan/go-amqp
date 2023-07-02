package amqp

import (
	"go-amqp/src/amqp/consumers"
)
func Init(){

	//launch consumers in go routines
	go consumers.Consumer()
	go consumers.Consumer1()
}
