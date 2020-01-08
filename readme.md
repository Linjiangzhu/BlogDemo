# A simple golang blog program for beginners
This project is simple enough for beginners to understand mvc structured web applicaiton.
I will rebuild the blog in the future with gin+gorm+react if I have time. Writing web template like jsp is really outdated.
### Required packages
- [gin-gonic/gin](https://github.com/gin-gonic/gin)
- [xs/xid](https://github.com/rs/xid)
- [go-sql-driver/mysql](https://github.com/go-sql-driver/mysql)
### program directory
```
├── assets
│   ├── css
│   │   └── main.css
│   ├── images
│   └── js
├── controller
│   ├── admin.go
│   ├── blog.go
│   ├── index.go
│   └── login.go
├── main.go
├── middleware
│   └── middleware.go
├── model
│   ├── database.go
│   └── model.go
├── readme.md
├── routes
│   └── router.go
└── view
    ├── admin
    │   ├── dashboard.html
    │   └── login.html
    ├── blog
    │   ├── blog.html
    │   ├── edit.html
    │   ├── new.html
    │   └── post.html
    ├── home
    │   └── index.html
    └── partials
        ├── footer.html
        └── header.html
```