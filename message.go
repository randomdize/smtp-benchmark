package main

import (
	"net/mail"
	"os"
	"path/filepath"
	"strings"
)

func getArfMail(file string) (*mail.Message, error) {
	p, err := filepath.Abs(file)
	if err != nil {
		return nil, err
	}
	f, err := os.Open(p)
	if err != nil {
		return nil, err
	}
	msg, err = mail.ReadMessage(f)
	if err != nil {
		return nil, err
	}
	return mail.ReadMessage(msg.Body)
}

func getEmlMail(file string) (*mail.Message, error) {
	p, err := filepath.Abs(file)
	if err != nil {
		return nil, err
	}
	f, err := os.Open(p)
	if err != nil {
		return nil, err
	}
	return mail.ReadMessage(f)
}

func getDefaultMail() (*mail.Message, error) {
	msg := "Date: Mon, 23 Jun 2015 11:40:36 -0400\n" +
		"From: RobotFrom <From@example.com>\n" +
		"To: RobotTo <To@example.com>\n" +
		"Subject: Gophers at Gophercon\n" +
		"\n" +
		"Message body\n" +
		"\n"

	r := strings.NewReader(msg)
	return mail.ReadMessage(r)
}
