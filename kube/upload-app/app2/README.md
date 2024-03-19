- Be default this app runs on 8081
- To change the port on which the app start listening APP_PORT=8082
- There are two endpoints
    - /ping --> To check whether it is working or not
    - /upload --> To upload a png image.

            ```
            curl --location 'localhost:50090/upload' --form 'myFile=@"/home/jiten/Desktop/Designer (5).png"'
            
            ```
- To build docker image

```
docker build -t jpalaparthi/hitachi_upload_app:v01 .
```

- To push the image to docker hub

```
docker push jpalaparthi/hitachi_upload_app:v01
```