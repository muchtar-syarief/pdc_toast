package pdc_toast

import "text/template"

func ToasTemplate() *template.Template {
	toastTemplate := template.New("toast")
	toastTemplate.Parse(`
[Windows.UI.Notifications.ToastNotificationManager, Windows.UI.Notifications, ContentType = WindowsRuntime] | Out-Null
[Windows.UI.Notifications.ToastNotification, Windows.UI.Notifications, ContentType = WindowsRuntime] | Out-Null
[Windows.Data.Xml.Dom.XmlDocument, Windows.Data.Xml.Dom.XmlDocument, ContentType = WindowsRuntime] | Out-Null

$APP_ID = '{{if .AppID}}{{.AppID}}{{else}}Windows App{{end}}'

$template = @"
<toast 
    {{if .Scenario}}
    scenario="{{ .Scenario }}" 
    {{end}}

    {{if .ActivationType}}
    activationType="{{.ActivationType}}" 
    {{end}}

    {{if .ActivationArguments}}
    launch="{{.ActivationArguments}}" 
    {{end}}

    {{if .Duration}}
    duration="{{.Duration}}" 
    {{end}}

    {{if .UseButtonStyle}}
    useButtonStyle='{{ .UseButtonStyle}}'
    {{end}}
    >
    <visual>
        <binding template="ToastGeneric">
            {{if .Title}}
            <text><![CDATA[{{.Title}}]]></text>
            {{end}}

            {{range .Messages}}
            <text 
            hint-callScenarioCenterAlign="{{ .AlignCenter }}"
            ><![CDATA[{{.Message}}]]></text>
            {{end}}

            {{if .Icon}}
            <image 
            placement="{{ .Icon.Placement }}" 
            hint-crop="{{ .Icon.HintCrop }}"
            src="{{.Icon.Src}} " 
            />
            {{end}}
        </binding>
    </visual>
    {{if ne .Audio "silent"}}
	<audio src="{{ .Audio }}" loop="{{ .Loop }}" />
	{{else}}
	<audio silent="true" />
	{{end}}
    {{if .Actions}}
    <actions>
        {{range .Actions}}
        <action 

        {{if .Type}}
        activationType="{{.Type}}" 
        {{end}}

        {{if .Label}}
        content="{{.Label}}" 
        {{end}}

        {{if .Arguments}}
        arguments="{{.Arguments}}" 
        {{end}}

        {{if .ImageUri}}
        imageUri="{{ .ImageUri }}" 
        {{end}}

        {{if .ButtonStyle}}
        hint-buttonStyle="{{ .ButtonStyle }}" 
        {{end}}

        {{if .Tooltip}}
        hint-toolTip="{{ .Tooltip }}"
        {{end}}
        />
        {{end}}
    </actions>
    {{end}}
</toast>
"@

$xml = New-Object Windows.Data.Xml.Dom.XmlDocument
$xml.LoadXml($template)
$toast = New-Object Windows.UI.Notifications.ToastNotification $xml
[Windows.UI.Notifications.ToastNotificationManager]::CreateToastNotifier($APP_ID).Show($toast)
    `)

	return toastTemplate
}
