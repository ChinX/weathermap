FROM ubuntu:network

RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
    echo "Asia/Shanghai" > /etc/timezone

WORKDIR /home/apps/
ADD fusionweather .

ENTRYPOINT ["/home/apps/fusionweather"]
