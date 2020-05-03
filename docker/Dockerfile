FROM httpd

COPY cigo-vhost.conf /usr/local/apache2/conf/extra/

RUN printf "Listen 8081\n\
LoadModule cgi_module modules/mod_cgi.so\n\
LoadModule rewrite_module modules/mod_rewrite.so\n\
Include conf/extra/cigo-vhost.conf\n" >>\
 /usr/local/apache2/conf/httpd.conf && \
  mkdir /usr/local/apache2/htdocs/iogo/
