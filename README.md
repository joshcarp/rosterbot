# rosterbot

<html>
<a href="https://slack.com/oauth/v2/authorize?scope=incoming-webhook,commands,chat:write&client_id=1367393582980.1445120201280"><img alt=""Add to Slack"" height="40" width="139" src="https://platform.slack-edge.com/img/add_to_slack.png" srcset="https://platform.slack-edge.com/img/add_to_slack.png 1x, https://platform.slack-edge.com/img/add_to_slack@2x.png 2x" /></a>
</html>

- A slack app for rostering

## Use

- Roster a message every day at 23:00 UTC
```
/roster add "0 23 * * *" "message" @user1 @user2 @user3 

```

- Remove rosters for current channel 
```
/roster remove
```
