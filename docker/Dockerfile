FROM golang:1.18.3-bullseye

RUN apt update && apt install -y \
    curl \
    nginx \
    nano \
    net-tools\
    supervisor \
    cron \
    && mkdir "/usr/local/nvm"

COPY docker/conf/supervisord.conf     /etc/supervisor/docker_supervisord.conf

ENV NVM_DIR /usr/local/nvm

RUN curl https://raw.githubusercontent.com/nvm-sh/nvm/v0.39.1/install.sh | bash \
    && . $NVM_DIR/nvm.sh \
    && nvm install 16.15.1 && nvm use 16.15.1 \
    && nvm alias default 16.15.1

ENV NODE_PATH $NVM_DIR/versions/node/v16.15.1/lib/node_modules
ENV PATH      $NVM_DIR/versions/node/v16.15.1/bin:$PATH

WORKDIR /app

COPY nuxt-app /app

RUN npm install --location=global yarn \
    && yarn install \
    && npm run build

WORKDIR /usr/src/app

COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY . .
RUN rm -rf nuxt-app && go build -v

RUN rm -rf /etc/nginx/sites-enabled/default
COPY docker/nginx/site.conf /etc/nginx/conf.d/

CMD ["/usr/bin/supervisord", "-c", "/etc/supervisor/docker_supervisord.conf"]