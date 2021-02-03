FROM node:15-alpine

WORKDIR /usr/src/app

RUN apk add --no-cache g++ make python

RUN npm install --quiet node-gyp -g

COPY package.json ./

RUN npm install

COPY . .

# EXPOSE 3000

CMD [ "npm", "start" ]
