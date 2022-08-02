


<h1 align="center"> 🚀 Welcome to this project 🚀 </h1>

<h3 align="center"> Pls, take a look in the software we are producing </h3>
<div align="center">
<a href="https://github.com/RodolfoMRibeiro/" target="_blank">
<img align="center" src=https://img.shields.io/badge/github-%2324292e.svg?&style=for-the-badge&logo=github&logoColor=white alt=github style="margin-bottom: 5px;" /> </a>
<a href="https://www.linkedin.com/in/rodolfomarquesribeiro/" target="_blank">
<img align="center" src=https://img.shields.io/badge/linkedin-%231E77B5.svg?&style=for-the-badge&logo=linkedin&logoColor=white alt=linkedin style="margin-bottom: 5px;" /> </a>
</div> 

<br> <br>

<h3 align="center"> . . . </h3>

# Hey 👋, Read more about!  




### Glad to see you here!  

That's an API make to a Q2 test which going to testing our seniority. We're not high level programers, but being honest, we're in begin our careers and decide to give it a try.

This same API REST, in instance, was made following a really close structure as many other golang programs they using there. Although, in specialy this software, which'll never go to production, aims to show an user registration that simulates a bank internal process. 

A little more about this sample bank internal process... It contains a route to register the user, properly said, the accounts and even simulate some transactions that commonly happen inside a payment sector.

<br>
<p align="center"> The instructors are Known as Gui and Dimi and are teaching us 'til become seniors </p>


<div align="center">
<a href="https://git-scm.com/" target="_blank" rel="noreferrer"> <img style="margin: 10px" src="https://www.vectorlogo.zone/logos/git-scm/git-scm-icon.svg" alt="git" width="40" height="40" align="center"/> </a> 
<a href="https://golang.org" target="_blank" rel="noreferrer"> <img style="margin: 10px" src="https://raw.githubusercontent.com/devicons/devicon/master/icons/go/go-original.svg" alt="go" width="40" height="40" align="center"/> </a> 
<a href="https://www.mysql.com/" target="_blank" rel="noreferrer"> <img style="margin: 10px" src="https://raw.githubusercontent.com/devicons/devicon/master/icons/mysql/mysql-original-wordmark.svg" alt="mysql" width="40" height="40" align="center"/> </a> 
<a href="https://ubuntu.com/" target="_blank" rel="noreferrer"> <img style="margin: 10px" src="https://profilinator.rishav.dev/skills-assets/linux-original.svg" alt="Linux" width="40" height="40" align="center"/> </a>
<a href="https://www.docker.com/" target="_blank" rel="noreferrer"> <img style="margin: 10px" src="https://profilinator.rishav.dev/skills-assets/docker-original-wordmark.svg" alt="Docker" width="40" height="40" align="center"/> </a>
</div>  
<h6 align="center"> < Tools Used /> </h6>


## Database Modeling
<table><tr><td valign="top" width="50%">

- 🔭 Relational database
  

- 🌱 Coverage: 1º, 2º, 3º Normal forms
  

- ❓ Logic modeling
  

- ⚡ Generated by package gorm


</td><td valign="top" width="50%">

<div align="center">
    
<img style="margin: auto" src="https://user-images.githubusercontent.com/89111957/181146633-b602e870-492d-4529-8242-b82f94a4b6bc.png" align="center" style=" width=200px height= 200px " />
</div>  


</td></tr></table>  
<br> <br> 

## Use the following command to run docker


```

$ docker-compose up -d

```
<br> <br> 




## Project Architecture

```bash
.
├── configs
│    ├── configFuncs.go
│    └── configs.go
│  
├── db
│    ├── seed
│    |    └── seed.go
|    |
│    ├── database.go
│    └── runner.go
|
├── integrations 
|    ├── integrations_test.go
│    └── integrations.go
|
├── library
│    └── consts.go
|
├── module
│    ├── controllers
│    |    ├── controller_transaction.go
│    |    ├── controller_account.go
│    │    └── controller_user.go
│    │
│    ├── models
│    |    ├── model_transaction.go
│    |    ├── model_interfaces.go
│    |    ├── model_category.go
│    |    ├── model_account.go
│    |    ├── model_status.go
│    │    └── model_user.go
│    │
|    ├── repositories
│    |    ├── repository_transaction.go
│    |    ├── repository_account.go
│    │    └── repository_user.go
│    │ 
│    ├── services
│    │    └── service.go
│    │
│    └── routers
│         ├── routes_transaction.go
│         ├── routes_account.go
│         └── routes_user.go
|
├── server
│    └── server.go
|
├── routers 
│    └── routes.go
|
├── util
│    ├── utilities_test.go
│    └── utilities.go
|
├── .env-example
├── .gitignore
├── docker-compose.yml
├── go.mod
├── go.sum
├── main.go
└── README.md


```

<hr>
<h3 align="center"> ⚡ Thanks for ur coming ⚡ </h3>
<h4 align="center"> . . . </h4>

