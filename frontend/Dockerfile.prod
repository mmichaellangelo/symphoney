FROM alpine:latest

# Set working directory
WORKDIR /app

# Install dependencies
RUN apk add --update --no-cache \
    nodejs \
    npm \
    nginx

# Copy node package files
COPY package*.json .

# Install node dependencies
RUN npm install

# Copy all files to /app
COPY . .

# Build static site
RUN npm run build

RUN chown -R nginx:nginx /app/build
RUN chmod -R 755 /app/build


# Copy nginx configuration file
COPY nginx.conf /etc/nginx/nginx.conf

# Expose port for HTTP (HTTPS TODO)
EXPOSE 80

# Start the server
CMD ["nginx", "-g", "daemon off;"]
