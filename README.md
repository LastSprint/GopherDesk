# GopherDesk
Help Desk built on top of Trello with slack integration

1. Create a ticket on specific trello Board
2. Read specific ticket (by id) from trello board
3. If ticket was updated send notification to user who created the ticket

# Ticket creation

1. User opened the form and server got webhook class and send the form back to user
2. User fill a form in slack and send it (by pressing "send")
3. Server get form result and validate it
4. Server load user info (like E-mail, Name and other staff)
5. Server creates a ticket on Trello board

# Ticket Reading

1. User can check status of ticket by typing a command in slack
2. Server get webhook call and return ticket info (status and other)

# Ticket Update

Trello has awesome webhooks API where you can register event which can listen updates of this model (on which you registered the hook)

```
POST https://api.trello.com/1/tokens/{TOKEN}/webhooks/
Content-Type: application/json

{
  "key": "{KEY}",
  "callbackURL": "{URL}",
  "idModel": "{MODELS}",
  "description": "Board hook"
}
```

So when something is changed - our server notify users about

