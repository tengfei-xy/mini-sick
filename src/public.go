package main
import (
	"time"
	"math/rand"
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