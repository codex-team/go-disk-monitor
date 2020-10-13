package main

import (
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"
)

var codexBotURL = flag.String("webhook", "", "notification URI from CodeX Bot")
var alertLevel = flag.Int("alert", 90, "disk usage percentage to trigger notification")
var path = flag.String("path", "/", "path to the disk volume")
var debug = flag.Bool("debug", false, "show debug information")

// Hostname - get current server hostname
func Hostname() string {
	name, err := os.Hostname()
	if err != nil {
		log.Fatalf("error while getting hostname: %v", err)
	}

	return name
}

func main() {
	flag.Parse()
	disk := DiskUsage(*path)
	if *debug {
		log.Printf("All: %.2f GB\n", float64(disk.All)/float64(GB))
		log.Printf("Used: %.2f GB\n", float64(disk.Used)/float64(GB))
		log.Printf("Free: %.2f GB\n", float64(disk.Free)/float64(GB))
		log.Printf("Percent: %.0f%%\n", 100*(float64(disk.Used)/float64(disk.All)))
	}

	percentUsed := int(100 * (float64(disk.Used) / float64(disk.All)))
	hostname := Hostname()

	if percentUsed >= *alertLevel {

		alertMessage := fmt.Sprintf("🔥🚒 Running out of space `%.2fGB(%d%%)` on server %s", float64(disk.Free)/float64(GB), percentUsed, hostname)

		// notify via CodeX Bot
		if *codexBotURL != "" {
			data := url.Values{}
			data.Set("message", alertMessage)
			data.Set("parse_mode", "Markdown")

			_, err := MakeHTTPRequest("POST", *codexBotURL, []byte(data.Encode()), map[string]string{
				"Content-Type": "application/x-www-form-urlencoded",
			})
			if err != nil {
				log.Fatalf("error while sending webhook: %v", err)
			}
		}

		log.Println(alertMessage)
	}
}
