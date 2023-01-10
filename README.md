# tg-alarm-button

- Create a bot: https://t.me/botfather;
- Start a chat with the bot; 
- Visit `https://api.telegram.org/botACCESS_TOKEN/getUpdates` and find the chat id; 
- Put your chat id and access token into the `cred.go` file;
- Build an app: `go build -tags cred`

Sample message:
```
🔥 THE ALARM BUTTON WAS PUSHED 🔥

🖥 EVILCORP-NET
👤 EVILCORP-NET\John.Doe
🔌 192.168.111.8/24, ::1/128, 127.0.0.1/8, fe80::5efe:a0a:b2b/128
```