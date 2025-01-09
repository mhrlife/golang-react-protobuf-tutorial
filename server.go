package main

import (
	"ProtobufTutorial/internal/schema"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"google.golang.org/protobuf/proto"
	"io"
	"math/rand"
)

func main() {
	board := &schema.Board{
		Width:  1000,
		Height: 1000,
	}

	board.Pixels = make([]uint32, board.Width*board.Height)

	for range board.Height {
		for range board.Width {
			board.Pixels = append(board.Pixels, uint32(rand.Intn(10)))
		}
	}

	echoServer := echo.New()

	echoServer.Use(middleware.CORS())

	echoServer.GET("/board", func(c echo.Context) error {
		boardBytes, _ := proto.Marshal(board)

		return c.Blob(200, "application/protobuf", boardBytes)
	}, middleware.Gzip())

	echoServer.POST("/request", func(c echo.Context) error {
		body, err := io.ReadAll(c.Request().Body)
		if err != nil {
			return err
		}

		request := &schema.Request{}
		err = proto.Unmarshal(body, request)
		if err != nil {
			return err
		}

		switch request.GetRequest().(type) {
		case *schema.Request_Ping:

		case *schema.Request_GetMe:
		}
	})

	echoServer.Logger.Error(echoServer.Start(":8086"))
}
