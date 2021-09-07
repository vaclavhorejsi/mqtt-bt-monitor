package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"time"

	"github.com/bamzi/jobrunner"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var cfg config
var runCount int

var c mqtt.Client

type config struct {
	Delay   int
	Debug   bool
	Devices []device
	MQTT    struct {
		Server   string
		Port     string
		Topic    string
		Retained bool
		Username string
		Password string
	}
}

type device struct {
	Name string
	MAC  string
}

type status struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Status    string `json:"status"`
	Timestamp int64  `json:"timestamp"`
}

type checkDevice struct {
}

func (e checkDevice) Run() {
	i := runCount % len(cfg.Devices)

	_, err := exec.Command("l2ping", "-c", "1", cfg.Devices[i].MAC).Output()

	msg := status{
		ID:        cfg.Devices[i].MAC,
		Name:      cfg.Devices[i].Name,
		Timestamp: time.Now().Unix(),
	}

	if err != nil {
		msg.Status = "offline"
	} else {
		msg.Status = "online"
	}

	msgJSON, _ := json.Marshal(msg)

	if cfg.Debug {
		fmt.Println(string(msgJSON))
	}
	c.Publish(cfg.MQTT.Topic+"/"+cfg.Devices[i].Name, 1, cfg.MQTT.Retained, msgJSON)
	c.Publish(cfg.MQTT.Topic, 1, false, "online")

	runCount++
}

func main() {
	configFile, err := ioutil.ReadFile("./config.json")
	if err != nil {
		panic("Unable to read config.json file!")
	}
	if err := json.Unmarshal(configFile, &cfg); err != nil {
		panic("Unable to parse config.json file!")
	}
	fmt.Println("Config.json loaded")

	jobrunner.Start()
	if len(cfg.Devices) > 0 {
		jobrunner.Every(time.Duration(cfg.Delay)*time.Second, checkDevice{})
	}

	hn, _ := os.Hostname()

	opts := mqtt.NewClientOptions().AddBroker("tcp://" + cfg.MQTT.Server + ":" + cfg.MQTT.Port)
	opts.Username = cfg.MQTT.Username
	opts.Password = cfg.MQTT.Password
	opts.SetAutoReconnect(true)
	opts.SetConnectRetry(true)
	opts.SetCleanSession(true)
	opts.SetClientID(hn)
	opts.SetWill(cfg.MQTT.Topic, "offline", 1, false)
	opts.SetOnConnectHandler(func(mqtt.Client) {
		c.Publish(cfg.MQTT.Topic, 1, false, "online")
	})

	c = mqtt.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	select {}
}
