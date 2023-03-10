# OMS

Use Server -  postgres                  

Use DB - ordertry           

port - 5432     



To run the program - go run main.go     


Post Order - 
    curl --location 'http://localhost:8080/orders' \
    --header 'Content-Type: application/json' \
    --data '{
        "id": 35,
        "status": "first invoice",
                "items": [
                    {
                        "id": 1,
                        "description": "Essential",
                        "price": 125.01,
                        "quantity": 15
                    }
                    
                
                ]
        }'      
        

Update Order - 
    curl --location --request PATCH 'http://localhost:8080/order/{order_id}' \
    --header 'Content-Type: application/json' \
    --data '{
        "id": 5,
        "status": "pending invoice",
        "Item": [
            {
                "id": 125,
                "description": "laptop",
                "price": 125.01,
                "quantity": 15
            }
        ]
    }'
        
        
Get All Orders - 
    GET 'http://localhost:8080/orders'          
    

Get Order By ID - 
    GET http://localhost:8080/order/{id}
