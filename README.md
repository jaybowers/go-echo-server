# Quick Start

Start the server with:

```
mkdir -p ~/go && \
    cd ~/go && \
    mkdir -p src/forgerock.com && \
    git clone git@github.com:jaybowers/go-echo-server.git src/forgerock.com/go-echo-server && \
    go build forgerock.com/go-echo-server && \
    ./go-echo-server
```

Then test with:

```
curl -H'X-Echo: world' -d woot 'localhost:8080/hello?there'
```
