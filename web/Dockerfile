# Development stage
FROM node:16-alpine AS development

RUN npm install -g pnpm

WORKDIR /app

COPY package.json pnpm-lock.yaml ./

RUN pnpm install

COPY . .

EXPOSE 5173

CMD ["pnpm", "run", "dev", "--host"]

# Builder stage
FROM development AS builder
RUN pnpm run build

# Server stage
FROM nginx:1.23.1-alpine

COPY --from=builder /app/dist /usr/share/nginx/html
COPY ./proxy/nginx.conf /etc/nginx/conf.d/default.conf

