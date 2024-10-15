FROM node:18.17

WORKDIR /app
COPY ./api/typescript/src/package.json ./
RUN npm install

CMD ["npm", "run", "start"]