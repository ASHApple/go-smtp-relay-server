# Go-SMTP-Relay-Server
Simple SMTP Relay Server writen with Go

Introduction
------------

I started this project because my [Nextcloud](https://github.com/nextcloud/server) Admin Mail not working on SSL/TLS, STARTLS.

Installation
------------

    $ cp config.example.yml config.yml
    $ vi config.yml OR $ nano config.yml
    $ go get github.com/mhale/smtpd gopkg.in/yaml.v3
    $ go build main.go

Packages
--------

  - [smtpd](https://github.com/mhale/smtpd)
  - [yaml.v3](https://gopkg.in/yaml.v3)

Contributing
------------

Pull requests are always welcome.

License
-------

Copyright (c) 2019, HyeongJong Choi. Released under the MIT [License](LICENSE).