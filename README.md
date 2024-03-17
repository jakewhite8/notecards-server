<a name="readme-top"></a>
# Notecard Server

<!-- PROJECT SHIELDS -->
[![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]


<!-- LOGO -->
<br />
<div align="center">

  <h3 align="center">Notecards Server</h3>

  <p align="center">
    A Gin REST API
    <br />
    <br />
    <a href="https://github.com/jakewhite8/notecards-server/issues">Report Bug</a>
    ·
    <a href="https://github.com/jakewhite8/notecards-server/issues">Request Feature</a>
  </p>
</div>

<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li>
      <a href="#run-the-app-locally">Run The App Locally</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#contact">Contact</a></li>
    <li><a href="#acknowledgments">Acknowledgments</a></li>
  </ol>
</details>

## About The Project

This is the server for the Notecards application. The Notecards server accepts requests from the client, performs CRUD operations on a Postgres database, then sends responses back to the client. The corresponding User Interface is located in the [Notecards repository](https://github.com/jakewhite8/Notecards)

<p align="right">(<a href="#readme-top">back to top</a>)</p>

### Built With

The server is built with the Go web framework, Gin. It uses GORM to interact with the Postgres database hosted on a Amazon RDS.

* [![Go][Go]][Go-url]
* [![Gin][Gin]][Gin-url]

<p align="right">(<a href="#readme-top">back to top</a>)</p>

## Run The App Locally
### Prerequisites
The server was built with version 1.20.11 of Go

### Installation
1. Clone the repo
   ```sh
   git clone https://github.com/jakewhite8/notecards-server.git
   ```
2. From inside the notecards-server directory, install the project dependencies by running:
   ```sh
   go mod download
   ```   
3. Start the server:
   ```sh
   go run main.go
   ```
4. Navigating to localhost:8080 should display the message: Notecards Server

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- CONTACT -->
## Contact

Jake White - jake.white@colorado.edu

Project Link: [https://github.com/jakewhite8/notecards-sever](https://github.com/jakewhite8/notecards-server)

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- ACKNOWLEDGMENTS -->
## Acknowledgments

Resources I found helpful and would like to give credit to
* [Img Shields](https://shields.io)
* [Best-README-template](https://github.com/othneildrew/Best-README-Template/blob/master/README.md#readme-top)


<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- MARKDOWN LINKS & IMAGES -->
[contributors-shield]: https://img.shields.io/github/contributors/jakewhite8/notecards-server.svg?style=for-the-badge
[contributors-url]: https://github.com/jakewhite8/notecards-server/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/jakewhite8/notecards-server.svg?style=for-the-badge
[forks-url]: https://github.com/jakewhite8/notecards-server/network/members
[stars-shield]: https://img.shields.io/github/stars/jakewhite8/notecards-server.svg?style=for-the-badge
[stars-url]: https://github.com/jakewhite8/notecards-server/stargazers
[issues-shield]: https://img.shields.io/github/issues/jakewhite8/notecards-server.svg?style=for-the-badge
[issues-url]: https://github.com/jakewhite8/notecards-server/issues
[Go]: https://img.shields.io/badge/go-grey?style=for-the-badge&logo=go
[Go-url]: https://go.dev/
[Gin]: https://img.shields.io/badge/gin-green?style=for-the-badge&logo=gin
[Gin-url]: https://github.com/gin-gonic/gin
[Expo-url]: https://expo.dev/
