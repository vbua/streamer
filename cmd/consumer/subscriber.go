package main

import (
	"encoding/json"
	"github.com/nats-io/stan.go"
	"github.com/vbua/streamer"
	"log"
)

var cache = make(map[string]string)

func subscribeToNats() {
	sc, err := stan.Connect(streamer.NutsClusterId, "streamer_consumer")
	if err != nil {
		log.Fatal("Couldn't connect: ", err.Error())
	}

	_, err = sc.Subscribe(streamer.NutsSubject, func(m *stan.Msg) {
		log.Println("Got a new message")
		var o Order
		err := json.Unmarshal(m.Data, &o)
		if err != nil {
			m.Ack()
			log.Println("Wrong json: ", err.Error())
			return
		}

		// валидация json
		err = validate.Struct(&o)
		if err != nil { // если валидацию не проходит, тогда игнорим
			m.Ack()
			log.Println("Wrong json: ", err.Error())
			return
		}

		// пишем в кэш
		cache[o.OrderUid] = string(m.Data)

		// пишем в базу
		insertOrderToDb(o.OrderUid, string(m.Data))

		// говорим стримингу, что все прошло успешно
		m.Ack()
		log.Println("Notified streaming all is good")
	}, stan.DurableName("streamer"), stan.SetManualAckMode())

	if err != nil {
		log.Fatal("Couldn't subscribe: ", err.Error())
	}
}
