package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/zmwilliam/go-amqp-logstash-es/amqp"
)

func main() {
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter text: ")
		text, _ := reader.ReadString('\n')

		m := Message{
			Text:      text,
			CreatedAt: time.Now(),
		}

		content, _ := json.Marshal(m)
		amqp.SendMessage("q_logs", content)
	}

}

type Message struct {
	Text      string
	CreatedAt time.Time
}
