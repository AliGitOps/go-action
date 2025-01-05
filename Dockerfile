FROM node:18

WORKDIR /react-app

COPY ./* /react-app/
#COPY src/ /react-app/src
#COPY package*.json /react-app/


CMD ["npm","start"]


