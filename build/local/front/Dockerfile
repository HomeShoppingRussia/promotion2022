FROM node:16

WORKDIR /app

COPY . .
COPY package*.json ./
RUN npm install

EXPOSE 8080

CMD ["npm", "run", "serve"]