package pdc_toast

type ToastDuration string

const (
	SHORT_DURATION ToastDuration = "short"
	LONG_DURATION  ToastDuration = "long"
)

type ToastAudio string

const (
	DEFAULT         ToastAudio = "ms-winsoundevent:Notification.Default"
	IM              ToastAudio = "ms-winsoundevent:Notification.IM"
	MAIL            ToastAudio = "ms-winsoundevent:Notification.Mail"
	REMINDER        ToastAudio = "ms-winsoundevent:Notification.Reminder"
	SMS             ToastAudio = "ms-winsoundevent:Notification.SMS"
	LOOPING_ALARM   ToastAudio = "ms-winsoundevent:Notification.Looping.Alarm"
	LOOPING_ALARM2  ToastAudio = "ms-winsoundevent:Notification.Looping.Alarm2"
	LOOPING_ALARM3  ToastAudio = "ms-winsoundevent:Notification.Looping.Alarm3"
	LOOPING_ALARM4  ToastAudio = "ms-winsoundevent:Notification.Looping.Alarm4"
	LOOPING_ALARM5  ToastAudio = "ms-winsoundevent:Notification.Looping.Alarm5"
	LOOPING_ALARM6  ToastAudio = "ms-winsoundevent:Notification.Looping.Alarm6"
	LOOPING_ALARM7  ToastAudio = "ms-winsoundevent:Notification.Looping.Alarm7"
	LOOPING_ALARM8  ToastAudio = "ms-winsoundevent:Notification.Looping.Alarm8"
	LOOPING_ALARM9  ToastAudio = "ms-winsoundevent:Notification.Looping.Alarm9"
	LOOPING_ALARM10 ToastAudio = "ms-winsoundevent:Notification.Looping.Alarm10"
	LOOPING_CALL    ToastAudio = "ms-winsoundevent:Notification.Looping.Call"
	LOOPING_CALL2   ToastAudio = "ms-winsoundevent:Notification.Looping.Call2"
	LOOPING_CALL3   ToastAudio = "ms-winsoundevent:Notification.Looping.Call3"
	LOOPING_CALL4   ToastAudio = "ms-winsoundevent:Notification.Looping.Call4"
	LOOPING_CALL5   ToastAudio = "ms-winsoundevent:Notification.Looping.Call5"
	LOOPING_CALL6   ToastAudio = "ms-winsoundevent:Notification.Looping.Call6"
	LOOPING_CALL7   ToastAudio = "ms-winsoundevent:Notification.Looping.Call7"
	LOOPING_CALL8   ToastAudio = "ms-winsoundevent:Notification.Looping.Call8"
	LOOPING_CALL9   ToastAudio = "ms-winsoundevent:Notification.Looping.Call9"
	LOOPING_CALL10  ToastAudio = "ms-winsoundevent:Notification.Looping.Call10"
	SILENT          ToastAudio = "silent"
)

type HintCrop string

const (
	SQUARE HintCrop = ""
	CIRCLE HintCrop = "circle"
)

type Placement string

const (
	NO_PLACEMENT       Placement = ""
	HERO               Placement = "hero"
	APP_LOGO_OVERRIDES Placement = "appLogoOverride"
)

type ActionPlacement string

const (
	COTEXT_MENU              ActionPlacement = "contextMenu"
	ACTION_PLACEMENT_DEFAULT ActionPlacement = ""
)

type ActionType string

const (
	PROTOCOL   ActionType = "protocol"
	FOREGROUND ActionType = "foreground"
	BACKGROUND ActionType = "background"
)

type MessagePlacement string

const (
	ATTRIBUTE MessagePlacement = "attribution"
)

type ActionButtonStyle string

const (
	SUCCESS  ActionButtonStyle = "Success"
	CRITICAL ActionButtonStyle = "Critical"
)

type ScenarioType string

const (
	URGENT       ScenarioType = "urgent"
	INCOMINGCALL ScenarioType = "incomingCall"
	REMAINDER    ScenarioType = "reminder"
	ALARM        ScenarioType = "alarm"
)
