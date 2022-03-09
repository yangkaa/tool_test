package main

import (
	"encoding/json"
	"fmt"
	"github.com/thinkeridea/go-extend/exnet"
	"net/http"
)

var locationsCache map[string]string

func main() {
	http.HandleFunc("/", getIP)
	http.ListenAndServe(":8000", nil)
}

func getIP(w http.ResponseWriter, r *http.Request) {
	ip := exnet.ClientPublicIP(r)
	if ip == "" {
		ip = exnet.ClientIP(r)
	}
	local := getIPLocation(ip)
	fmt.Fprintf(w, fmt.Sprintf("IP: [%s]\nLocal: [%s]", ip, local))
}

func getIPLocation(ip string) string {
	if locationsCache == nil {
		locationsCache = make(map[string]string)
	}
	if location, ok := locationsCache[ip]; ok {
		return location
	}
	res, err := http.Get(fmt.Sprintf("https://ip.zxinc.org/api.php?type=json&ip=%s", ip))
	if err != nil {
		return ""
	}
	if res.Body != nil {
		defer res.Body.Close()
		var info = make(map[string]interface{})
		if err := json.NewDecoder(res.Body).Decode(&info); err == nil {
			if location, ok := info["data"].(map[string]interface{})["location"]; ok {
				locationsCache[ip] = location.(string)
			}
			return locationsCache[ip]
		} else {
			fmt.Println(err)
		}
	}
	return ""
}
