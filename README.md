#Cloud Vision
    
* 


### Install Libraries    
* go get github.com/gorilla/mux
* go get -u cloud.google.com/go/vision/apiv1



* Json Send (POST) to Googe Cloud Vision API

      {

          "requests": 
          [
      
               {
    
                    "image": {
    
                        
                        "content":"/9j/7QBEUGhvdG9...image contents...eYxxxzj/Coa6Bax//Z"
                   
                    },
                    
    
                    "features": [
    
                    {
    
                    
                        "type":"TEXT_DETECTION", 
                                   
                        "maxResults":1
    
                 
                     }
                
                    ]
        
                }
            ]
  
       }

* Json Expected to get after parsing (GET)

        {
        
        "name": ""
        "lastNameP": ""
        "lastNameF": ""
        "Gender": ""
        "Birthday": ""
        "Email": ""    *
        "Password": "" *
        "C.P.": ""
        "VIN": ""      *
        
        }