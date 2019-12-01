package main

import (
	"log"
	"net"
	"net/smtp"
	"os"

	"github.com/mhale/smtpd"
	"gopkg.in/yaml.v3"
)

// Config Struct
type Config struct {
	Server struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	} `yaml:"server"`
	Mail struct {
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		SMTP     string `yaml:"smtp"`
		Port     string `yaml:"port"`
	} `yaml:"mail"`
}

var cfg Config

func mailHandler(origin net.Addr, from string, to []string, data []byte) {
	log.Printf("[↓] from: %s, to: %s", from, to)
	sendMail(from, to, data)
}

func sendMail(from string, to []string, data []byte) {
	auth := smtp.PlainAuth("", cfg.Mail.Username, cfg.Mail.Password, cfg.Mail.SMTP)

	err := smtp.SendMail(cfg.Mail.SMTP+":"+cfg.Mail.Port, auth, from, to, data)
	if err != nil {
		panic(err)
	} else {
		log.Printf("[↑] from: %s, to: %s", from, to)
	}
}

func main() {
	f, err := os.Open("config.yml")
	if err != nil {
		panic(err)
	}

	var config Config
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&config)
	if err != nil {
		panic(err)
	}

	cfg = config

	host := cfg.Server.Host
	port := cfg.Server.Port

	log.Printf("[↑↓] %s:%s", host, port)

	smtpd.ListenAndServe(host+":"+port, mailHandler, "SMTP-Relay", "")
}
