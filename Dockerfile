FROM node:24-alpine3.22 AS build

WORKDIR /app

RUN corepack enable pnpm

COPY package.json .
COPY pnpm-lock.yaml .
COPY pnpm-workspace.yaml .

RUN pnpm i --frozen-lockfile 

COPY . .

RUN pnpm run build

FROM nginx:1.29.4-alpine

WORKDIR /usr/share/nginx/html

COPY --from=build /app/dist .
