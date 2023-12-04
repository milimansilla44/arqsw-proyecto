FROM node:16.15.1

WORKDIR /FrontEnd

COPY . /FrontEnd

COPY package.json .
COPY package-lock.json .

RUN npm install

RUN npm run build

EXPOSE 3000

CMD ["npm", "start","build"]

# Path: FrontEnd\docker-compose.yml
