# gin-framework-sample-project
The best practice of Gin RESTful Web with GORM

# config
## set  up config
config.yaml -> 
    environment: dev
so config load from config.dev.yaml   
retrive data  like 
config.Config.GetString("app.url") 

# db
    mysql 

    CREATE TABLE users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(100) NOT NULL,
    password VARCHAR(100) NOT NULL,
    email VARCHAR(100) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);


    mongo
# Test comand
go  test test/*.go -v
# curl 

 

 curl http://localhost:8088/api/v1/user/1

curl -X POST -H 'Content-Type: application/json' -d '{"<key>":"<value>", ...}'
curl -X POST -H 'Content-Type: application/json' -d '{"username":"john_doe","email":"john@example.com"}'  http://localhost:8088/api/v1/user/
  

  curl -X POST -H 'Content-Type: application/json' -d '{"username":"john_doe","password":"john@example.com"}'  http://localhost:8088/api/v1/auth/
  