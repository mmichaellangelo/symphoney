# Message Types

## HTTP
- Create Room
    - **POST** http://{api}/rooms
- Add Member to Room
    - **POST** http://{api}/rooms/{room}/members
- Remove Member from Room
    - **DELETE** http://{api}/rooms/{room}/members/{member}
- Update Room Properties
    - **PUT/PATCH** http://{api}/rooms/{room}
- Get Room Properties
    - **GET** http://{api}/rooms/{room}
- Get All Room Properties
    - **GET** http://{api}/rooms
- Delete Room
    - **DELETE** http://{api}/rooms/{room}
- Update Member Properties
    - **PUT/PATCH** http://{api}/rooms
- Get All Member Properties for Room
    - **GET** http://{api}/

## Websocket
- Room Socket
    - **WS** ws://{api}/rooms/{room}
    - Packet:
        ```
        {
            "roomID": 
            "memberID":
            "data": {
                "x":
                "y":
            }
        }
        ```
### Data
```
```
