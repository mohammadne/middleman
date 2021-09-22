<!-- LOGO -->
<p align="center">
  <img src="assets/logo.png" />
</p>

<!-- BADGES -->
<p align="center">
  <img src="https://img.shields.io/github/release/mohammadne/middleman.svg?style=for-the-badge">
  <img src="https://img.shields.io/codecov/c/gh/mohammadne/middleman?logo=codecov&style=for-the-badge">
  <img src="https://img.shields.io/github/license/mohammadne/middleman?style=for-the-badge">
  <img src="https://img.shields.io/github/stars/mohammadne/middleman?style=for-the-badge">
  <img src="https://img.shields.io/github/downloads/mohammadne/middleman/total.svg?style=for-the-badge">
</p>

<!-- TITLE -->
# MIDDLE MAN
> `MIDDLE MAN` will simulate a real world proxy server that clients
> will request to the proxy server (reverse proxy) and as its role is
> load-balancing between servers and caching some requests based on
> logic in proxy server

## How to run

- run servers in `first` terminal

  ``` zsh
  go run cmd/root.go servers
  ```

- run proxy in `second` terminal

  ``` zsh
  go run cmd/root.go proxy
  ```

- run clients in `third` terminal

  ``` zsh
  go run cmd/root.go clients
  ```

- post a test `cURL` body

  ``` zsh
  curl --header "Content-Type: application/json" \
  --request POST --data '{"key":"key","value":"value"}' \
  http://localhost:8090/objects
  ```

- get specific hashId's body with `cURL`

  ``` zsh
  curl --header "Content-Type: application/json" \
  --request GET http://localhost:8090/objects/hashId
  ```
