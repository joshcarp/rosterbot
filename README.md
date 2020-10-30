# rosterbot

<html>
<a href="https://slack.com/oauth/v2/authorize?scope=incoming-webhook,commands,chat:write&client_id=1367393582980.1445120201280"><img alt=""Add to Slack"" height="40" width="139" src="https://platform.slack-edge.com/img/add_to_slack.png" srcset="https://platform.slack-edge.com/img/add_to_slack.png 1x, https://platform.slack-edge.com/img/add_to_slack@2x.png 2x" /></a>
</html>

- A slack app for rostering

## Use

- Roster a message every day at 23:00 UTC
```
/roster add "0 23 * * *" "message" @user1, @user2, @user3 
// 23:00 UTC: message @user1
// 23:00 Tomorrow: message @user2
// 23:00 The day after tomorrow: message @user2
```


- Remove rosters for current channel 
```
/roster remove
```
## Caveats

- Doesn't support full cron syntax:
```
* * * * * is supported
0 /10 * * is not supported
```

- All time must be specified in UTC time, not local time

## About
- Hosted on google cloud functions
- Firestore to store data (webhooks and subscriptions)
- Uses google cloud scheduling to execute every minute (* * * * *) in order to filter out firestore subscriptions for that time