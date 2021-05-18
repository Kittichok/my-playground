FROM node:12-slim

RUN apt-get -y update
RUN apt-get -y upgrade
RUN apt-get install -y sqlite3 libsqlite3-dev

WORKDIR /usr/src/app

COPY package*.json ./

RUN yarn

COPY . .

# EXPOSE 3000

CMD [ "yarn", "start" ]
