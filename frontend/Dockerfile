##################################
# BASE STAGE

FROM node:23-alpine as base

# Set working directory
WORKDIR /app

# Copy node package files
COPY package*.json ./

# Install node dependencies
RUN npm install

# Copy rest of source files
COPY . .

##################################
# DEV TARGET

FROM base as dev

CMD ["npm", "run", "dev"]

##################################
# BUILD STAGE

FROM base as build

# Get API address from build args
ARG API_ADDRESS=symphoney.xyz

# Set environment var
ENV API_ADDRESS=${API_ADDRESS}

# Build static site
RUN npm run build

##################################
# PROD TARGET

FROM nginx:1.27.4-alpine as prod

# Copy built webapp to nginx image
COPY --from=build /app/build /app

# Copy nginx configuration file
COPY nginx.conf /etc/nginx/nginx.conf

# Ports to expose for HTTP/S
EXPOSE 80
EXPOSE 443

# Start the server
CMD ["nginx", "-g", "daemon off;"]
