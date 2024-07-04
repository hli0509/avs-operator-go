package main

import (
	"avs-operator-go/config"
	"avs-operator-go/operator"
)

func main() {
	config := config.LoadConfig()

	op := operator.NewOperator(config)
	op.RegisterOperator()
	
	if err := op.MonitorNewTasks(); err != nil {
		panic(err)
	}
}
