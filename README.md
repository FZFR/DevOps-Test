# README
## How to Run

### Build, Push, and Run Docker Image

1. Build the Docker image (local):

```bash
docker build -t testing/welcome:local .
```
2. Run the Docker image (local port 8000):

```bash
docker run -p 8000:5000 testing/welcome:local
```
3. Run the application using docker-compose (using ghcr and local tag):
- for prod (using ghcr)
```bash
docker-compose up prod
```
- for local
```bash
docker-compose up local
```
**You can access the application at http://localhost:8000**

4. Workflow Dispatch and Deploy to VM
- Workflow Dispatch:
  click this link [here](https://github.com/FZFR/devops-test/actions/workflows/action-workflow-dispatch.yml), then click 'Run Workflow' to choose a branch and input the desired version. Finally, click the 'Run workflow' green button to execute.  
  ![Alt Text](https://media.discordapp.net/attachments/732815516158394418/1150367232631783514/image.png)

- Deploy to VM
  Deploy to your VM only requires adding a 'release' tag to the last commit, with the note that the 'release' tag must follow the following forma `*.*.*` which means there are three variables that need to be filled, for example, `1.0.1`, then push it, and it will automatically build and deploy to the connected VM.


5. K8S config and explanations
This YAML configuration file includes all the specified Kubernetes resources in single file in [here!](./k8s.yaml)  
- ConfigMap:
The ConfigMap stores application configuration as key-value pairs.
- Secret:
The Secret is used to store sensitive data like passwords in an encoded form.
- Deployment:
The Deployment manages the replication of the application. In this case, it deploys an application using the image ghcr.io/fzfr/DevOps-Test:latest.
- Service:
The Service directs traffic to the application based on the label selector app: my-app.
- Ingress:
The first Ingress directs HTTP traffic with the host "localhost" to the my-service.
- Ingress with TLS (letsencrypt):
The second Ingress configures HTTPS with an SSL certificate from Let's Encrypt for the host "fazza.cybernethicc.com".