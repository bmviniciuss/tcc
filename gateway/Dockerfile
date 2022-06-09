FROM node:16-alpine as builder

RUN npm install -g pnpm

WORKDIR /usr/src/app
COPY package*.json ./
COPY tsconfig*.json ./
COPY . .
RUN pnpm install
RUN pnpm run build


FROM node:16-alpine as runner
RUN npm install -g pnpm

ENV NODE_ENV=production
WORKDIR /usr/src/app
COPY package*.json ./
RUN pnpm install --prod
COPY --from=builder /usr/src/app/dist/ dist/
EXPOSE 3000

CMD [ "node", "dist/application/index.js" ]