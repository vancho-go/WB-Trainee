/*
=== Базовая задача ===
Создать программу печатающую точное время с использованием NTP библиотеки.Инициализировать как go module.
Использовать библиотеку https://github.com/beevik/ntp.
Написать программу печатающую текущее время / точное время с использованием этой библиотеки.
Программа должна быть оформлена с использованием как go module.
Программа должна корректно обрабатывать ошибки библиотеки: распечатывать их в STDERR и возвращать ненулевой код выхода в OS.
Программа должна проходить проверки go vet и golint.
*/

package ntp_time

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/beevik/ntp"
)

func GetCurretTime() time.Time {
	time, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		_, err := io.WriteString(os.Stderr, err.Error())
		if err != nil {
			fmt.Printf("Write to stderr error: %s", err)
		}
	}
	return time
}
