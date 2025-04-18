# Workshop 2: Healthcheck

## Build Image
มี 3 ภาษาให้เลือก `Go`, `Javascript` และ `Rust`
เมื่อเลือกได้แล้ว ให้ไปที่ภาษานั้น (ก่อนหน้านี้ต้องอยู่ที่ healthchecks ก่อน สามารถตรวจสอบได้ด้วย command `pwd`)
```sh
cd <directory_namge>

# ตัวอย่าง
cd go
```

เมื่อมาที่ directory ของภาษาที่เราต้องการสร้างแล้ว ให้ใช้คำสั่ง
```sh
docker build -t <your_image_name> .

# ตัวอย่าง
docker build -t go-image .
```

เมื่อสร้างเสร็จแล้ว ให้ตรวจสอบว่ามี image ในเครื่องแล้ว โดยใช้คำสั่ง
```sh
docker image ls
```

ต้องเห็นชื่อ image ที่เราสร้างไว้ ยกตัวอย่างเช่น 
| REPOSITORY | TAG | IMAGE ID | CREATED | SIZE | 
| :--------  | :----- | :----------- | :------------ | :----- |
| go-image   | latest | 1234567890ab | 2 minutes ago | 15.9MB | 

หรือถ้าใช้ tool อื่นๆ สามารถเข้าไปดูที่หน้า Images ได้

## Run Container

ใช้คำสั่ง
```sh
docker run -it -d -p <node_port>:<container_port> --name <your_container_name> <your_image_name>

# ตัวอย่าง
docker run -it -d -p 8080:80 --name go-container go-image
```

เมื่อสร้างเสร็จแล้ว ให้ตรวจสอบว่ามี container ในเครื่องแล้ว โดยใช้คำสั่ง
```sh
docker ps
```

ต้องเห็นชื่อ container ที่เรารันไว้ ยกตัวอย่างเช่น
| CONTAINER ID | IMAGE | COMMAND | CREATED | STATUS | PORTS | NAMES |
| :----------- | :------- | :------------- | :------------ | :---------- | :----------------------------- | :----------- |
| ba0987654321 | go-image | "/healthcheck" | 2 minutes ago | Up 1 minute | 8080/tcp, 0.0.0.0:8080->80/tcp | go-container |

หรือถ้าใช้ tool อื่นๆ สามารถเข้าไปดูที่หน้า Containers ได้

## Test container app

- ### Homepage checking

ให้ลองไปที่ web browser แล้วใส่ URL
```
localhost:8080
```

หรือเปิดอีก Terminal แล้วใช้คำสั่ง
```sh
curl localhost:8080
```
จะเจอข้อความต้อนรับ ตัวอย่างจาก Go
```
Hello this is code from Go!
```

ถ้าใช้ web browser ใน **macOS** 

:arrow_right: ให้กด `CMD+Shift+C` แล้วเลือก Network แล้วกด Reload page ให้เลือก localhost เพื่อดู status

ถ้าใช้ web browser ใน **Windows**

:arrow_right: ให้กด `F12` แล้วเลือก Network แล้วกด Reload ให้เลือก localhost เพื่อดู status

- ### Liveness checking

ที่ web browser ใส่ URL
```
localhost:8080/health/liveness
```

หรือเปิดอีก Terminal แล้วใช้คำสั่ง
```sh
curl localhost:8080/health/liveness
```

จะเจอข้อความ
```
200 OK, it lives!
```

- ### Rediness checking

ที่ web browser ใส่ URL
```
localhost:8080/health/readiness
```

หรือเปิดอีก Terminal แล้วใช้คำสั่ง
```sh
curl localhost:8080/health/readiness
```

จะเจอข้อความ
```
200 OK, it's ready!
```

## :fire:Quiz:fire:
ให้ลองเข้าไปใน container (ใส่ container_ID แค่ 3 แรก) โดยใช้คำสั่ง

```sh
docker exec -it <container_ID> sh

# ตัวอย่าง
docker exec -it ba0 sh
```

จะขึ้นแบบนี้
```sh
/app #
```

ต่อมาให้ลบ /tmp/ready ออก โดยใช้คำสั่ง
```sh
rm -rf /tmp/ready
```

ลองเข้า url `localhost:8080/health/readiness` และ `localhost:8080/health/liveness` อีกครั้งเพื่อดูสถานะ

:arrow_forward: จะได้ข้อความและ status อะไร?

(สามารถกด `ctrl+D` เพื่อออกจาก container ได้)

:star:**<u>หมายเหตุ**</u> ถ้าสร้างอันใดอันหนึ่งแล้ว จะสร้างอันต่อไปให้ลบ Container ของเก่าก่อน