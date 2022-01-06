FROM node:lts-alpine3.15 as builder
WORKDIR /app
COPY . .
RUN yarn install \
    && yarn build

FROM nginx:1.21.5-alpine
COPY --from=builder /app/dist/ /usr/share/nginx/html
