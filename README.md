# myblog-search

Application for searching my blog articles at https://blog.seanlee.site/

There are 145 articles: 105 articles in Chinese and 40 articles in English.
All English articles have their counterpart in Chinese.
There are pairs of English and Chinese articles for the same topic.
The search bar provided by blogger.com does not recognize the above relationship,
so there are duplication in the search results when a search keyword exists in both languages.
This application is to avoid the duplication in search results.

The search application is based on Vespa Text Search Tutorial
at https://docs.vespa.ai/en/tutorials/text-search.html.
It is deployed on Google Kubernetes Engine at http://search.seanlee.site/search/?query=%E9%AD%9A

The back end is a single-node Vespa server. It is deployed as a stateful set with persistent
volumes for storing search index.

The middleware is a stateless Golang program to append parameters for Vespa to return search
results in JSON format.

The frond end is a stateless reverse proxy by NGINX. It forwards queries to the middleware and
render search results by Vue.js.

The last component of this search application is a cralwer that downloads the blog articles
in ATOM format, convert the articles into Vespa document format in json format by Golang, and then
feed the Vespa documents into the backend. It is deployed as a Kubernetes CronJob with a
static persisent volume to retain the download blog feed. The retained feed is used for
requesting only the recent updated blog feed instead of full feed. Also, the retained feed can
be used for rebuilding/refeeding the search index.

The following is the data flow of this search application:
<p>
  HTTP client --> NGINX --> Middleware --> Vespa <-- Crawler <-- Blog
</p>

Docker Compose is also used, but only in local development environment. It is for practicing and
comparing the functional difference between Kubernetes and Docker Compose.
