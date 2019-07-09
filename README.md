# docker-nginx-proxy
Nginx reverse proxy and load balancing to dockerized Go application.

## Installation
1. Install Docker and Docker Compose (https://docs.docker.com/install/linux/docker-ce/ubuntu).
2. Clone Github repository to you machine: `git clone https://github.com/devops787/docker-nginx-proxy.git`
3. Build Docker images: `docker-compose build`
4. Run Docker containers: `docker-compose up -d`

After successful installation process you should get 4 running containers:
```
$ docker ps
CONTAINER ID        IMAGE                          COMMAND                  CREATED             STATUS              PORTS                                      NAMES
61c293e9836a        devops787/docker-nginx-proxy   "./app"                  8 minutes ago       Up 8 minutes        3000/tcp                                   dockernginxproxy_server3_1
06107c09fa85        devops787/docker-nginx-proxy   "./app"                  8 minutes ago       Up 8 minutes        3000/tcp                                   dockernginxproxy_server1_1
f9dcd3607740        nginx                          "nginx -g 'daemon of…"   8 minutes ago       Up 8 minutes        0.0.0.0:80->80/tcp, 0.0.0.0:443->443/tcp   dockernginxproxy_nginx_1
a8e80cdf3068        devops787/docker-nginx-proxy   "./app"                  8 minutes ago       Up 8 minutes        3000/tcp                                   dockernginxproxy_server2_1
```

Containers:
1. `devops787/docker-nginx-proxy` - Go application 
2. `nginx` - Nginx 

The ports are mapped to localhost machine, so you can access them via 127.0.0.1 address.

### Example
```
$ curl -i http://127.0.0.1

HTTP/1.1 200 OK
Server: nginx/1.17.1
Date: Tue, 09 Jul 2019 17:56:19 GMT
Content-Type: text/plain; charset=utf-8
Content-Length: 8
Connection: keep-alive

Server 2
```