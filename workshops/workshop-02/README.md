# Ingress on local workshop

## Install ingress nginx controller 
```bash
kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/controller-v1.12.1/deploy/static/provider/baremetal/deploy.yaml 
```

**or new version:** [Link](https://kubernetes.github.io/ingress-nginx/deploy/#bare-metal-clusters)
You can install ingress controller via bare-metal-cluster section

## Create Service Type LoadBalancer for EntryPoint
```yaml
apiVersion: v1
kind: Service
metadata:
 name: ingress-nginx-controller-loadbalancer
 namespace: ingress-nginx
spec:
 selector:
   app.kubernetes.io/component: controller
   app.kubernetes.io/instance: ingress-nginx
   app.kubernetes.io/name: ingress-nginx
 ports:
   - name: http
     port: 80
     protocol: TCP
     targetPort: 80
   - name: https
     port: 443
     protocol: TCP
     targetPort: 443
 type: LoadBalancer
```

## Mapping `External-IP` into `hosts` file like this

```bash
$ kubectl get svc ingress-nginx-controller-loadbalancer -n ingress-nginx

OUTPUT: 

NAMESPACE       NAME                                    TYPE           CLUSTER-IP      EXTERNAL-IP     PORT(S)                      AGE
ingress-nginx   ingress-nginx-controller-loadbalancer   LoadBalancer   10.43.45.94     192.168.106.4   80:31184/TCP,443:31741/TCP   9h

$ vi <hosts file> e.g. vi /etc/hosts
# Mastering DevOps
192.168.106.4 jojo.saritrat
# End of Mastering DevOps
```

## Create Application and then expose it
```bash
$ kubectl create deploy nginx --image nginx
$ kubectl expose deploy nginx --port 80 --target-port 80
```

## Create Ingress
```yaml
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
 name: nginx
 annotations:
   nginx.ingress.kubernetes.io/rewrite-target: /$2
   nginx.ingress.kubernetes.io/ssl-redirect: "false"
spec:
 ingressClassName: nginx
 rules:
 - host: jojo.saritrat
   http:
     paths:
       - path: /
         pathType: Prefix
         backend:
           service:
             name: nginx
             port:
               number: 80
```

## Note:
if cannot connect to external IP via `colima` please run command like this:

```bash
colima stop
colima start --kubernetes --network-address
```

while starting the colima or curl the new External-IP `please allow the permission`

Enjoy
