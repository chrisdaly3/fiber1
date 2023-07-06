## This is a template repository. Some notes below:

1. The files outlined here are inteded to help with bootstrapping new API web services in Go.
2. Most, if not all, of the code in here is boilerplate, and **SHOULD** be expected to change when implemented in your own projects
3. This stack relies on the [Fiber](https://gofiber.io/) + [GORM](https://gorm.io/) frameworks, for Web routing + ORM interactions respectively.
4. The db is a SQLite database, however this should be able to be hot-swapped pretty easily in [database/database.go](database/database.go).
5. For hot-reload functionality, I make use of the [air](https://github.com/cosmtrek/air) package. It's impressive how painlessly simple it is to use.


### Thanks for reading. Feel free to fork at your leisure.