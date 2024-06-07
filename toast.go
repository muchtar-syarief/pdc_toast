package pdc_toast

import (
	"bytes"
	"os"
	"os/exec"
	"path/filepath"
	"syscall"

	"github.com/gofrs/uuid"
)

type Notification interface {
	Push() error
	SetIcon(icon Icon) error
	SetMessages(msgs []Message) Notification
	SetAction(useButtonStyle bool, actions []Action) error
}

type NotificationImpl struct {
	AppID               string
	Title               string
	Scenario            ScenarioType
	DisplayTimestamp    string
	Messages            []Message
	Icon                Icon
	ActivationArguments string
	Actions             []Action
	Loop                bool
	UseButtonStyle      bool
	ActivationType      ActionType    // default "protocol"
	Audio               ToastAudio    // default "DEFAULT"
	Duration            ToastDuration // default "short"
}

func (n *NotificationImpl) SetMessages(msgs []Message) Notification {
	n.Messages = msgs
	return n
}

func (n *NotificationImpl) SetIcon(icon Icon) error {
	absfile, err := filepath.Abs(icon.Src)
	if err != nil {
		return err
	}
	icon.Src = absfile
	n.Icon = icon
	return nil
}

func (n *NotificationImpl) SetAction(useButtonStyle bool, actions []Action) error {
	n.UseButtonStyle = useButtonStyle

	for _, item := range actions {
		absfile, err := filepath.Abs(item.ImageUri)
		if err != nil {
			return err
		}
		item.ImageUri = absfile
		n.Actions = append(n.Actions, item)
	}

	return nil
}

func (n *NotificationImpl) Push() error {
	xml, err := n.buildXML()
	if err != nil {
		return err
	}

	return n.invokeTemporaryScript(xml)
}

func (n *NotificationImpl) buildXML() (string, error) {
	var out bytes.Buffer

	template := ToasTemplate()
	err := template.Execute(&out, n)
	if err != nil {
		return "", err
	}

	return out.String(), nil
}

func (n *NotificationImpl) invokeTemporaryScript(content string) error {
	id, _ := uuid.NewV4()
	file := filepath.Join(os.TempDir(), id.String()+".ps1")
	defer os.Remove(file)

	bomUtf8 := []byte{0xEF, 0xBB, 0xBF}
	out := append(bomUtf8, []byte(content)...)
	err := os.WriteFile(file, out, 0600)
	if err != nil {
		return err
	}
	cmd := exec.Command("PowerShell", "-ExecutionPolicy", "Bypass", "-File", file)
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	if err = cmd.Run(); err != nil {
		return err
	}
	return nil
}

func NewNotification(appID, title string) Notification {
	return &NotificationImpl{
		AppID:          appID,
		Title:          title,
		ActivationType: PROTOCOL,
		Duration:       SHORT_DURATION,
		Audio:          DEFAULT,
	}
}
