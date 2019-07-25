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

    cd client && go run main.go "http://localhost:8080"
