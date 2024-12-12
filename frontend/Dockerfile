FROM node:lts-alpine as build

WORKDIR /app
COPY package*.json ./
RUN npm install
COPY . .
RUN npm run build

FROM nginx:stable-alpine
# Copia a saida do build para a pasta do nginx
COPY --from=build /app/dist /usr/share/nginx/html
# Copia o arquivo de configuração do nginx para o SPA
COPY ./nginx/default.conf /etc/nginx/conf.d/default.conf
EXPOSE 80
CMD [ "nginx", "-g", "daemon off;" ]