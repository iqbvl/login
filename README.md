# Boilerplate Iseng

- Golang
- Go-Chi
- GoChi/jwtauth
- Docker
- Docker-compose

# Maintainer
Iqbal Abdurrahman at iqbal.jtk09@gmail.com

# How to deploy
1. Install Docker
2. If you are on linux : 
    sudo docker-compose up --build --detach
3. Use a rest tool or do curl:
    a. Login [POST] 
        curl : 
        ```sh curl -d '{"username":"iqbvl", "password":"password"}' -H "Content-Type: application/json" -X POST http://localhost:8080/login ```
        Content-Type : application/json
        Body : 
        ```sh
        {
            "username" : "iqbvl",
            "password" : "iqbvl"
        }
        ```
    b. Register [POST]
        Content-Type : application/json
        Body : 
        ```sh
        {
            "emailAddress": "iqbvlz@gmail.com",
            "username" : "iqbvl",
            "password" : "iqbvl"
        }
        ```
    c. SendOTP [POST]
        Content-Type : application/json
        Body : 
        ```sh
        {
            "otp" : 123456
        }
        ```
    d. ForgotPassword [POST]
        Content-Type : application/json
        Body : 
        ```sh
        {
            "emailAddress": "iqbvlz@gmail.com"
        }
        ```
    e. Dashboard [GET]
        Header : Authorization : Bearer {token}