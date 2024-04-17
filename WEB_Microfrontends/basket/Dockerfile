FROM node:21-alpine
WORKDIR /app

COPY package.json /app

RUN npm install

COPY . /app

EXPOSE 3002

CMD ["npm", "start"]
