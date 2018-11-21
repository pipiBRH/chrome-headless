package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
)

func GetOutboundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		res := getWS()
		fmt.Fprintf(w, res)
	})
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func getWS() string {
	url := fmt.Sprintf("http://%v:9222/json/version", GetOutboundIP())
	req, _ := http.NewRequest("GET", url, nil)
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return ""
	}
	return string(body)
}
