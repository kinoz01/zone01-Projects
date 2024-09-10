FROM node:21-bookworm-slim

WORKDIR /app/

COPY package.json package-lock.json ./ 

RUN npm install

RUN apt-get update -y

RUN apt-get install golang python3 curl -y

RUN curl https://sh.rustup.rs/ -sSf | sh -s -- -y

COPY . /app/

CMD ["npm", "start"]
