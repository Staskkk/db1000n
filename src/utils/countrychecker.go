package utils

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
	"runtime"
)

func openBrowser(url string) {
	switch runtime.GOOS {
	case "windows":
		_ = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		_ = exec.Command("open", url).Start()
	}

	log.Printf("Please open %s", url)
}

// CheckCountry allows to check which country the app is running from
func CheckCountry(countriesToAvoid []string) {
	type IPInfo struct {
		Country string `json:"country"`
	}

	ipCheckerURI := "https://api.myip.com/"

	resp, err := http.Get(ipCheckerURI)
	if err != nil {
		log.Println("Can't check users country. Please manually check that VPN is enabled or that you have non Ukrainian IP address.")

		return
	}

	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Printf("Can't close connection to: %s", ipCheckerURI)

			return
		}
	}()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Can't check users country. Please manually check that VPN is enabled or that you have non Ukrainian IP address.")

		return
	}

	ipInfo := IPInfo{}

	if err = json.Unmarshal(body, &ipInfo); err != nil {
		log.Println("Can't check users country. Please manually check that VPN is enabled or that you have non Ukrainian IP address.")

		return
	}

	for _, country := range countriesToAvoid {
		if ipInfo.Country == country {
			log.Printf("Current country: %s. You might need to enable VPN.", ipInfo.Country)
			openBrowser("https://arriven.github.io/db1000n/vpn/")

			return
		}
	}

	log.Printf("Current country: %s", ipInfo.Country)
}
