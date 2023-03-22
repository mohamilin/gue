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

A

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- GETTING STARTED -->

## Getting Started

This is an example of how you may give instructions on setting up your project locally. To get a local copy up and running follow these simple example steps.

### Installation

#### Prerequisites

- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)
- [Git](https://git-scm.com/)

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
   docker pull nginx:alpine && docker pull golang:1.20.2 && docker pull scratch
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

## Debugging

## Cleanup

After you finish using this application, you can clean up the docker images by running this command

```sh
    docker-compose down --rmi all
```

Remember to remove the images that you pull before running docker-compose

```sh
    docker image rm mysql:8.0
    docker image rm nginx:alpine
    docker image rm golang:1.20.2
    docker image rm scratch
```

<p align="right">(<a href="#readme-top">back to top</a>)</p>
