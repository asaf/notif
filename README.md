# notif

_notif_ is a standalone binary that sends notifications.

It currently supports SMTP only.


# why ?

Because sending HTML emails via Linux CLI requires too much effort and we needed a way to simply pipe HTML content into an emails.

# how to use

```bash
echo "<b>Nice Body</b>" | ./notif smtp --port 1025 -from from@me.com -to to@you.com -subject "Nice Subject"
```

