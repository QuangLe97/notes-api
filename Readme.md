# TỔNG QUAN

## Architecture diagram
![](../../../Architecture.png)

![](../../../process.png)
## API

**GET ALL NOTES**
* METHOD : GET
* URL : host:8787/notes
* PARAMS :
  * page : int
  * limit: int
  * < name field >.< [contains,equals,in] >: str  ex: name.contains
  * sort :str
  
**GET DETAIL NOTE**
* METHOD : GET
* URL : host:8787/notes/ID

**CREATE NOTE**
* METHOD : POST
* URL : host:8787/notes
* BODY : {"name":str, "content":str}

**UPDATE NOTE**
* METHOD : PUT
* URL : host:8787/notes/ID
* BODY : {"name":str, "content":str}

**DELETE NOTE**
* METHOD : DELETE
* URL : host:8787/notes/ID


# START SERVICE
SET CÁC ENV
- DB_PATH=storage.db
- RABBITMQ_URL=42.119.139.251:31678/
- RABBITMQ_USER=rabbit-admin
- RABBITMQ_NAME=note
- RABBITMQ_PASS=xHc2zUkq4PZLeQ2C

> Run file main.go : start service CURD note

> Run file consumer.go : start service ghi log