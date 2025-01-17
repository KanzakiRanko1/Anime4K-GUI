package main

import (
	"fmt"
	g "github.com/AllenDang/giu"
	"log"
	"os"
	"time"
)

func formatMegabytes(bytes int64) string {
	return fmt.Sprintf("%d MB", bytes/(1024*1024))
}

func formatMillis(millis int64) string {
	calc := millis / 1000
	seconds := millis / 1000 % 60
	minutes := (calc % 3600) / 60
	hours := calc / 3600

	if hours < 1 {
		return fmt.Sprintf("%02d:%02d", minutes, seconds)
	}

	return fmt.Sprintf("%02d:%02d:%02d", hours, minutes, seconds)
}

func logDebug(message string, additionalLineBefore bool) {
	if debug {
		logMessage("[Debug] "+message, additionalLineBefore)
	}
}

func logMessage(message string, additionalLineBefore bool) {
	if additionalLineBefore {
		logs += "\n"
	}

	logs += fmt.Sprintf("[%s] %s\n", time.Now().Format("15:04:05"), message)
	g.Update()
}

func handleSoftError(message string, logs string) {
	logMessage("Critical error occurred, upscaling has been stopped", true)
	logMessage("Please make issue on GitHub with logs, upscaling settings and your computer specification", false)
	logMessage(message, true)
	logMessage(logs, false)

	for i := 0; i < len(animeList); i++ {
		animeList[i].Status = Error
	}

	currentSpeed = "Speed:"
	currentTime = "Time:"
	g.Update()
}

func calcFinished() int {
	i := 0
	for _, anime := range animeList {
		if anime.Status == Finished {
			i++
		}
	}

	return i
}

func checkDebugParam() {
	if len(os.Args) != 0 {
		for _, arg := range os.Args {
			if arg == "--debug" {
				debug = true
				break
			}
		}
	}
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
