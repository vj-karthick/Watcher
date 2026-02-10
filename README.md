**🚀 Kubernetes CRD Watcher Project**
**📌 Project Title**

**Kubernetes Custom Resource Definition (CRD) Watcher using Go**

**📖 Description**

This project demonstrates how to create a **Custom Resource Definition (CRD)** in Kubernetes and build a **Go-based watcher** that continuously monitors custom resources and logs events when new custom objects are created.

It shows how Kubernetes can be extended using CRDs and how external programs (controllers/watchers) can interact with the Kubernetes API using Go.

**🎯 Purpose of the Project**

- Extend Kubernetes with a custom resource type

- Define a new object using CRD

- Watch and react to custom objects using Go

- Understand Kubernetes API interaction

- Build foundation for Kubernetes controllers/operators

**🧱 Technologies Used**

- **Kubernetes**
- **Custom Resource Definitions (CRD)**
- **Go (Golang)**
- **kubectl**
- **Docker Desktop Kubernetes Cluster**

**📂 Project Structure**
k8s-crd-demo/
│
├── main.go              # Go watcher program
├── crd.yaml             # CRD definition file
├── sample-image.yaml    # Sample custom resource object
├── go.mod               # Go module file
├── go.sum               # Go dependency checksum file
└── logs/
    └── crd.log          # Log file for detected CRD events

**⚙️ How It Works (Flow)**

1. crd.yaml defines a new Kubernetes resource type (ImageScan)

2. Kubernetes registers this new type using CRD

3. main.go connects to Kubernetes API

4. Go watcher listens for new custom objects

5. When a new CRD object is created:
     - It is detected
     - Event is logged in logs/crd.log

**▶️ How to Run the Project**

**Step 1: Apply CRD**

  kubectl apply -f crd.yaml

**Step 2: Run Go Watcher**

  go run main.go

**Step 3: Create Custom Resource**

  kubectl apply -f sample-image.yaml

**Step 4: Check Logs**

  type logs\crd.log   (Windows)


**📥 Input**

  sample-image.yaml (Custom resource object)

**📤 Output**

  Console message:
    New ImageScan detected and logged
  Log file:
    logs/crd.log

**🧠 Learning Outcomes**

 - Understanding Kubernetes extensibility

 - CRD creation and usage

 - Kubernetes API communication

 - Go client usage

 - Watcher/controller concept

 - Event-driven architecture in Kubernetes

**💡 Use Case**

This architecture is useful for:

 - Security scanning systems

 - Image vulnerability scanners

 - DevOps automation tools

 - Kubernetes operators/controllers

 - Monitoring systems

 - Event-driven infrastructure

**👨‍💻 Author**

Karthick M
B.E CSE Graduate
GitHub: https://github.com/vj-karthick




