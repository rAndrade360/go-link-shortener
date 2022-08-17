# api

### Request

```bash
#GRPC call
$ grpcurl -plaintext -d '{"url":"http://google.com/gmail"}' \
localhost:8081 ShortUrl/Short

#HTTP call
$ curl -i -XPOST localhost:8080 -d '{"url": "http://google.com/iaiaoo"}'
```