FROM node:18

WORKDIR /react-app

COPY public/ /react-app/public
COPY node_modules/ /react-app/node_modules
COPY src/ /react-app/src
COPY package.json /react-app/package.json
COPY package-lock.json /react-app/package-lock.json


CMD ["npm","start"]


