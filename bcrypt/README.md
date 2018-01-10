# Bcrypt
Simple function to bcrypt and validate bcrypt hashes

## Configuration

```yaml
  bcrypt:
    lang: go
    handler: ./bcrypt
    image: nicholasjackson/func_bcrypt
```

## Hashing a string
Simply call the function with the string you would like to has as a string to the request

```
$ echo "password" | faas-cli invoke bcrypt
$2a$10$W879WXMlNfuN3dkcs9CJA.tTB3bLGsqMfS6DWOoBlDx7GJXv0rld6
```

### Response
Hashed version of the string

## Comparing a string and a hashed string
To compare a string to its hash call the function passing the hash and the password separated by a single space to the request.  
`"[hash] [string]"`

```
echo "$2a$10$W879WXMlNfuN3dkcs9CJA.tTB3bLGsqMfS6DWOoBlDx7GJXv0rld6 password" | faas-cli invoke --query action=validate bcyrpt
ok
```

### Response
**ok** - The hash and the string match  
**err: not equal** - The hash and the string do not match
