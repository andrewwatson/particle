package particle

import (
	"bufio"
	"crypto/tls"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type Device struct {
	Id        string            `json:"id"`
	Name      string            `json:"name"`
	LastApp   string            `json:"last_app"`
	Connected bool              `json:"connected"`
	LastHeard time.Time         `json:"last_heard"`
	ProductId int               `json:"product_id"`
	Variables map[string]string `json:"variables"`
	Functions []string          `json:"functions"`
}

const DEVICE_URL string = "https://api.particle.io/v1/devices/"

func GetDevice(Id string, token string) Device {
	var client *http.Client
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client = &http.Client{tr, nil, nil, 0 * time.Second}

	req, err := http.NewRequest("GET", DEVICE_URL+Id+"?access_token="+token, nil)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	reader := bufio.NewReader(resp.Body)

	var device Device
	line, err := reader.ReadString('}')
  // log.Printf(line)
	if err != nil {
		log.Fatal(err)
	} else {
		err := json.Unmarshal([]byte(line), &device)
		if err != nil {
			log.Fatal(err)
		}
	}

	return device
}
