FROM node:20.16-slim

RUN npm install -g bun && bun upgrade

WORKDIR /front

COPY . .
RUN bun install && bun run build

CMD ["bun", "run", "dev"]