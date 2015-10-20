FROM scratch
MAINTAINER Johannes Scheuermann <johannes.scheuermann@inovex.de>
ADD hello-webserver /hello-webserver
CMD ["/hello-webserver"]
EXPOSE 8000
