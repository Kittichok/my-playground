FROM node:14-alpine

RUN apk update
RUN apk add git sqlite python make g++ > /dev/null

WORKDIR /usr/src/app

COPY package*.json ./

RUN npm install

COPY . .

EXPOSE 4000

CMD [ "npm", "start" ]
