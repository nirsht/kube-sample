package manifest

const Pod = `
apiVersion: v1
kind: Pod
metadata:
  labels:
    app: nginx
    env: staging
  name: demo
  namespace: default
spec:
  containers:
    - image: nginx
      imagePullPolicy: IfNotPresent
      name: demo
      ports:
        - containerPort: 80
          name: http
          protocol: TCP
  dnsPolicy: ClusterFirst
  restartPolicy: Always
  terminationGracePeriodSeconds: 0
`
