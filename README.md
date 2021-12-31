This is a simple Golang service to send SMS to users using Twilio


Usage:

`curl -X POST -H "Content-Type: application/json" -d '{"to":"your number here","body":"Hello World"}' http://localhost:3006/sms`

example response:
```json
{
    "status": "success",
    "message": "message sent successfully"
    "message sid": "427ca519-fad5-470b-a33b-51774aec29f3"
}
```