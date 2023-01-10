package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
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
	r, err := http.Post(url, "application/json", bytes.NewBuffer(postBody))
	if err != nil {
		log.Fatalln(err)
	}
	if r.StatusCode != 200 {
		b, _ := io.ReadAll(r.Body)
		log.Fatalln(string(b))
	}
	log.Println("OK")
}

func getUsername() string {
	u, err := user.Current()
	if err != nil {
		log.Println(err.Error())
		return "Unknown username"
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
		return "Unknown address"
	}
	for _, addr := range addrs {
		ips = append(ips, addr.String())
	}
	return strings.Join(ips, ", ")
}
