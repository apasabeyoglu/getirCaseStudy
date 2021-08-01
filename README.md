# Getir Case Study

> Go application for Getir case study.  I've tried to keep everything simple. I haven't use any foldering such as 'model' etc. because of the simplicity.

---

## Table of Contents

- [Installation](#installation)
- [Example](#example)
- [Tests](#tests)
- [Technologies](#technologies)

## Installation

### Clone

- Clone this repo to your local machine using `https://github.com/apasabeyoglu/getirCaseStudy`

### Setup

- If you want to use the application ony our local machine, please make sure you have installed the needed applications(shown in technologies section).
  
- Also I have used Redis as my in memory database. For local testing make sure that you have set environment variable for "REDIS_URL". I have implemented a solution in the tests those need to use Redis.

- I have used Heroku for deployment. However, Redis addon for Heroku gave me some headaches, but it is working flawlessly after some configuration on Heroku.

- Only framework I've used for this project is testify which is only for testing.

---

## Example

> Deployed Heroku application link

- https://secure-falls-62285.herokuapp.com/

> Endpoints for the different requests

- (GET) https://secure-falls-62285.herokuapp.com/getir/mongo (This endpoint used for getting data from Getir's MongoDB)
- (GET) https://secure-falls-62285.herokuapp.com/getir (This endpoint used for getting data from Redis)
- (POST) https://secure-falls-62285.herokuapp.com/getir (This endpoint used for saving data to Redis)

## Tests

> I have created unit tests for this project which can be seen in the test files. Project have 70% statement coverage.

---

## Technologies

- Go 1.16
- MongoDB Community Edition@4.2
- Heroku
