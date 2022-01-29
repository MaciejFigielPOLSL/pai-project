<h1>Opis</h1>

Jest to prosty przykład aplikacji serwera napisanego w języku Golang wraz ze stroną internetową wykorzystującą framework Vue.js. Jest to strona typu blog, w której można przeglądać dodane artykułu (i dodawać nowe po zalogowaniu).

<h2>Instalacja</h2>
<h3>Wykorzystanie dockera</h3>
W repozytorium jest spakowana ostatnia wersja kontenera server.tar, którą można załadować poleceniem
``` shell
docker load < server.tar
```
i następnie uruchomić poleceniem
``` shell
docker run -p 8080:8080 server
```

Alternatywnie można samemu zbudować obraz dockera za pomocą polecenia
``` shell
docker build -f Dockerfile -t <tag> .
```
Obraz korzysta z dwóch obrazów pośrednich kompilujących stronę www oraz aplikację serwera.

<h3>Instalacja natywna</h3>
Aby zainstalować aplikację można skorzystać ze skryptu `build.sh` na systemie Linux. Na systemie Windows można zastosować odpowiednio zmodyfikowane instrukcje.

Do kompilacji wymagane są programy: golang, npm, vue/cli (instalacja za pomocą polecenia 'npm install -g @vue/cli' w katalogu frontend)


<h1 id="">API</h1>

> Scroll down for code samples, example requests and responses. Select a language for code samples from the tabs above or the mobile navigation menu.

<h1 id="-debug">debug</h1>

## get__api_clear

> Code samples

```shell
# You can also use wget
curl -X GET /api/clear

```

```http
GET /api/clear HTTP/1.1

```

`GET /api/clear`

*Clear all database*

Clear all database

<h3 id="get__api_clear-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|None|

<aside class="success">
This operation does not require authentication
</aside>

<h1 id="-user">user</h1>

## post__api_login

> Code samples

```shell
# You can also use wget
curl -X POST /api/login \
  -H 'Content-Type: application/x-www-form-urlencoded' \
  -H 'Accept: application/json'

```

```http
POST /api/login HTTP/1.1

Content-Type: application/x-www-form-urlencoded
Accept: application/json

```

`POST /api/login`

*Login*

Allows user to log into the system

> Body parameter

```yaml
username: string
password: string

```

<h3 id="post__api_login-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|body|body|object|true|none|
|» username|body|string|true|User login|
|» password|body|string|true|User password|

> Example responses

> 200 Response

```json
{}
```

