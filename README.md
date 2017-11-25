# OpenFaaS Functions

## Echo
Simple function which echos the data posted to it
[s3upload/README.md](s3upload/README.md)

```yaml
  echo:
    lang: go
    handler: ./echo
    image: nicholasjackson/func_echo
```

## S3 Upload
Upload a file to a s3 bucket  
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
