FROM node:21-alpine
WORKDIR /app

COPY package.json /app

RUN npm install

COPY . /app

EXPOSE 3001

CMD ["npm", "start"]