<h3 id="post__api_login-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|OK|Inline|
|400|[Bad Request](https://tools.ietf.org/html/rfc7231#section-6.5.1)|Bad Request|Inline|
|401|[Unauthorized](https://tools.ietf.org/html/rfc7235#section-3.1)|Unauthorized|Inline|
|500|[Internal Server Error](https://tools.ietf.org/html/rfc7231#section-6.6.1)|Internal Server Error|Inline|

<h3 id="post__api_login-responseschema">Response Schema</h3>

<aside class="success">
This operation does not require authentication
</aside>

## post__api_logout

> Code samples

```shell
# You can also use wget
curl -X POST /api/logout \
  -H 'Accept: application/json'

```

```http
POST /api/logout HTTP/1.1

Accept: application/json

```

`POST /api/logout`

*Logout*

Allows user to log out from the system

> Example responses

> 200 Response

```json
{}
```

<h3 id="post__api_logout-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|OK|Inline|
|400|[Bad Request](https://tools.ietf.org/html/rfc7231#section-6.5.1)|Bad Request|Inline|
|500|[Internal Server Error](https://tools.ietf.org/html/rfc7231#section-6.6.1)|Internal Server Error|Inline|

<h3 id="post__api_logout-responseschema">Response Schema</h3>

<aside class="success">
This operation does not require authentication
</aside>

## post__api_private_deleteme

> Code samples

```shell
# You can also use wget
curl -X POST /api/private/deleteme

```

```http
POST /api/private/deleteme HTTP/1.1

```

`POST /api/private/deleteme`

*Delete user*

Removes user from system

<h3 id="post__api_private_deleteme-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|

<aside class="success">
This operation does not require authentication
</aside>

## post__api_private_me

> Code samples

```shell
# You can also use wget
curl -X POST /api/private/me \
  -H 'Content-Type: application/x-www-form-urlencoded' \
  -H 'Accept: application/json'

```

```http
POST /api/private/me HTTP/1.1

Content-Type: application/x-www-form-urlencoded
Accept: application/json

```

`POST /api/private/me`

*Modify user*

Allows user to modify its data

> Body parameter

```yaml
username: string
name: string
password: string
email: string
showName: true

```

<h3 id="post__api_private_me-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|body|body|object|true|none|
|» username|body|string|true|User login|
|» name|body|string|true|User name|
|» password|body|string|true|User password|
|» email|body|string|true|User email|
|» showName|body|boolean|true|Show name or not|

> Example responses

> 200 Response

```json
"string"
```

<h3 id="post__api_private_me-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|OK|string|
|401|[Unauthorized](https://tools.ietf.org/html/rfc7235#section-3.1)|Unauthorized|string|
|500|[Internal Server Error](https://tools.ietf.org/html/rfc7231#section-6.6.1)|Internal Server Error|string|

<aside class="success">
This operation does not require authentication
</aside>

## post__api_register

> Code samples

```shell
# You can also use wget
curl -X POST /api/register \
  -H 'Content-Type: application/x-www-form-urlencoded' \
  -H 'Accept: */*'

```

```http
POST /api/register HTTP/1.1

Content-Type: application/x-www-form-urlencoded
Accept: */*

```

`POST /api/register`

*Register new user*

Add new user to system

> Body parameter

```yaml
username: string
email: string
password: string

```

<h3 id="post__api_register-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|body|body|object|true|none|
|» username|body|string|true|User login|
|» email|body|string|true|Email address|
|» password|body|string|true|User password|

> Example responses

> 200 Response

<h3 id="post__api_register-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|OK|Inline|
|406|[Not Acceptable](https://tools.ietf.org/html/rfc7231#section-6.5.6)|Not Acceptable|Inline|

<h3 id="post__api_register-responseschema">Response Schema</h3>

<aside class="success">
This operation does not require authentication
</aside>

<h1 id="-articles">articles</h1>

## get__articles_

> Code samples

```shell
# You can also use wget
curl -X GET /articles/ \
  -H 'Accept: application/json'

```

```http
GET /articles/ HTTP/1.1

Accept: application/json

```

`GET /articles/`

*Returns all articles in system*

Returns all articles in system

> Example responses

> 200 Response

```json
{
  "AddDate": "string",
  "AuthorId": "string",
  "AuthorName": "string",
  "ID": 0,
  "Text": "string",
  "Title": "string"
}
```

<h3 id="get__articles_-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|OK|[server.ArticleResponse](#schemaserver.articleresponse)|

<aside class="success">
This operation does not require authentication
</aside>

## post__articles_

> Code samples

```shell
# You can also use wget
curl -X POST /articles/ \
  -H 'Content-Type: application/json'

```

```http
POST /articles/ HTTP/1.1

Content-Type: application/json

```

`POST /articles/`

*Add article*

Allows to add article to system

> Body parameter

```json
{
  "text": "string",
  "title": "string"
}
```

<h3 id="post__articles_-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|body|body|[server.ArticleJsonBody](#schemaserver.articlejsonbody)|true|Article to add|

<h3 id="post__articles_-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|None|

<aside class="success">
This operation does not require authentication
</aside>

## get__articles_:articleId

> Code samples

```shell
# You can also use wget
curl -X GET /articles/:articleId \
  -H 'Accept: application/json'

```

```http
GET /articles/:articleId HTTP/1.1

Accept: application/json

```

`GET /articles/:articleId`

*Returns article with comments*

Returns full article data with array of its comments

<h3 id="get__articles_:articleid-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|articleId|path|integer|true|Article ID|

> Example responses

> 200 Response

```json
{
  "article": {
    "AddDate": "string",
    "AuthorId": "string",
    "AuthorName": "string",
    "ID": 0,
    "Text": "string",
    "Title": "string"
  },
  "comments": [
    {
      "ArticleId": 0,
      "AuthorName": "string",
      "Dislikes": 0,
      "ID": 0,
      "Likes": 0,
      "Text": "string"
    }
  ]
}
```

<h3 id="get__articles_:articleid-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|OK|[server.FullEntryResponseSwagger](#schemaserver.fullentryresponseswagger)|

<aside class="success">
This operation does not require authentication
</aside>

## get__articles_page_:page

> Code samples

```shell
# You can also use wget
curl -X GET /articles/page/:page \
  -H 'Accept: application/json'

```

```http
GET /articles/page/:page HTTP/1.1

Accept: application/json

```

`GET /articles/page/:page`

*Returns articles with pagination*

Returns articles with pagination

<h3 id="get__articles_page_:page-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|page|path|integer|true|Page to display|

> Example responses

> 200 Response

```json
{
  "AddDate": "string",
  "AuthorId": "string",
  "AuthorName": "string",
  "ID": 0,
  "Text": "string",
  "Title": "string"
}
```

<h3 id="get__articles_page_:page-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|OK|[server.ArticleResponse](#schemaserver.articleresponse)|

<aside class="success">
This operation does not require authentication
</aside>

## get__comment

> Code samples

```shell
# You can also use wget
curl -X GET /comment \
  -H 'Accept: application/json'

```

```http
GET /comment HTTP/1.1

Accept: application/json

```

`GET /comment`

*Get comments*

Get call comments in system

> Example responses

> 200 Response

```json
{
  "ArticleId": 0,
  "AuthorName": "string",
  "Dislikes": 0,
  "ID": 0,
  "Likes": 0,
  "Text": "string"
}
```

<h3 id="get__comment-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|OK|[server.CommentResponse](#schemaserver.commentresponse)|

<aside class="success">
This operation does not require authentication
</aside>

## post__comment

> Code samples

```shell
# You can also use wget
curl -X POST /comment \
  -H 'Content-Type: application/json'

```

```http
POST /comment HTTP/1.1

Content-Type: application/json

```

`POST /comment`

*Add comment*

Add comment for article

> Body parameter

```json
{
  "articleId": 0,
  "authorName": "string",
  "text": "string"
}
```

<h3 id="post__comment-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|body|body|[server.CommentJsonBody](#schemaserver.commentjsonbody)|true|Comment|

<h3 id="post__comment-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|None|

<aside class="success">
This operation does not require authentication
</aside>

## get__comment_get_:articleId

> Code samples

```shell
# You can also use wget
curl -X GET /comment/get/:articleId \
  -H 'Accept: application/json'

```

```http
GET /comment/get/:articleId HTTP/1.1

Accept: application/json

```

`GET /comment/get/:articleId`

*Returns article comments*

Returns all comments for article

<h3 id="get__comment_get_:articleid-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|articleId|path|integer|true|Article ID|

> Example responses

> 200 Response

```json
{
  "ArticleId": 0,
  "AuthorName": "string",
  "Dislikes": 0,
  "ID": 0,
  "Likes": 0,
  "Text": "string"
}
```

<h3 id="get__comment_get_:articleid-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|OK|[server.CommentResponse](#schemaserver.commentresponse)|

<aside class="success">
This operation does not require authentication
</aside>

## get__comment_like_:commentId

> Code samples

```shell
# You can also use wget
curl -X GET /comment/like/:commentId

```

```http
GET /comment/like/:commentId HTTP/1.1

```

`GET /comment/like/:commentId`

*Add dislike*

Add dislike to comment

<h3 id="get__comment_like_:commentid-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|commentId|path|integer|true|Comment ID|

<h3 id="get__comment_like_:commentid-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|

<aside class="success">
This operation does not require authentication
</aside>

# Schemas

<h2 id="tocS_server.ArticleJsonBody">server.ArticleJsonBody</h2>
<!-- backwards compatibility -->
<a id="schemaserver.articlejsonbody"></a>
<a id="schema_server.ArticleJsonBody"></a>
<a id="tocSserver.articlejsonbody"></a>
<a id="tocsserver.articlejsonbody"></a>

```json
{
  "text": "string",
  "title": "string"
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|text|string|false|none|none|
|title|string|false|none|none|

<h2 id="tocS_server.ArticleResponse">server.ArticleResponse</h2>
<!-- backwards compatibility -->
<a id="schemaserver.articleresponse"></a>
<a id="schema_server.ArticleResponse"></a>
<a id="tocSserver.articleresponse"></a>
<a id="tocsserver.articleresponse"></a>

```json
{
  "AddDate": "string",
  "AuthorId": "string",
  "AuthorName": "string",
  "ID": 0,
  "Text": "string",
  "Title": "string"
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|AddDate|string|false|none|none|
|AuthorId|string|false|none|none|
|AuthorName|string|false|none|none|
|ID|integer|false|none|none|
|Text|string|false|none|none|
|Title|string|false|none|none|

<h2 id="tocS_server.CommentJsonBody">server.CommentJsonBody</h2>
<!-- backwards compatibility -->
<a id="schemaserver.commentjsonbody"></a>
<a id="schema_server.CommentJsonBody"></a>
<a id="tocSserver.commentjsonbody"></a>
<a id="tocsserver.commentjsonbody"></a>

```json
{
  "articleId": 0,
  "authorName": "string",
  "text": "string"
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|articleId|integer|false|none|none|
|authorName|string|false|none|none|
|text|string|false|none|none|

<h2 id="tocS_server.CommentResponse">server.CommentResponse</h2>
<!-- backwards compatibility -->
<a id="schemaserver.commentresponse"></a>
<a id="schema_server.CommentResponse"></a>
<a id="tocSserver.commentresponse"></a>
<a id="tocsserver.commentresponse"></a>

```json
{
  "ArticleId": 0,
  "AuthorName": "string",
  "Dislikes": 0,
  "ID": 0,
  "Likes": 0,
  "Text": "string"
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|ArticleId|integer|false|none|none|
|AuthorName|string|false|none|none|
|Dislikes|integer|false|none|none|
|ID|integer|false|none|none|
|Likes|integer|false|none|none|
|Text|string|false|none|none|

<h2 id="tocS_server.FullEntryResponseSwagger">server.FullEntryResponseSwagger</h2>
<!-- backwards compatibility -->
<a id="schemaserver.fullentryresponseswagger"></a>
<a id="schema_server.FullEntryResponseSwagger"></a>
<a id="tocSserver.fullentryresponseswagger"></a>
<a id="tocsserver.fullentryresponseswagger"></a>

```json
{
  "article": {
    "AddDate": "string",
    "AuthorId": "string",
    "AuthorName": "string",
    "ID": 0,
    "Text": "string",
    "Title": "string"
  },
  "comments": [
    {
      "ArticleId": 0,
      "AuthorName": "string",
      "Dislikes": 0,
      "ID": 0,
      "Likes": 0,
      "Text": "string"
    }
  ]
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|article|[server.ArticleResponse](#schemaserver.articleresponse)|false|none|none|
|comments|[[server.CommentResponse](#schemaserver.commentresponse)]|false|none|none|

