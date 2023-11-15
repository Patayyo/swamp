# usercart - simple api with items

<h2>Requirements</h2>
1) Go 1.21
2) make
3) Docker

<h2>How to install</h2>

Execute in the repository root directory commands
1) `make docker-up`
2) `make migration`
3) `make run`
4) Then check <a href="http://localhost:8080/api/v1/healthcheck">localhost healthcheck</a> or <a href="http://localhost:8080/api/v1/items">check items</a>