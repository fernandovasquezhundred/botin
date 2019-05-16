package main

import (
	"fmt"
	"log"
	"time"

	"github.com/yanzay/tbot"
)
import bolt "go.etcd.io/bbolt"

func main() {
	bot := tbot.New("872512071:AAGqTVYoot2MJphPFtfGixmC2oo9UM9MEak")
	c := bot.Client()

	bot.HandleMessage(".*yo.*", func(m *tbot.Message) {
		c.SendChatAction(m.Chat.ID, tbot.ActionTyping)
		time.Sleep(1 * time.Second)
		c.SendMessage(m.Chat.ID, "hello!")
	})

	bot.HandleMessage("read", func(m *tbot.Message) {
		c.SendMessage(m.Chat.ID, "init")
		db, err := bolt.Open("C:/tmp/bolt.db", 0666, nil)
		if err != nil {
			fmt.Sprintf(err.Error())
		}

		defer db.Close()

		// TODO: read boltDB book index / bbolt
		// TODO: read next X words  starting from index
		// TODO: send words to bot
		// TODO: update index
	})

	err := bot.Start()
	if err != nil {
		log.Fatal(err)
	}
}
