# Overview
Very simple web service that keeps track of an internal current version code (build number) for a given app. Provides an easy-to-use rest API that allows:
* GET /versionCode/<app-id> - Current version code for a given app
* GET /versionCode/<app-id>/next - Increment, then retrieve the current version code for a given app
* DELETE /versionCode/<app-id> - Removes app data from the database
* PUT /versionCode/<app-id>/<version-code> - Sets the app version code with the given value
 


# Dependencies
  - [Go]
  - [go-restful]
  - [MongoDB]

# License
version-code-service is available under the MIT license. See the LICENSE file for more info.

[Go]:https://golang.org/
[go-restful]:https://github.com/emicklei/go-restful
[MongoDB]:https://www.mongodb.org/