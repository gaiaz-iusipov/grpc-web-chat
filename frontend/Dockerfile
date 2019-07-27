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
    && yarn run build \
    && yarn cache clean

FROM node:10-alpine
RUN yarn global add serve
RUN set -eux \
    && mkdir /dist \
    && chown node:node /dist
WORKDIR /dist
USER node
COPY --from=builder --chown=node /app/dist/ .
ENV PORT=3000
EXPOSE $PORT
CMD ["serve", "-s", "."]
