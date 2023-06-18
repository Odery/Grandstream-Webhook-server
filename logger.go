// loggerService.go

package main

import (
	"log"
	"os"
	"sync"
)

type LoggerService struct {
	file   *os.File
	logger *log.Logger
	mux    sync.Mutex
}

func NewLoggerService(filename string) (*LoggerService, error) {
	f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return nil, err
	}

	logger := log.New(f, "", log.LstdFlags)

	return &LoggerService{
		file:   f,
		logger: logger,
	}, nil
}

func (ls *LoggerService) Log(message string) {
	ls.mux.Lock()
	defer ls.mux.Unlock()

	log.Println(message)       // Print to console
	ls.logger.Println(message) // Print to file
}

func (ls *LoggerService) Close() error {
	return ls.file.Close()
}
