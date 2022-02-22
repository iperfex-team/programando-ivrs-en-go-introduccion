# Programando IVRs en Go - Introducción


## Asterisk Conf

iax.conf
```
[general]
mailboxdetail=yes
tos=ef
language=en
disallow=all
allow=ulaw
allow=alaw
allow=gsm

[100]
deny=0.0.0.0/0.0.0.0
secret=12345
transfer=yes
host=dynamic
type=friend
port=4569
qualify=yes
dial=IAX2/100
accountcode=
permit=0.0.0.0/0.0.0.0
requirecalltoken=yes
context=from-internal
secret_origional=8a283a10d3f1768e024d06a2c44489b8
callerid=Test <100>
setvar=REALCALLERIDNUM=
```

extension.conf
```
[general]
static=yes
writeprotect=no
autofallthrough=yes
extenpatternmatchnew=true
clearglobalvars=no
priorityjumping=yes
userscontext=default
AL_OPTIONS = trwW

[default]
exten => s,1,Playback(demo-congrats)
exten => h,1,Hangup()

[from-internal]
include => app-echo-test
include => fastagi

exten => 510,1,NoOp(QUEUE 510 ADMIN)
exten => 510,n,Wait(2)
exten => 510,n,Playback(vm-goodbye)
exten => 510,n,Hangup

exten => 520,1,NoOp(QUEUE 520 SALES)
exten => 520,n,Wait(2)
exten => 520,n,Playback(vm-goodbye)
exten => 520,n,Hangup

exten => 540,1,NoOp(QUEUE 540 SUPPORT)
exten => 540,n,Wait(2)
exten => 540,n,Playback(vm-goodbye)
exten => 540,n,Hangup

[app-echo-test]
exten => *43,1,Answer
exten => *43,n,Wait(1)
exten => *43,n,Playback(demo-echotest)
exten => *43,n,Echo()
exten => *43,n,Playback(demo-echodone)
exten => *43,n,Hangup
exten => h,1,Hangup

[hints]
exten = 100,hint,IAX/100

[fastagi]
exten => 12345,1,NoOp(DOCKER ASTERISK FASTAGI GO)
exten => 12345,n,AGI(agi://fastagi:8000, ivrdemo1,es)
exten => 12345,n,Hangup()
```

## Docker-compose

### RUN
```
docker-compose up -d
```

### PS
```
docker-compose ps
```

### DOWN
```
docker-compose down
```

### CLI ASTERISK
```
docker-compose exec voip asterisk -rvvvvvvvvvv
agi set debug on
```

## Asterisk IAX2 (Use Zoiper client) 

Download: **https://www.zoiper.com/en/voip-softphone/download/current**

user: 100

pass: 12345

port: 4569



