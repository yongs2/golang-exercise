package main

import (
	"encoding/json"
	"testing"
	"time"
)

type Data struct {
	ID               int64        `json:"id"`
	SomeForeignKeyID int64        `json:"some_foreign_key_id"`
	SomeForeignData  *ForeignData `json:"some_foreign_data"`
	TextData         string       `json:"text_data"`
}

type ForeignData struct {
	ID       int64  `json:"id"`
	TextData string `json:"text_data"`
}

func JSONMarshal(data *Data, b *testing.B) {
	_, err := json.Marshal(data)
	if err != nil {
		b.Fatalf("json marshal failed: %v", err)
	}
}

func JSONUnmarshal(data string, b *testing.B) {
	var dataStruct *Data
	err := json.Unmarshal([]byte(data), &dataStruct)
	if err != nil {
		b.Fatalf("json unmarshal failed: %v", err)
	}
}

func BenchmarkJSONMarshal(b *testing.B) {
	id := int64(time.Now().Nanosecond())
	data := &Data{
		ID:               id,
		SomeForeignKeyID: id,
		SomeForeignData: &ForeignData{
			ID:       id,
			TextData: "Hello World!",
		},
		TextData: "Hello World",
	}

	for i := 0; i < b.N; i++ {
		JSONMarshal(data, b)
	}
}

func BenchmarkJSONUnmarshal(b *testing.B) {
	dataStr := `{"id":922232846,"some_foreign_key_id":922232846,"some_foreign_data":{"id":922232846,"text_data":"Hello World!"},"text_data":"Hello World"}`
	for i := 0; i < b.N; i++ {
		JSONUnmarshal(dataStr, b)
	}
}
