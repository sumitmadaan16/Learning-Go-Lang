package main

import (
	"design_pattern/decorator"
	"design_pattern/factory"
	"design_pattern/observer"
	"design_pattern/singleton"
	"fmt"
)

func main() {
	fmt.Println()
	fmt.Println("Singleton Pattern....")
	singleton.SingletonPattern()
	fmt.Println()
	fmt.Println("Factory Pattern....")
	factory.FactoryPattern()
	fmt.Println()
	fmt.Println("Observer Pattern.... ")
	observer.ObserverPattern()
	fmt.Println()
	fmt.Println("Decorator Pattern....")
	decorator.DecoratorPattern()
}
