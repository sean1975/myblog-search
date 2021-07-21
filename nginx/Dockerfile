FROM nginx:1.21.0

COPY etc/nginx.conf /etc/nginx/nginx.conf
COPY etc/templates/default.conf.template /etc/nginx/templates/default.conf.template
COPY html/index.html /usr/share/nginx/html/index.html
COPY html/searchbar.html /usr/share/nginx/html/searchbar.html
COPY css/mystyle.css /usr/share/nginx/css/mystyle.css
COPY xslt/result.xsl /usr/share/nginx/xslt/result.xsl
