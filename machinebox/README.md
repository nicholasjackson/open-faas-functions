# Face detection with MachineBox
This function interfaces with machinebox and to perform facial recognition and outputs a copy of the inputted image with borders drawn round recognised faces.
To use this function you must have Facebox by [machinebox.io](https://machinebox.io) running, instructions for running machinebox can be found on the machinebox site.

```yaml
  machinebox:
    lang: go
    handler: ./machinebox
    image: nicholasjackson/machinebox
    environment:
      machinebox_url: http://localhost:8080
```

## Example
To detect faces, simply send an image to the faas cli and invoke the machinebox function
```bash
cat hashi.jpg | faas invoke machinebox > out.png
```
