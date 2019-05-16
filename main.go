package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/yanzay/tbot"
)
import bolt "go.etcd.io/bbolt"

func Find(filename string, from, to int) (string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer f.Close()
	n := 0
	scanner := bufio.NewScanner(f)

	text := ""
	for scanner.Scan() {
		n++
		if n >= from {
			text = text + scanner.Text()
		}
		if n > to {
			break
		}
	}
	return text, scanner.Err()
}

func main() {
	bot := tbot.New(os.Getenv("TELEGRAM_TOKEN"))
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
			fmt.Printf(err.Error())
		}
		defer db.Close()

		db.View(func(tx *bolt.Tx) error {

			b := tx.Bucket([]byte("books"))
			line := b.Get([]byte("book1"))

			linex, err := strconv.Atoi(string(line))
			fmt.Printf("The answer is: %s\n", line)

			found, err := Find("C:/zxc/learn/self/ego.txt",
				linex,
				linex + 2)
			c.SendMessage(m.Chat.ID, found)
			fmt.Println(found, err)
			return nil
		})

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
