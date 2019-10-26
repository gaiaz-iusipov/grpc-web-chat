FROM node:10-alpine as builder
ARG VUE_APP_API_URL=http://localhost:3000
RUN set -eux \
    && mkdir /app \
    && chown node:node /app
WORKDIR /app
USER node
COPY --chown=node . .
RUN set -eux \
    && yarn install \
    && yarn build \
    && yarn cache clean

FROM nginx:1.17-alpine
COPY --from=builder /app/dist/ /usr/share/nginx/html
