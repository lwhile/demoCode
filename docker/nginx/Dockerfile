FROM ubuntu16.04

MAINTAINER NGINX Docker Maintainers "docker-maint@nginx.com"


RUN apt-get update \
                        && apt-get install --no-install-recommends --no-install-suggests -y \
                                                                        ca-certificates \
                                                                                                                                                                        nginx-module-xslt \
                                                                                                                                                                                                                        nginx-module-geoip \
                                                                                                                                                                                                                                                                        nginx-module-image-filter \
                                                                                                                                                                                                                                                                                                                        nginx-module-perl \
                                                                                                                                                                                                                                                                                                                                                                        nginx-module-njs \
                                                                                                                                                                                                                                                                                                                                                                                                                        gettext-base \
                                                                                                                                                                                                                                                                                                                                                                                                                                && rm -rf /var/lib/apt/lists/*

# forward request and error logs to docker log collector
RUN ln -sf /dev/stdout /var/log/nginx/access.log \
        && ln -sf /dev/stderr /var/log/nginx/error.log

        EXPOSE 80 443

        CMD ["nginx", "-g", "daemon off;"]
