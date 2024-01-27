# Hackathon-app
## API

### API lookups
2. "/api/v2":
      - "/events":
          - "/"              POST  
          - "/"              GET 
          - "/:id"           GET 
          - "/:id/image"     GET
          - "/:id"           PATCH   
          - "/:id"           DELETE 
          - "/:id/buyticket" POST

      - "/contact"
          - "/:id"           POST 
<br>
"/api/v2":<br>
<br>
"/api/v2/events":<br>
"/api/v2/events/"              POST<br>
"/api/v2/events/"              GET<br>
"/api/v2/events/:id"           GET<br>
"/api/v2/events/:id/image"     GET<br>
"/api/v2/events/:id"           PATCH<br>
"/api/v2/events/:id"           DELETE<br>
"/api/v2/events/:id/buyticket" POST<br>
<br>
"/api/v2/contact/:id"          POST<br>

<br>
<h3>events</h3>
<hr>
<h4>"/api/v1/events/", method:GET.</h4>

Type | JSON | Headers
--- | --- | ---
Request | --- | ---
Response | [{ "id": "3124",  "name": "Hackathon Liquid",  "id_of_person": "123042193012",  "tags": "crypto,cybersecurity",  "description": "some information",  "images": "3123-2.png",  "start-date": "28-01-2024 ","price": 20,}... ] | ---
Error Response | { "message": "Some text" } | ---

<h4>"/api/v1/events/", method:POST.</h4>

Type | JSON | Headers 
--- | --- | --- 
Request | {"name": "Hackathon Liquid",  "id_of_person": "123042193012",  "tags": "crypto,cybersecurity",  "description": "some information",  "images": "3123-1.png",  "start-date": "28-01-2024", "price":20,} | ---
Response | { "Status": "ok" } | --- 
Error Response | { "message": "Some text" } | --- 

<h4>"/api/v1/events/:id", method:GET.</h4>

Type | JSON | Headers
--- | --- | --- 
Request | --- | --- 
Response | Meta data | ---
Error Response | { "message": "Some text" } | --- 

<h4>"/api/v1/events/:id/image"", method:GET.</h4>

Type | JSON | Headers
--- | --- | --- 
Request | --- | --- 
Response | metadata of image | ---
Error Response | { "message": "Some text" } | --- 

<h4>"/api/v1/events/:id", method:PUT.</h4>

Type | JSON | Headers 
--- | --- | --- 
Request | { "id_of_person": "123042193012", data: {"name": "Hackathon Liquid", "tags": "crypto,cybersecurity",  "description": "some information",  "images": "3123-1.png",  "start-date": "28-01-2024 ", "price": 20,}}  | ---
Response | { "Status": "ok" } | ---
Error Response | { "message": "Some text" } | --- 

<h4>"/api/v1/events/:id", method:DELETE.</h4>

Type | JSON | Headers 
--- | --- | --- 
Request | {"id_of_person": "123042193012"} | ---
Response | { "Status": "ok" } | --- 
Error Response | { "message": "Some text" } | --- 

<h4>"/api/v1/events/:id/buyticket", method:POST.</h4>

Type | JSON | Headers 
--- | --- | --- 
Request | {"id_of_person": "123042193012"} | ---
Response | { "Status": "ok" } | --- 
Error Response | { "message": "Some text" } | ---


## Models
### Events

id   | name | id_of_person | tags | description | images | start-date | price
--- | --- | --- | --- | --- | --- | --- | ---
3123 | Hackathon Liquid | 123042193012 | crypto,cybersecurity | some information | 3123-1.png,3123-2.png | 28-01-2024 | 20

## Configs
### .env
'''
DB_PASSWORD="password"
'''

### config.yml
'''
SERVER:
	Port: port
	
DB:
	Host: host
	Port: port
	Username: admin
	SSLmode: disable
'''

