![Symphoney](symphoney_small.png)

> **Officially hosted at [symphoney.xyz](http://symphoney.xyz)**

A program for creating interactive computer music experiences.

## How it works

Ever wondered if you could have your audience participate in a piece of electronic music? Well, that's the idea of this project. You'll create a "room" for them to join at which point they can scan a qr code you'll put on screen or enter the given 4-letter code in the website. When they join the room, they'll become a little circle on the screen and can move about by using their touchscreen. 

From there, you'll have to decide what to map each audience member to. We'll have presets that make it easy to get started (eventually).

## Self-hosting

Everything you need to self-host the program is here. You'll need Docker and Docker Compose installed on your server.  
The default port for the API is 8080 which will need to be exposed since the client webapp creates a connection directly to the API.  
The compose file is setup to use certbot for SSL. If you're fine without HTTPS just remove that section. Otherwise you'll need to [set up certbot](https://phoenixnap.com/kb/letsencrypt-docker).