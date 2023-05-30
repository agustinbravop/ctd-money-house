FROM node:latest AS build

ENV REACT_APP_API_BASE_URL=http://wallet.e2324cfffd41459e8978.brazilsouth.aksapp.io/api/v1

WORKDIR /build

COPY package.json .
COPY package-lock.json .
RUN npm ci

COPY . .
RUN npm run build

FROM nginx:alpine

COPY --from=build /build/build/ /usr/share/nginx/html

RUN rm /etc/nginx/conf.d/default.conf

COPY nginx.conf /etc/nginx/conf.d

CMD ["nginx", "-g", "daemon off;"]

EXPOSE 80