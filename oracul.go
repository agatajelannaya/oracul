package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	x := []string{
		"да",
		"конечно!",
		"вероятность велика",
		"вероятность мала",
		"неизвестно",
		"лучше пока не говорить тебе об этом",
		"нет",
		"возлюби ближнего своего",
		"переформулируй вопрос и задай снова",
	}
	fmt.Println("Подумайте о вопросе...")
	time.Sleep(3000 * time.Millisecond)
	fmt.Println("Приготовьтесь!")
	time.Sleep(2000 * time.Millisecond)
	fmt.Println("Оракул говорит: ", x[rand.Intn(len(x))])
}
