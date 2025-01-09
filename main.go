package main

import (
	"ProtobufTutorial/internal/schema"
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"google.golang.org/protobuf/proto"
	"math/rand"
	"os"
	"time"
)

func main() {
	watch("json", boardJSON)
	watch("proto", boardProto)
}

type Board struct {
	Width  int
	Height int
	Pixels []int // 0 = empty, otherwise player id
}

func boardJSON() {
	board := Board{
		Width:  1000,
		Height: 1000,
	}

	board.Pixels = make([]int, board.Width*board.Height)

	for range board.Height {
		for range board.Width {
			board.Pixels = append(board.Pixels, rand.Intn(10))
		}
	}

	// convert to json
	boardBytes, _ := json.Marshal(board)
	os.WriteFile("board.json", compress(boardBytes), 0644)
}

func boardProto() {
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

	boardBytes, _ := proto.Marshal(board)
	os.WriteFile("board.bin", compress(boardBytes), 0644)
}

func compress(file []byte) []byte {
	var writer bytes.Buffer
	w, _ := gzip.NewWriterLevel(&writer, gzip.BestCompression)
	w.Write(file)
	w.Close()

	return writer.Bytes()
}

func watch(name string, exec func()) {
	startTime := time.Now()
	exec()
	elapsedTime := time.Since(startTime)
	fmt.Println(name, "took", elapsedTime)
}
