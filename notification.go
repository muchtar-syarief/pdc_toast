package main

import (
	"bytes"
	"fmt"
)

type Notification struct {
	AppID               string
	Title               string
	DisplayTimestamp    string
	Messages            []Message
	Icon                Icon
	ActivationArguments string
	Actions             []Action
	Loop                bool
	ActivationType      string        // default "protocol"
	Audio               ToastAudio    // default "DEFAULT"
	Duration            ToastDuration // default "short"
}

func (n *Notification) Push() error {
	xml, err := n.buildXML()
	if err != nil {
		return err
	}

	fmt.Println(xml)
	return invokeTemporaryScript(xml)
}

func (n *Notification) buildXML() (string, error) {
	var out bytes.Buffer
	err := toastTemplate.Execute(&out, n)
	if err != nil {
		return "", err
	}

	return out.String(), nil
}
