FROM node:16-alpine

WORKDIR /app

COPY /src .

COPY package.json .
COPY package-lock.json .

RUN npm install

CMD ["npm", "start"]
