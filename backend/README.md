---
runme:
  id: 01HN6B1KH0KW12R2NKAGHJ4AME
  version: v2.2
---

# Hackathon-app

## API

### API lookups

#### List lookup

"/api/v2":

- "/sub":
   - "/"              POST
   - "/"              GET
   - "/:id"           GET
   - "/:id/image"     GET
   - "/:id"           PATCH
   - "/:id"           DELETE
   - "/:id/buyticket" POST

#### Point lookup

"/api/v2/sub/"              POST
"/api/v2/sub/"              GET
"/api/v2/sub/:id"           GET
"/api/v2/sub/:id"           Delete
"/api/v2/sub/:id"           PATCH

<br>
<h3>Subscription</h3>
<hr>
<h4>"/api/v1/sub/", method:GET.</h4>

Type | JSON | Headers
--- | --- | ---
Request | --- | ---
Response | [{ "channel_id":123, "name":"Name", "price": 20, "description": "some text", "user_id":"1230412", "link":"t.me/213dqw", "images":"img.png", "tags": "crypto,res", wallet: "12ed1o@31pjo"}... ] | ---
Error Response | { "message": "Some text" } | ---

<h4>"/api/v1/sub/", method:POST.</h4>

Type | JSON | Headers
--- | --- | ---
Request | {"channel_id":123, "name":"Name", "price": 20, "description": "some text", "user_id":"1230412", "link":"t.me/213dqw", "images":"img.png", "tags": "crypto,res", wallet: "12ed1o@31pjo"} | ---
Response | { "Status": "ok" } | ---
Error Response | { "message": "Some text" } | ---

<h4>"/api/v1/sub/:id", method:GET.</h4>

Type | JSON | Headers
--- | --- | ---
Request | --- | ---
Response | {"channel_id":123, "name":"Name", "price": 20, "description": "some text", "user_id":"1230412", "link":"t.me/213dqw", "images":"img.png", "tags": "crypto,res", wallet: "12ed1o@31pjo"} | ---
Error Response | { "message": "Some text" } | ---

<h4>"/api/v1/sub/:id", method:PUT.</h4>

Type | JSON | Headers
--- | --- | ---
Request | { "change_person": "1230412", "name":"Name", "price": 20, "description": "some text", "images":"img.png", "tags": "crypto,res", wallet: "12ed1o@31pjo"}  | ---
Response | { "Status": "ok" } | ---
Error Response | { "message": "Some text" } | ---

<h4>"/api/v1/sub/:id", method:DELETE.</h4>

Type | JSON | Headers
--- | --- | ---
Request | {"change_person": "123042193012"} | ---
Response | { "Status": "ok" } | ---
Error Response | { "message": "Some text" } | ---

## Models

### Subscriptions

channel_id | name | price | description | user_id | link | images | tags | wallet
--- | --- | --- | --- | --- | --- | --- | --- | ---
1233123 | TonAppchannel | 20 | Something | 12309812948 | t.me/qwed | 3123-2.png | crypto,pirat | ewrkljlk!@#k

## Configs

### .env

'''
DB_PASSWORD="password"
SECRET_KEY="secret_key"
'''

### config.yml

'''
Enviroment: "dev"

server:
Port: port

db:
Host: host
Port: port
Name: name
Username: admin
SSLmode: disable
'''

