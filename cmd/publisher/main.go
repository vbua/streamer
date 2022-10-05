package main

import (
	"fmt"
	"github.com/nats-io/stan.go"
	"github.com/vbua/streamer"
	"io"
	"log"
	"os"
	"time"
)

func main() {
	sc, err := stan.Connect(streamer.NutsClusterId, "streamer_publisher") // коннектимся к натс стриминг
	if err != nil {
		log.Fatal(err.Error())
	}

	files, err := os.ReadDir("models")
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		fmt.Println("Opening file: ", file.Name())

		file, err := os.Open("models/" + file.Name()) // открываем файл с моделью
		if err != nil {
			log.Fatal(err.Error())
		}
		res, err := io.ReadAll(file) // читаем весь файл
		if err != nil {
			log.Fatal(err.Error())
		}

		// шлем сообщение в натс стриминг
		if err := sc.Publish("order", []byte(res)); err != nil {
			log.Fatal(err.Error())
		}
		file.Close()
		time.Sleep(10 * time.Second)
	}
}
