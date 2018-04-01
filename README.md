# go-crawl-store
go-crawl-store a restful API server to crawler price for online stores (大潤發, 家樂福)

## Start using it
0. Please note that because of the net/html dependency, goquery requires Go1.1+

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


## How to Qurey:

- **Example** - Compare the Price for 可口可樂 with 大潤發 & 家樂福
```bash
curl http://127.0.0.1:8080/api/v1/ALL/可口可樂
```
<img  src="https://raw.githubusercontent.com/EddieYY/go-crawl-store/master/img/ALL_可口可樂.png">

- **Example** - Qurey the real time price for 可口可樂 with 家樂福
```bash
curl http://127.0.0.1:8080/api/v1/家樂福/可口可樂
```
<img  src="https://raw.githubusercontent.com/EddieYY/go-crawl-store/master/img/家樂福_可口可樂.png">


- **Example** - Qurey the real time price for 可口可樂 with 大潤發
```bash
curl http://127.0.0.1:8080/api/v1/大潤發/可口可樂
```
<img  src="https://raw.githubusercontent.com/EddieYY/go-crawl-store/master/img/大潤發_可口可樂.png">




