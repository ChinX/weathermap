FROM node:network

RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
    echo "Asia/Shanghai" > /etc/timezone

WORKDIR /home/apps/
COPY ./app.js ./
COPY ./bin ./bin
COPY ./package.json ./
COPY ./public ./public
COPY ./routes ./routes
COPY ./shutdown.sh ./
COPY ./start.sh ./
COPY ./views ./views
RUN npm install

ENTRYPOINT ["node", "/home/apps/bin/weathermapweb"]
