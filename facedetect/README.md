# Face Detection with OpenCV
Requires go-opencv template to build

```yaml
  facedetect:
    lang: go-opencv
    handler: ./facedetect
    image: nicholasjackson/func_facedetect
```

## Input
The input type of the function is controlled via a http header and accepts jpeg images, or base64 encoded jpegs.

## Output
The face data is returned as an array of rectangles...

```json
{
  "Faces":[
    {"Min":{"X":193,"Y":91},"Max":{"X":302,"Y":200}},
    {"Min":{"X":459,"Y":107},"Max":{"X":573,"Y":221}}
  ],
  "Bounds":{"Min":{"X":0,"Y":0},"Max":{"X":800,"Y":600}}
}
```

## Testing
The images must be base64 encoded jpegs you can test the function from the command line using the following command

```bash
$ cat pic.jpg | faas-cli --gateway http://192.168.1.113:8080/ invoke facedetect
# or
$ cat pic.jpg | base64 | faas-cli --gateway http://192.168.1.113:8080/ invoke facedetect
```
