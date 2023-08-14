FROM node:lts-alpine3.18 as builder
WORKDIR /app
COPY . .
RUN yarn install
RUN yarn build

FROM nginx:stable-alpine3.17-slim
COPY --from=builder /app/dist/ /usr/share/nginx/html
