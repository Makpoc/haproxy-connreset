Steps to reproduce
==================

Start HAProxy
-------------

<pre>
cd docker 
(TAG=1.8 ; docker build -t haproxy-eof:${TAG} . && \
docker run -it --rm --name haproxy-syntax-check haproxy-eof:${TAG} haproxy -c -f /usr/local/etc/haproxy/haproxy.cfg && \
docker run -it --rm --network=host --name haproxy-eof haproxy-eof:$TAG)
</pre>

Start the server
----------------

    cd server && go run main.go 2690

Run the client
--------------

### Use the go client

    cd client && go run main.go "http://localhost:8080"

### Use curl
```bash
curl -kv \
  -X PUT \
  -H "Expect:" \
  --data-binary @/tmp/put_body \
  "http://127.0.0.1:8080/" \
  --next \
  -X PATCH -d '{"message": value}' \
  -H "Content-type: application/json" \
  "http://127.0.0.1:8080/"
```
