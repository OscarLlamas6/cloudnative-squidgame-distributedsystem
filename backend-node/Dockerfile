FROM node:14
# Create app directory
WORKDIR /usr/src/app
# Install app dependencies
COPY package.json ./
RUN npm install
RUN touch .env
RUN chmod 777 .env
# Bundle app source
COPY . .
EXPOSE 4000
CMD [ "npm", "start" ]