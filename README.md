
![Logo](https://github.com/JuanDiegoCastellanos/kafka_golang/blob/main/banner.png)


##
| [![MIT License](https://img.shields.io/badge/License-MIT-green.svg)](https://choosealicense.com/licenses/mit/) | ![GitHub commit activity](https://img.shields.io/github/commit-activity/w/JuanDiegoCastellanos/kafka_golang) |  ![Static Badge](https://img.shields.io/badge/go-main_language-blue) | ![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/JuanDiegoCastellanos/kafka_golang) | ![GitHub Tag](https://img.shields.io/github/v/tag/JuanDiegoCastellanos/kafka_golang) | ![GitHub Release](https://img.shields.io/github/v/release/JuanDiegoCastellanos/kafka_golang)  | ![GitHub Repo stars](https://img.shields.io/github/stars/JuanDiegoCastellanos/kafka_golang?style=flat-square)  |
|--------------|--------------|--------------|--------------|--------------|--------------|--------------|


## Appendix

This a project with golang using kafka for data streaming, using my own built docker images. 


## Features

- Concurrency
- Mutex/ RWLock and RWUnlock
- Multiple persistence drivers
- Highly available services  
- SQLC techniques
- Apache kafka
- Event Driven Architecture
- Docker


## Dependecies
- require
- fake "only for tests"
- testify
- lib/pq
- net
- encoding/json
- gin
- gorilla/websocket
- gorilla/mux
- httptest
- gomock
- confluent-kafka-go.v1/kafka
## Deployment

To deploy this project run

1- First download all the dependecies:
```bash
  go mod download
```
2- Compile the app:
```bash
  go build main.go
```
3- Run the app:
```bash
  ./kafka_golang
```

And always make sure about the dependecies version, sometimes one project might fail because of a deprecated dependency.
## Support

For support, email juandiego.castellanosjerez@gmail.com

