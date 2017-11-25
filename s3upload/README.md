# S3 Upload
Simple function to upload a file to a S3 bucket

```yaml
s3store:
  lang: go
  handler: ./s3store
  image: nicholasjackson/func_s3store
  environment:
    bucket: "my-s3-bucket"
    accessKeyID: "xxxxxxxxxxxxxxxxx"
    secretAccessKey: "xxxxxxxxxxxxxxxxxxxxxxxxx"
    region: "eu-west-1"

```

## Environment variables
*bucket*: Name for the s3 bucket  
*accessKeyID*: AWS access key to use for storage access  
*secretAccessKey*: AWS secret to use for storage access  
*region*: - AWS region  

## Payload
```json
{
  "filename": "myfile.jpg",
  "permissions": "public-read",
  "data_base64": "abc2131="
}
```

### Valid values for permissions
private | public-read | public-read-write | aws-exec-read | authenticated-read | bucket-owner-read | bucket-owner-full-control
