whoami
======

whoami is a Go microservice that returns information over HTTP(S) about the host that it is running on. The primary purpose of this service is to debug and benchmark containerised network services.


Configuration
-------------

    Usage of whoami:
      -addr string
            server listen address (default "127.0.0.1:8080")
      -tls
            enable transport security
      -tls-cert string
            tls certificate (default "cert.pem")
      -tls-key string
            tls private key (default "cert.key")


Example
-------

    $ curl http://hostip:32465
    {"addrs":["::1/128","127.0.0.1/8","172.16.2.3/24"],"hostname:"3ei73ioe84c"}
