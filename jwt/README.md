# JWT - Experimental
This simple function uses an experimental version of the watchdog which validates the signature for a provided JWT token.  The use case for this is to enabled fine-grained authorization checks within your functions.  Authentication and the creation of the JWT is handled externally, potentially is is possible to create a function for JWT authentication or you could use a service like [http://auth0.com](http://auth0.com).

For more info on JWT please see: [https://jwt.io](https://jwt.io)

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

## Environment variables
**jwt**: Set to true to enable JWT authorization for this function
**jwt_public_key**: Base64 encoded RSA public key to be used to validate the JWT signature 

## Testing
Once deployed the function can be tested by adding the `X-Authorization-JWT` header containing a valid JWT, the example below has been signed with the private key corresponding to the public key in the settings above.

```bash
$ curl -v -H "X-Authorization-JWT: eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2Nlc3NMZXZlbCI6InVzZXIiLCJleHAiOjE1MTE5NDU4ODMsInVzZXJJRCI6ImFiY3NkMjMyamZqZiJ9.V_hU9pUuAB54UzOP6DATemD1XZq-sMDn5pRMXg-Jekwf5mFy5CYrG4XiF-c01qYJLww7XuIaAycHIr6xIwlGv9jDacbn9wQ96ncxpDnpLy7eCFu-nZ9GpWdOc4Mt4LV6oTuYDYqrXfUD7W-cnE0l1X6_Js9bGQHeGPPhX4PnP9qRkqgIOyvyym25iVUT91CY7gf0a62oRuIqA4Rjim7HIENWOjLJp1jNhrfW2Cc7DnO0CMGVCFJ1XEanrM31S6OSfKWo8zrA3BgMHBahX2S22GZi3mwOv43QpIDCRZqTp9wPG08i1gEGkk-lYo1nGoYUL0YXqAlFlve6oO8hsBe6g1SQg5RJU8YOacro6gwBRwiNSdcQA_KYiy766qVRBFywFyWFLUkP5Yge8iElNhAKcxULD5ziCWHm_cMEKg1YnvdbQsBqGT3kTbIn6F3_q5aC3ieLdM2R0pEfKX0MZlA0I7h7ocL7Pw-qY_iYf40YLvI68M2ITv_C9_zvtPM0sqhq_PBmqLbRtJxKcO6nOT3edjNqBiSQ-ssOEBrLWRw-wG0CkCVS3ew7Q-VH8i8MXR7Psf7aeqKvDZtqudw56SsZGGpCQd0dWpLaKoKSnJzmoy4U9SiNTnJiusOGhTYzoDW-Ti-5twrECZLeSbo8iPHBDAfdqiPuP1kLsFQxr_hMU6k" http://localhost:8080/function/jwt
```

If the JWT is valid then the claims encoded into it will be provided to your function as environment variables prefixed with `claim_`.

```
Claims:
claim_userID=abcsd232jfjf
claim_accessLevel=user
claim_exp=%!s(float64=1.511945883e+09)
```

If the JWT is not validated with the public key the a `401 Unauthorized` status code will be returned.
