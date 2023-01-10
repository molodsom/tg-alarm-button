package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/user"
	"strings"
)

func main() {
	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", accessToken)
	postBody, _ := json.Marshal(map[string]string{
		"chat_id": chatId,
		"text": fmt.Sprintf("🔥 THE ALARM BUTTON WAS PUSHED 🔥"+"\n\n"+"🖥️ %s"+"\n"+"👤 %s"+"\n"+"🔌 %s",
			getHostname(), getUsername(), getIPs(),
		),
	})
	responseBody := bytes.NewBuffer(postBody)
	_, err := http.Post(url, "application/json", responseBody)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("OK")
}

func getUsername() string {
	u, err := user.Current()
	if err != nil {
		log.Fatalf(err.Error())
		return ""
	}
	return u.Username
}

func getHostname() string {
	hostname, err := os.Hostname()
	if err != nil {
		return "Unknown hostname"
	}
	return hostname
}

func getIPs() string {
	var ips []string
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, addr := range addrs {
		ips = append(ips, addr.String())
	}
	return strings.Join(ips, ", ")
}
