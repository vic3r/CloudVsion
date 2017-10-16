#Cloud Vision
Google Cloud API allows developers understand and know the content of an image through 
encapsulating powerful models of machine learning in a simple api.

####About us

We are currently trying to connect the Google Cloud Vision API with another api, 
this with the goal of reconfiguring the actual text that we obtain from Google. 

### Requirements

* Go
* Python 2.7
* Google Cloud SDK
* Google Cloud Platform

### Install Libraries    
* go get github.com/gorilla/mux
* go get -u cloud.google.com/go/vision/apiv1


### How to run locally

1. Install Go and Python
2. Must have a Google Account
3. Must have Google Cloud Platform (free trial with a credit card)
4. Create a porject in GCP or set a project you already have
5. Set the billing for the project you're using
6. Download Google-Cloud-sdk
7. Run the script inside the Google-Cloud-sdk/bin folder -> ./gcloud
8. Run the command ./gcloud auth application-default login
9. Move this repo inside go/src folder


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

* Expected Json to get after parsing (GET)

        {
        
            "name": ""
            "lastNameP": ""
            "lastNameF": ""
            "Gender": ""
            "Birthday": ""
            "Email": ""    * (Not)
            "Password": "" * (Not)
            "C.P.": ""
            "VIN": ""      * (Not)
        
        }
        
        
###How to deploy

1. Run the application with go run -filename and python -filename
2. Select via args the image that you want to pass
3. (working in third option with endpoints)