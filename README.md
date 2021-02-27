# Example Usage

```go

var sender = go_msteams.Sender{webhookUrl: "TEAMS_WEBHOOK_URL"}

var message = go_msteams.NewMessage("My Message")
message.SetColor("#00FF00")


section := message.AddSection("My Section","Section title","http://via.placeholder.com/150x150")

section.AddFact("Due Date","23 octobre 2025")

err := sender.SendMessage(message)
if err != nil {
    t.Error(err)
}


```