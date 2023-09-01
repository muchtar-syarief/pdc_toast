package main

func main() {
	n := Notification{
		AppID: "Test",
		Title: "Ini Title",
		Messages: []Message{
			{
				Message: "hahaha",
			},
			{
				Message: "jiahahahha",
			},
		},
		ActivationType: string(PROTOCOL),
		Audio:          DEFAULT,
		Duration:       SHORT_DURATION,
		Actions: []Action{
			{
				Label:       "Halo",
				Arguments:   "callId=123",
				Type:        PROTOCOL,
				ButtonStyle: SUCCESS,
			},
			{
				Label:       "Close",
				Arguments:   "callId=123",
				Type:        BACKGROUND,
				ButtonStyle: CRITICAL,
			},
		},
	}

	n.Push()
}
