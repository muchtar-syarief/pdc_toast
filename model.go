package pdc_toast

type Action struct {
	Label       string
	Arguments   string
	ImageUri    string
	Type        ActionType
	Placement   ActionPlacement
	ButtonStyle ActionButtonStyle
	Tooltip     string
}

type Icon struct {
	Src       string
	HintCrop  HintCrop
	Placement Placement
}

type Message struct {
	AlignCenter bool
	Message     string
	Placement   MessagePlacement
}
