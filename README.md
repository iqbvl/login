# Boilerplate

- Golang
- Go-Chi
- Go-Chi/jwtauth
- Docker
- Docker-compose
- Dep

# Maintainer

Iqbal Abdurrahman at iqbal.jtk09@gmail.com

# How to deploy

1. Install Docker

2. If you are on linux : 
    sudo docker-compose up --build --detach to deploy

3. Make sure you dep installed before

4. Use a rest tool or do curl:
    a. Login :
        
      curl :
                
        curl -d '{"username":"iqbvl", "password":"password"}' -H "Content-Type: application/json" -X POST http://localhost:8080/login 
        
     Content-Type : application/json
        
     Body :
        
        {
            "username" : "iqbvl",
            "password" : "iqbvl"
        }

    b. Register [POST]
        
      Content-Type : application/json
        
      Body : 
        
        {
            "emailAddress": "iqbvlz@gmail.com",
            "username" : "iqbvl",
            "password" : "iqbvl"
        }
        

    c. SendOTP [POST]
        
      Content-Type : application/json
      
      Body : 
    
        {
            "otp" : 123456
        }
        
    d. ForgotPassword [POST]
        
      Content-Type : application/json
      
      Body : 
    
        
        {
            "emailAddress": "iqbvlz@gmail.com"
        }
        

    e. Dashboard [GET]
      
      Authorization : Bearer[space]{token_generated_from_login}
      
# Happy Coding

Happy coding Everyone
