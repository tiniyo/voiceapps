# Voice applications

golang tiniyo voice call apps. this voice application has direct call and restaurent voice ivr.

# Build docker file.

```bash
$ docker build -t voiceapps:latest .
```
# Run docker file on port 8080

```bash
$ docker run -p 8080:8080 voiceapps:latest
```

# Install and configure ngrok
## Install snapd & ngrok
```bash
sudo apt update
sudo apt install snapd
sudo snap install ngrok
```
## Run ngrok for 8080 port
```bash
$ ngrok http 8080
ngrok by @inconshreveable                                                                                                                                                    
Session Status                online
Session Expires               1 hour, 49 minutes
Version                       2.3.35
Region                        United States (us)
Web Interface                 http://127.0.0.1:4040
Forwarding                    http://04f88df36fbf.ngrok.io -> http://localhost:8080
Forwarding                    **https://04f88df36fbf.ngrok.io** -> http://localhost:8080
                                                                                                   
Connections                   ttl     opn     rt1     rt5     p50     p90                     
                              0       0       0.00    0.00    0.00    0.00
```


# Configure application on Tiniyo platform. 
use public **https://04f88df36fbf.ngrok.io** url from ngrok and configure application on tiniyo for the sip endpoint to test.
