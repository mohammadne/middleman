package client

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/labstack/gommon/random"
	"github.com/mohammadne/middleman/internal/handlers"
	"github.com/mohammadne/middleman/pkg/model"
)

const (
	keyLength   = 5
	valueLength = 10
	sleepTime   = time.Second / 5
)

func Setup(path string, requestsNum int) {
	rand.Seed(time.Now().UnixNano())
	handler := handlers.ClientHandler{RequestUrl: path}

	// handler.Get("amir")

	// handler.Post(
	// 	&Body{
	// 		Key:   "mohammad",
	// 		Value: "loves google",
	// 		Cache: true,
	// 	},
	// )

	for index := 0; index < requestsNum; index++ {
		handler.Post(
			&model.Body{
				Key:   random.String(keyLength),
				Value: random.String(valueLength),
				Cache: index%2 == 0,
			},
		)

		time.Sleep(sleepTime)
	}

	fmt.Println("Finish client requests")
}
