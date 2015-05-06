# Overview
Very simple web service that keeps track of an internal current version code (build number) for a given app. Provides an easy-to-use rest API that allows:
* GET /versionCode/&lt;app-id&gt; - Current version code for a given app
* GET /versionCode/&lt;app-id&gt;/next - Increment, then retrieve the current version code for a given app
* DELETE /versionCode/&lt;app-id&gt; - Removes app data from the database
* PUT /versionCode/&lt;app-id&gt;/&lt;version-code&gt; - Sets the app version code with the given value
 
# Installation

First make sure you have a working [Go installation] and the [GOPATH] environment variable is setted up.
You must also have [MongoDB installed] and its daemon ([mongod]) up and running.

Then to get and build this project in the terminal type:
```sh
$ go get github.com/mindera/version-code-service
$ go install github.com/mindera/version-code-service
```
The generated binary file can be found at:
```
$GOPATH/bin/version-code-service
```

# Usage

The service can be launch by typing in the terminal:
```sh
$ ./version-code-service
```

By default the service will be listening on port 8080. If you want to specify another port:
```sh
$ ./version-code-service -port 8081
```

# Dependencies
  - [Go]
  - [go-restful]
  - [mgo]
  - [MongoDB]

# License
version-code-service is available under the MIT license. See the LICENSE file for more info.

[Go]:https://golang.org/
[GOPATH]:https://golang.org/doc/code.html#GOPATH
[MongoDB installed]:http://docs.mongodb.org/manual/installation/
[mongod]:http://docs.mongodb.org/manual/reference/program/mongod/
[Go installation]:https://golang.org/doc/install#install
[go-restful]:https://github.com/emicklei/go-restful
[mgo]:https://labix.org/mgo
[MongoDB]:https://www.mongodb.org/
