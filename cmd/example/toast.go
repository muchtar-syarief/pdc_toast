package main

import "github.com/muchtar-syarief/pdc_toast"

func main() {
	var err error

	toast := pdc_toast.NewNotification("PDC Shop", "New Order.!")
	toast.SetMessages([]pdc_toast.Message{
		{
			AlignCenter: true,
			Message:     "Hei you..!!!",
		}, {
			AlignCenter: true,
			Message:     "You have new order. Go packaging and send to delivery service.",
			Placement:   pdc_toast.ATTRIBUTE,
		},
	})

	err = toast.SetIcon(pdc_toast.Icon{
		Src:       "./icons/order.png",
		HintCrop:  pdc_toast.SQUARE,
		Placement: pdc_toast.APP_LOGO_OVERRIDES,
	})
	if err != nil {
		panic(err)
	}

	toast.SetAction(true, []pdc_toast.Action{
		{
			Label:       "Reply",
			Arguments:   "action=2",
			Type:        pdc_toast.PROTOCOL,
			ButtonStyle: pdc_toast.CRITICAL,
			ImageUri:    "./icons/chat.png",
		},
	})

	err = toast.Push()
	if err != nil {
		panic(err)
	}
}
