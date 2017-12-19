# Tweet
Simple function to send tweets, via the Twitter API

## Configuration

```yaml
  tweet:
    lang: go
    handler: ./tweet
    image: nicholasjackson/func_tweet
    environment:
      twitter_consumer_key: xxxxxxxxx
      twitter_consumer_secret: xxxxxxxxxxxx
      twitter_access_token: xxxxxxxx
      twitter_access_token_secret: xxxxxxx

```
All keys and secrets can be found at https://dev.twitter.com/apps/YOUR_APP_ID/show

**twitter_consumer_key** application consumer key  
**twitter_consumer_secret** application consumer secret  
**twitter_access_token** twitter access token  
**twitter_access_token_secret** twitter access token secret  

## Request
Request is a simple json object containing the message to send
```json
{
  "Text": "Test Message",
  "Image": "[image encoded as a base64 string]"
}
```

## Response
Response is a simple json object containing a http status code and detailed error or success message
```json
{
  "Code": 200,
  "Message": "Tweet sent"
}
```

## Testing
You can test the function using the faas CLI

```bash
$ echo '{ "Text": "Testing something else" }' | faas-cli invoke tweet
```

or with an image

```bash
$ echo '{ "Text": "Hey, a tweet with an image", "Image": "$(cat ./mypicture.png | base64)" }' | faas-cli invoke tweet
```
