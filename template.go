package main

import "text/template"

var toastTemplate *template.Template

func init() {
	toastTemplate = template.New("toast")
	toastTemplate.Parse(`
[Windows.UI.Notifications.ToastNotificationManager, Windows.UI.Notifications, ContentType = WindowsRuntime] | Out-Null
[Windows.UI.Notifications.ToastNotification, Windows.UI.Notifications, ContentType = WindowsRuntime] | Out-Null
[Windows.Data.Xml.Dom.XmlDocument, Windows.Data.Xml.Dom.XmlDocument, ContentType = WindowsRuntime] | Out-Null

$APP_ID = '{{if .AppID}}{{.AppID}}{{else}}Windows App{{end}}'

$template = @"
<toast activationType="{{.ActivationType}}" launch="{{.ActivationArguments}}" duration="{{.Duration}}">
    <visual>
        <binding template="ToastGeneric" displayTimestamp="{{ .DisplayTimestamp}}">
            {{if .Icon}}
            <image placement="{{ .Icon.Placement }}" hint-crop="{{ .Icon.HintCrop }}" src="{{.Icon.Src}} " />
            {{end}}
            {{if .Title}}
            <text><![CDATA[{{.Title}}]]></text>
            {{end}}
            {{range .Messages}}
            <text hint-callScenarioCenterAlign="{{ .AlignCenter }}"><![CDATA[{{.Message}}]]></text>
            {{end}}
        </binding>
    </visual>
    {{if ne .Audio "silent"}}
	<audio src="{{.Audio}}" loop="{{.Loop}}" />
	{{else}}
	<audio silent="true" />
	{{end}}
    {{if .Actions}}
    <actions>
        {{range .Actions}}
        <action activationType="{{.Type}}" content="{{.Label}}" arguments="{{.Arguments}}" imageUri="{{ .ImageUri }}" hint-buttonStyle="{{ .ButtonStyle }}" hint-toolTip="{{ .Tooltip }}"/>
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
}

// <group>
//         <subgroup>
//           <text hint-style="base">52 attendees</text>
//           <text hint-style="captionSubtle">23 minute drive</text>
//         </subgroup>
//         <subgroup>
//           <text hint-style="captionSubtle" hint-align="right">1 Microsoft Way</text>
//           <text hint-style="captionSubtle" hint-align="right">Bellevue, WA 98008</text>
//         </subgroup>
//       </group>
