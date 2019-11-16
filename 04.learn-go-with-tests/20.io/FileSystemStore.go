package main

import (
	"io"
	"encoding/json"
)

type FileSystemStore struct {
	database io.Reader
}

func (f *FileSystemStore) GetLeague() []Player {
	var league []Player
	json.NewDecoder(f.database).Decode(&league)
	return league
}
