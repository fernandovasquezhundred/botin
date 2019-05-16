package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/yanzay/tbot"
)
import bolt "go.etcd.io/bbolt"

func Find(fname string, from, to int, needle []byte) (bool, error) {
	f, err := os.Open(fname)
	if err != nil {
		return false, err
	}
	defer f.Close()
	n := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		n++
		if n < from {
			continue
		}
		if n > to {
			break
		}
		if bytes.Index(scanner.Bytes(), needle) >= 0 {
			return true, nil
		}
	}
	return false, scanner.Err()
}

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

		db.View(func(tx *bolt.Tx) error {

			b := tx.Bucket([]byte("books"))
			v := b.Get([]byte("book1"))
			fmt.Printf("The answer is: %s\n", v)

			found, err := Find("C:\tmp\EgoIstheEnemybyRyanHoliday.txt", 18, 27, []byte("Hello World"))
			c.SendMessage(found, "init")
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
