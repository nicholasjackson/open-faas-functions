# OpenFaaS Functions

## Echo
Simple function which echos the data posted to it.  
[s3upload/README.md](s3upload/README.md)

```yaml
  echo:
    lang: go
    handler: ./echo
    image: nicholasjackson/func_echo
```

## S3 Upload
Upload a file to a s3 bucket.  
[s3upload/README.md](s3upload/README.md)

```yaml
  s3upload:
    lang: go
    handler: ./s3upload
    image: nicholasjackson/func_s3upload
    environment:
      bucket: "mybucket.name"
      accessKeyID: "xxxxxxxxxxxx"
      secretAccessKey: "xxxxxxxxxx"
      region: "eu-west-1"
```

## JWT Authorization - Experimental
Validate JWT tokens with their corresponding public key and provide the claims to a function as environment variables
[jwt/README.md](jwt/README.md)

```yaml
  jwt:
    lang: gojwt
    handler: ./jwt
    image: nicholasjackson/func_jwt
    environment:
      # enable jwt authorization for this function
      jwt: true
      # base64 public key to use for validating jwt signature
      jwt_public_key: "LS0tLS1CRUdJTiBQVUJMSUMgS0VZLS0tLS0KTUlJQ0lqQU5CZ2txaGtpRzl3MEJBUUVGQUFPQ0FnOEFNSUlDQ2dLQ0FnRUFuV3E1OFRXYUNlNkx4cFFwVXNwMgpINit2ekpnZnkwcGl0Zk1JUGpDZklTVEQxc1BGRUM3ZG9zVTlRZjRlamF3WGxKTDd2Kzc4aHBEczlFZHdIQW1ZCmU4bENSMS9WVC9XZElROHdTdVd6MXNrMkt0ZFZPYVE2bk9HZytCd0NnVWcvRkRYcHEvQ2FHb0NLOUF4RUh0aGEKL0xMYWR5TW1HelZUbEJzZ0M5TDY2S1VlR0E5bGRxUmpKVGtsYnMxQVB0VVNsaXpqZTdqMXNlSXdTcWhFSDNrUAowN1U4cTlRZFI1bGNtakdITWVlOFcwWitaWG90U0JFUzBwWThTMVZTKzJRQ0wzenp4dlkvaG9NL080NEo5TENnCnF1cWxqL1Z6bTBsb1NHQXNCVmRmOGlLeSt2eWN2eWRrRVpxMzZtRWFMNmtnYk5OL24zZXhiVGJKVUZHVVY2L0EKQW44Nmhhc0k2dTZCTVdsODRHdFNQeEZsM3hkTGN5WUF3TUVVTGY4KzIyQWE4NFZzMjB1dUtSZExOOU9IWnFtRgpWUkJ5cnM0bEZDdWR3bU1XSmk3SjQ4N3ZGWUJITG1zMG1XZGQxK3ZjaUtHQXhLOWNKY25IOUt6UiswdnVFUkUwCjVNdG93WjNRUU5EVk5Ham1TYlJrSVJOMFdsK3VHV29XOUVjSWh3U3BCc1p1MTRoZDlvVThYa0lia1RXVHRyb0YKck5WK2pJdkh4TjE2K0dOWHZKaVptYi9KanVmQnRGUllPaDR0VXo3c08rZ1JUSGFWU0tES0FHRDdiL2MyQVM1bwo5dUVqTHdZaGZhbEJhWHZTZVJ2NW9UeWJZV1M2ZVBrWFZ2enZsVnZBRmhIKzVpSGhwdVFLVm13RnYrKzlEQk9VCkRKMVJXaGJRSnBhN3Y2SHU4S0d0aDlFQ0F3RUFBUT09Ci0tLS0tRU5EIFBVQkxJQyBLRVktLS0tLQo="
```

## Tweet
Simple function to send a tweet using the twitter API
[tweet/README.md](tweet/README.md)

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
