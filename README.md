# Todo List

This project is used to learn about the GIN framework. It is my first project and it's learned form [Q1mi](https://github.com/Q1mi/bubble)

## Quick Start

1. Install

```
git clone https://github.com/haoran-mc/Gin-Todo.git
```

2. Create a database

```sql
CREATE DATABASE Gin_Todo;
```

3. Configure database connection information

```go
conf.db   = "Gin_Todo"     // database name
conf.user = "root"         // user name
conf.pwd  = "haoran232"    // user password
conf.host = "127.0.0.1"    // host location
conf.port = 3306           // MySQL port
```

4. Run

```
go run main.go
```

5. Enjoy it!

Open the url `http://127.0.0.1:8002/` in your browser
