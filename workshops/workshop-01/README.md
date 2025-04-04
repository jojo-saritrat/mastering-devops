# Workshop 1: Traditional Deployment

## Instruction

```bash
# Run app-1
docker run -it -d -p 8080:80 \
 -v $(pwd)/workshops/workshop-01/app-1.html:/usr/share/nginx/html/index.html \
 nginx

# Run app-2 
docker run -it -d -p 8081:80 \
 -v $(pwd)/workshops/workshop-01/app-2.html:/usr/share/nginx/html/index.html \
 nginx

# Run app-3
docker run -it -d -p 8082:80 \
 -v $(pwd)/workshops/workshop-01/app-3.html:/usr/share/nginx/html/index.html \
 nginx

# Running Load Balancer
docker run -d \
  --network=host \
  -p 80:80 \
  -v $(pwd)/workshops/workshop-01/nginx.conf:/etc/nginx/conf.d/default.conf \
  nginx
```

**Note:**
- ไม่จำเป็นต้อง build image สามารถรันจากคำสั่งข้างต้นได้เลย
- option `--network=host` เพื่อให้ container นั้นใช้ network ของ host ซึ่งจะทำให้ localhost:8080 และ localhost:8081 อ้างอิงไปยัง host จริง ๆ