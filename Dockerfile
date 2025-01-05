FROM node:18

WORKDIR /react-app

COPY public/ /react-app/public
COPY src/ /react-app/src
COPY package*.json /react-app/


CMD ["npm","start"]


