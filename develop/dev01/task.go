package dev01

import (
	"fmt"
	"os"
	"time"

	"github.com/beevik/ntp"
)

type Time struct {
}

func (t *Time) Now() time.Time {
	// Получаем текущее время с использованием NTP-сервера
	response, err := ntp.Query("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка при запросе NTP-сервера: %s\n", err)
		os.Exit(1)
	}

	// Вычисляем разницу между локальным временем и точным временем из NTP-ответа
	offset := response.ClockOffset.Nanoseconds()
	return time.Now().Add(time.Duration(offset))
}
