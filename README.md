# voiceapps

golang tiniyo voice call apps. this voice application has direct call and restaurent voice ivr.

# build docker file.

```bash
$ docker build -t voiceapps:latest .
```
# run docker file on port 8080

```bash
$ docker run -p 8080:8080 voiceapps:latest
```

# install and configure ngrok
## install snapd
```bash
sudo apt update
sudo apt install snapd
sudo snap install ngrok
```
## Run ngrok for 8080 port
```bash
$ ngrok http 8080
```

# use public https:// url from ngrok and configure application on tiniyo for the sip endpoint to test. 
