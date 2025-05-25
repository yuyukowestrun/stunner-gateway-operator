# STUNner Gateway Operator 🚀

Welcome to the STUNner Gateway Operator repository! This project aims to simplify the deployment and management of STUNner as a Kubernetes Gateway. STUNner leverages WebRTC technology to enhance communication in cloud-native environments.

[![Download Latest Release](https://img.shields.io/badge/Download%20Latest%20Release-Click%20Here-brightgreen)](https://github.com/yuyukowestrun/stunner-gateway-operator/releases)

## Table of Contents

- [Introduction](#introduction)
- [Features](#features)
- [Installation](#installation)
- [Usage](#usage)
- [Configuration](#configuration)
- [Contributing](#contributing)
- [License](#license)
- [Support](#support)

## Introduction

STUNner is a powerful gateway designed to facilitate WebRTC communications in Kubernetes. It acts as a bridge between clients and servers, allowing seamless connectivity. This operator automates the deployment and management of STUNner, making it easier for developers to integrate WebRTC into their applications.

## Features

- **Kubernetes Native**: Built specifically for Kubernetes, ensuring compatibility and ease of use.
- **WebRTC Support**: Optimized for WebRTC applications, providing low-latency communication.
- **Automatic Scaling**: Automatically scales resources based on demand.
- **Easy Configuration**: Simple configuration options to get started quickly.
- **Monitoring and Logging**: Built-in monitoring and logging capabilities for better observability.

## Installation

To install the STUNner Gateway Operator, follow these steps:

1. **Clone the Repository**:
   ```bash
   git clone https://github.com/yuyukowestrun/stunner-gateway-operator.git
   cd stunner-gateway-operator
   ```

2. **Build the Operator**:
   Ensure you have the necessary tools installed, such as Go and Kubernetes CLI. Then run:
   ```bash
   make build
   ```

3. **Deploy to Kubernetes**:
   You can deploy the operator using the provided Kubernetes manifests:
   ```bash
   kubectl apply -f deploy/
   ```

4. **Download and Execute Latest Release**:
   Visit [Releases](https://github.com/yuyukowestrun/stunner-gateway-operator/releases) to download the latest release. Execute the downloaded file to complete the installation.

## Usage

Once installed, you can start using the STUNner Gateway Operator by creating a custom resource definition (CRD). Here’s a basic example:

```yaml
apiVersion: stunner.example.com/v1
kind: Stunner
metadata:
  name: my-stunner
spec:
  replicas: 2
  service:
    type: LoadBalancer
```

Apply this configuration using:
```bash
kubectl apply -f my-stunner.yaml
```

This will create a STUNner instance with two replicas, and expose it via a LoadBalancer service.

## Configuration

The STUNner Gateway Operator supports various configuration options to tailor its behavior. Here are some key settings:

- **Replicas**: Defines the number of STUNner instances to run.
- **Service Type**: Choose between `ClusterIP`, `NodePort`, or `LoadBalancer`.
- **Logging Level**: Set the verbosity of logs (e.g., `info`, `debug`).

You can specify these options in your CRD YAML file as shown above.

## Contributing

We welcome contributions! If you would like to contribute to the STUNner Gateway Operator, please follow these steps:

1. Fork the repository.
2. Create a new branch for your feature or bug fix.
3. Make your changes and commit them with clear messages.
4. Push your changes and create a pull request.

Please ensure your code adheres to the existing style and includes tests where applicable.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Support

If you encounter any issues or have questions, feel free to open an issue on GitHub. For additional resources, visit our [Releases](https://github.com/yuyukowestrun/stunner-gateway-operator/releases) section to find the latest updates and downloads.

---

Thank you for your interest in the STUNner Gateway Operator! We hope this project helps you in your WebRTC endeavors.