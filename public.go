package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"math/rand"
	"os"
	"time"

	// my lib
	pnt "print"
)

func createUserID() string {
	rand.Seed(time.Now().UnixNano())
	const pool = "qazwsxedcrfvtgbyhnujmikolpQAZWSXEDCRFVTGBYHNUJMIKOLP1234567890"
	bytes := make([]byte, 15)
	for i := 0; i < 15; i++ {
		bytes[i] = pool[rand.Intn(len(pool))]
	}
	return string(bytes)
}

// JSON -> struct
func parseJSON(unmsg *[]byte, v interface{}) error {

	dec := json.NewDecoder(bytes.NewReader(*unmsg))
	for {
		if err := dec.Decode(&v); err == io.EOF {
			break
		} else if err != nil {
			return err
		}
	}
	return nil
}

// struct -> JSON
func reParseJSON(v interface{}) []byte {
	textbyte, err := json.Marshal(v)
	if err != nil {
		pnt.Errorwd(v, err)
	}
	return textbyte
}

func notExistFile(fn string) bool {
	_, err := os.Stat(fn)
	return errors.Is(err, os.ErrNotExist)
}
func existFile(fn string) bool {
	_, err := os.Stat(fn)
	if err == nil {
		return true
	}
	return errors.Is(err, os.ErrExist)
}
