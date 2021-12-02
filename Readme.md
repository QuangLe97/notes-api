#TỔNG QUAN

##API

**GET ALL NOTES**
* METHOD : GET
* URL :localhost:8787/notes
* PARAMS :
  * page : int
  * limit: int
  * < name field >.< [contains,equals,in] >: str
  * sorts
  
**GET DETAIL NOTE**
* METHOD : GET
* URL :localhost:8787/notes/ID

**CREATE NOTE**
* METHOD : POST
* URL :localhost:8787/notes
* BODY : {"name":str, "content":str}

**UPDATE NOTE**
* METHOD : PUT
* URL :localhost:8787/notes/ID
* BODY : {"name":str, "content":str}

**DELETE NOTE**
* METHOD : DELETE
* URL :localhost:8787/notes/ID


##START SERVICE
SET CÁC ENV
- DB_PATH=storage.db
- RABBITMQ_URL=42.119.139.251:31678/
- RABBITMQ_USER=rabbit-admin
- RABBITMQ_NAME=note
- RABBITMQ_PASS=xHc2zUkq4PZLeQ2C

Run file main.go : start service CURD note
Run file consumer.go : start service ghi log