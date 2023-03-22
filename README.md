<a name="readme-top"></a>

<!-- PROJECT LOGO -->
<br />
<div align="center">

  <h3 align="center">GUE Ecosystem</h3>

  <h4 align="center">
    Technical Test
  </h4>
</div>

## About The Project

This project is a technical test for GUE Ecosystem. This project is a simple API that consist of 3 services, which are: Auth, Product and Order service. The Auth service is responsible for handling user authentication and authorization. The Product service is responsible for handling product CRUD. The Order service is responsible for handling order CRUD.

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- GETTING STARTED -->

## Getting Started

This is an example of how you may give instructions on setting up your project locally. To get a local copy up and running follow these simple example steps.

### Installation

#### Prerequisites

You need to install these tools before you can run this project

- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)
- [Git](https://git-scm.com/)

#### Steps

1. Clone this repository
   ```sh
   git clone git@github.com:dodysat/gue.git
   ```
   or use https
   ```sh
   git clone https://github.com/dodysat/gue.git
   ```
1. Prepare Docker Images
   ```sh
   docker pull nginx:alpine && docker pull golang:1.20.2
   ```
1. Open Terminal and go to project directory
   ```sh
    cd gue
   ```
1. Run Docker Compose
   ```sh
    docker-compose up --build -d
   ```

<p align="right">(<a href="#readme-top">back to top</a>)</p>

## Usage

### API Documentation & Testing

Please refer to the [API Documentation](https://documenter.getpostman.com/view/1712036/2s93RL2c2t) for more information about the API.

or use

[![Run in Postman](https://run.pstmn.io/button.svg)](https://app.getpostman.com/run-collection/1712036-78f06683-32fb-49cf-ad48-d4dc343131c8?action=collection%2Ffork&collection-url=entityId%3D1712036-78f06683-32fb-49cf-ad48-d4dc343131c8%26entityType%3Dcollection%26workspaceId%3D11d8c495-6222-4984-bd87-9bc040a4b21b)

Collection Variables :

- `url` : `http://localhost:8878`

> ps: I use Postman Test, so you didn't need to manually copy the token from the response.

## Debugging

This project is run using docker-compose. You can change configuration in docker-compose.yml file. You can also change the configuration in each service's Dockerfile.

by default all ports are not exposed to the host machine except for the nginx server. You can change the nginx Exposure in docker-compose.yml file in case you want to change the port.

### Database

Each service has its own mysql server. there are 3 databases, which are: auth-database, product-database, order-database. For debugging purpose, you can access the database using mysql client but you need to expose the port in docker-compose.yml file.

### Application (Services)

Each service has its own docker container. If you want to access URL from the host machine, you need to expose the port in docker-compose.yml file.

### Api Gateway

The api gateway is using nginx server with minimal configuration. You can change the configuration in ./config/nginx.conf file.

By default, the api gateway is using port **8878** but you can change it in docker-compose.yml file.

## Cleanup

After you finish using this application, you can clean up the docker images by running this command

```sh
    docker-compose down --rmi all
```

<p align="right">(<a href="#readme-top">back to top</a>)</p>
