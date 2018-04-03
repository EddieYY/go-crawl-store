# go-crawl-store
go-crawl-store a restful API server to crawler price for online stores (大潤發, 家樂福)

## Start using it

1. Download and install it:

```sh
 $ go get github.com/EddieYY/go-crawl-store
```
2. To run (test):

```sh
 $ cd $GOPATH/src/github.com/EiddieYY/go-crawl-store
 $ go run main.go

```

<img  src="https://raw.githubusercontent.com/EddieYY/go-crawl-store/master/img/go-crawl-stire_server_run.png">


## How to Query:

- **Example** - Compare the price for 可口可樂 with 大潤發 & 家樂福
```bash
curl http://127.0.0.1:8080/api/v1/ALL/可口可樂
```
<img  src="https://raw.githubusercontent.com/EddieYY/go-crawl-store/master/img/ALL_可口可樂.png">

- **Example** - Query the real time price for 可口可樂 with 家樂福
```bash
curl http://127.0.0.1:8080/api/v1/家樂福/可口可樂
```
<img  src="https://raw.githubusercontent.com/EddieYY/go-crawl-store/master/img/%E5%AE%B6%E6%A8%82%E8%A4%94_%E5%8F%AF%E5%8F%A3%E5%8F%AF%E6%A8%82.png">


- **Example** - Query the real time price for 可口可樂 with 大潤發
```bash
curl http://127.0.0.1:8080/api/v1/大潤發/可口可樂
```
<img  src="https://raw.githubusercontent.com/EddieYY/go-crawl-store/master/img/大潤發_可口可樂.png">




