# How To Run

- make run

# APIs

- Get Insights of a website
    - 'POST' `43.207.178.99:8080/insight`
    - "Body" -> {"url": "https://www.suitejar.com/blog"}

- Get all the previously fetched insights
  - 'GET'  `43.207.178.99:8080/insights`

- Add to Favourite
    - 'GET' `43.207.178.99:8080/insights/:id`

- Delete a insight
    - 'DELETE' `43.207.178.99:8080/insights/:id`



# Tech Stack
    - Golang
    - Gin
    - Gorm
    - Sqlite3
    - net/html
    - net/http
    - AWS EC2
