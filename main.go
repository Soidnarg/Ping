package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	probing "github.com/prometheus-community/pro-bing"
)

func main() {
	count := 2
	timer := 2000

	if len(os.Args) >= 4 {
		if whoPing, err := strconv.Atoi(os.Args[3]); err == nil && whoPing != 0 {
			count = whoPing
		}
	}

	if len(os.Args) >= 3 {
		if whoPing, err := strconv.Atoi(os.Args[2]); err == nil && whoPing != 0 {
			timer = whoPing
		}
	}

	if len(os.Args) >= 2 {
		if os.Args[1] == "-h" || os.Args[1] == "--help" {
			fmt.Println(Help_string)
			return
		}

		whoPing := os.Args[1]
		if whoPing == "" || strings.Count(whoPing, ".") != 2 {
			log.Println("network to ping needed, example 192.168.1")
		} else {
			log.Println(strings.Count(whoPing, "."))
			for ip := 1; ip < 255; ip++ {
				pinger, err := probing.NewPinger(whoPing + "." + strconv.Itoa(ip))
				if err != nil {
					log.Println("ERRORE CREATORE PING")
				}
				pinger.Count = count
				pinger.Timeout = time.Duration(timer) * time.Millisecond
				pinger.SetPrivileged(true)
				pinger.Run()
				/* Non ci interessano eventuali errori */
				//		if err != nil {
				//			log.Println("ERRORE ESECUZIONE PING 192.168.1." + strconv.Itoa(ip) + err.Error())
				//		}
				stats := pinger.Statistics()
				if stats.PacketsRecv > 0 {
					log.Println("192.168.1." + strconv.Itoa(ip))
				}
			}
		}
	} else {
		log.Println("network to ping needed, example 192.168.1")
	}
}
