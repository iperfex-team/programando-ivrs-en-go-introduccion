# Programando-IVRs-en-Golang

## Build

### Asterisk
```
cd asterisk-iax2
make
```

### FastAGI DEV
```
cd fastagi
make dev
```

### FastAGI PROD
```
cd fastagi
make
```

## Build Multi Architecture (Production)
```
docker run --privileged --rm tonistiigi/binfmt --install arm64,riscv64,arm
docker buildx create --name multiarchitecture
docker buildx use multiarchitecture 


#asterisk-iax2
cd asterisk-iax2
docker buildx build --platform linux/amd64,linux/arm64 -t cnsoluciones/asterisk-iax2:latest --push .
cd ..
    
#fastagi
cd fastagi
docker buildx build --platform linux/amd64,linux/arm64 -t cnsoluciones/fastagi:latest --push .
cd ..
```
