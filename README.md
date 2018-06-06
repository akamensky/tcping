# Ping TCP port with tcping

tcping is a tool that allows to verify reachability of TCP port.
The purpose of this tool is to allow for a monitoring of reachability 
of specific service on remote (or local) host.
Unlike standard ping, fping or hping, this tool uses simple TCP 
connection to verify if port is listening and is agnostic of application 
protocols etc.

Each check will simply open a connection and if successful it will 
immediately close it. While some applications may not like this, most 
will handle this will and will not cause any significant resource 
consumption.

####  Get it

Currently: to start using this you will need Go language compiler, 
which you can get from [golang.org](https://golang.org/).

Compilation:

```sh
# go get -u -v -x github.com/akamensky/tcping
```

Installation:
```sh
# mv ~/go/bin/tcping /usr/local/bin/tcping
```

#### Use it

Simple use case:
```
# tcping -s www.github.com -p 443
```

For more detailed usage information check included help message:
```
# tcping -h
```

#### Know it

Currently `tcping` does DNS resolution for every check. This allows 
to test systems that are spread across multiple servers and see it 
more from user perspective.

#### License

See included LICENSE file for license information.  