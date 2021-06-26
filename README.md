# myblog-search

Application for searching my blog articles at http://diaryofsean.blogspot.com/

There are 145 articles: 105 articles in Chinese and 40 articles in English.
All English articles have their counterpart in Chinese.
There are pairs of English and Chinese articles for the same topic.
The search bar provided by blogger.com does not recognize the above relationship,
so there are duplication in the search results when a search keyword exists in both languages.
This application is to avoid the duplication in search results.

The search application is based on Vespa Text Search Tutorial
at https://docs.vespa.ai/en/tutorials/text-search.html

The frond end is a reverse proxy by NGINX and XSLT to transform search results from XML into HTML

Live demo at http://34.87.254.125/search/%E9%AD%9A on Google Cloud Platform
