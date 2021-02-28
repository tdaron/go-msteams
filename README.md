# Example Usage

```go

var sender = go_msteams.Sender{WebhookUrl: "TEAMS_WEBHOOK_URL"}

var message = go_msteams.NewMessage("My Message")
message.SetColor("#00FF00")


section := message.AddSection("My Section","Section title","http://via.placeholder.com/150x150")

section.AddFact("Due Date","23 octobre 2025")

button = go_msteams.NewActionCard("First action")
button.AddInput(go_msteams.NewDateInput("duedate","Due date",false))
button.AddAction(go_msteams.NewHTTPPostAction("Send","POST_URL","duedate={{duedate.value}}"))

message.AddButton(button)


openuri := go_msteams.NewOpenUri("Little link")
openuri.AddTarget(go_msteams.NewOpenUriTarget("https://google.com","default"))
message.AddButton(openuri)


err := sender.SendMessage(message)
if err != nil {
    panic(err)
}


```
