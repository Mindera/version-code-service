# Overview
Very simple web service that keeps track of an internal current version code (build number) for a given app. Provides an easy-to-use rest API that allows:
* GET /versionCode/&lt;app-id&gt; - Current version code for a given app
* GET /versionCode/&lt;app-id&gt;/next - Increment, then retrieve the current version code for a given app
* DELETE /versionCode/&lt;app-id&gt; - Removes app data from the database
* PUT /versionCode/&lt;app-id&gt;/&lt;version-code&gt; - Sets the app version code with the given value
 


# Dependencies
  - [Go]
  - [go-restful]
  - [MongoDB]

# License
version-code-service is available under the MIT license. See the LICENSE file for more info.

[Go]:https://golang.org/
[go-restful]:https://github.com/emicklei/go-restful
[MongoDB]:https://www.mongodb.org/
