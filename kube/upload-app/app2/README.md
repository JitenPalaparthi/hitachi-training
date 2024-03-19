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

- port forwarding from ssh

```
ssh -p <port-to-connect-to-vm> -L <hostport>:localhost:<port-inside-vm> username@IP 
ssh -p 23213 -L 50090:192.168.49.2:32575 jiten@34.145.43.55
```

- kubectl port-forward service

```
kubectl port-forward service/uploadapp-service 8082:8081 -n dev
```
- kubectl port-forward pod

```
kubectl port-forward pod/uploadapp-service-abcd-234 8082:8081 -n dev
```
