FROM node:20-alpine as base

# set the working directory to /app. This is the directory where the commands will be run. We can use any directory name but /app is a standard convention
WORKDIR /app

COPY package*.json ./

RUN addgroup app && adduser -S -G app app

RUN chown -R app:app .

USER app

RUN npm install

COPY . .

EXPOSE 8080

# development build
FROM base as development

CMD npm run generate; npm run dev;

# production build
FROM base as production

USER root

RUN npm install pm2 -g

USER app

CMD npm run generate; npm run build; npm run start
