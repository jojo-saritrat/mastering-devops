FROM node:23.10-alpine as builder
USER node
WORKDIR /app
COPY ./package*.json ./
RUN npm install
COPY --chown=node . .
RUN npm run build

FROM nginx:1.27-alpine
WORKDIR /var/www/app
COPY --from=builder /app/dist/ /var/www/app/