It's crawler website using Go language.

## Features

- Go language
- Docker
- Elastic Search
- MVC pattern
- Microservices
- Singleton -> Concurrent -> Distribute

## Installation and go packages

- go language
- docker
- elasticsearch
- go get golang.org/x/text
- go get -v github.com/gpmgo/gopm
- gopm get -g -v golang.org/x/text
- gopm get -g -v golang.org/x/net/html
- go get gopkg.in/olivere/elastic.v6

## Algorithm

![Algorithm](./images/Algorithm.png)

## Framework

![Frame1](.\images\Frame1.png)

![Frame2](./images/Frame2.png)

## Architecture

![Architecture](.\images\Architecture.png)



## Usage for Concurrent

- Start Docker.
- Run Script "docker run -d -p 9200:9200 elasticsearch:6.8.23"
- Run "crawler/main.go", to start the singleton crawler.
- Run "crawler/frontend/starter.go", to view the result in the website.
- Visit "http://localhost:8888/" in your browser
- Type in query string with REST format. such as "女 && Age>20"

## Usage for Distribute

- Start Docker.
- Run Script "docker run -d -p 9200:9200 elasticsearch"
- Open a Terminal, execute: crawler_distributed\persist\server>go run ItemSaver.go --port=1234
- Open a Terminal, execute: crawler_distributed\worker\server>go run worker.go --port=9000
- Open a Terminal, execute: crawler_distributed\worker\server>go run worker.go --port=9001
- Open a Terminal, execute: crawler_distributed>go run main.go --itemsaver_host=":1234" --worker_hosts=":9000,:9001"
- Run "crawler/frontend/starter.go", to view the result in the website.
- Visit "http://localhost:8888/" in your browser
- Type in query string with REST format. such as "男 && 已购车"