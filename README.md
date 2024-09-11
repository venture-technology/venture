<h1 align="center"> 🌬️ Venture </h1>

<h1 align="center"> We are security, speed, and technology. We are Venture </h1>

<p align="center">
  <img src="https://i.imgur.com/yieDOSJ.png"/>
</p>

# 👋 Hello, this is the official documentation for the first version of Venture Backend!

### ❓ Who are we?

We are the creators of the Venture app, available for Android and iOS. The app promises to simplify something that continues to cause you headaches. We promise to do our best for everyone, advancing education.

- How do we help drivers?

Our focus was initially on truly making the drivers' lives easier. But, after analyzing and thoroughly reviewing, we saw the possibility of helping everyone.

Drivers face serious issues with receiving payments every month on a fixed date. We solved this problem by requesting automatic payments via Boleto, Debit Card, and Credit Card through external platforms.

Additionally, drivers previously had to check daily which children would or would not attend the next day and create a complete route based on these changes.

We solved this problem as well. Parents can now add in advance whether their child will go to school the next day, along with an optional justification.

Creating the route is no longer a driver’s problem. With the help of an external API, we handle this for them, taking into account who is going or not going to school.

Drivers need to accept an invitation from a school to become part of it, creating a partnership. When a parent registers their child in a school, they can check all available drivers at that school. However, only the school can send the invitation and manage whether the driver is still part of the school.

- How do we help parents and children?

Parents can register all their children in the app and enroll them in different schools with different drivers. When choosing a school, they can see all the drivers partnered with that school.

Parents can view the driver's profile, ratings, benefits, and, of course, the monthly fee.

In the future, we plan to add real-time tracking of the driver’s route while they are transporting their child.

- How do we help schools?

With the app, schools can easily manage drivers, inviting them and hiring them as actual business partners.

It will also be simpler and safer to understand which driver is responsible for the child during the trip.

- Currently, we are still under construction, so please be aware if the service is not complete.

# 🛠️ Dependencies and Languages

### 🔵 Language

Due to its low learning curve and rich web tools, we chose Go, the language from Google.

---

> Before diving into the dependencies and trying to run this service in isolation, we recommend visiting the official Venture repository to run the complete infrastructure. You can also run directly from the Dockerfile or use Go commands for automatic library installation. However, if you really want to check our dependencies, here they are:

### 🫁 Dependencies

- AWS SDK Go
- JWT Go
- Gin Gonic
- Go Redis
- Postgresql
- Gin Swagger
- Files Swagger
- Yaml V2
- Kakfa Go
- QRCode GO
- Base64x
- TOML Go


### 🛢 Migrations

- Please, install golang-migration

> Linux
```bash
$ curl -L https://packagecloud.io/golang-migrate/migrate/gpgkey | apt-key add -
$ echo "deb https://packagecloud.io/golang-migrate/migrate/ubuntu/ $(lsb_release -sc) main" > /etc/apt/sources.list.d/migrate.list
$ apt-get update
$ apt-get install -y migrate
```

> Windows
```bash
$ scoop install migrate
```

- To create a new migration
```bash
migrate create -ext sql -dir database/migrations description-of-migration
```
> f.e: migrate create -ext sql -dir database/migrations add_profile_image_to_drivers

- To run your migration
```bash
migrate -path=database/migrations -database "postgres://user:password@localhost:5432/mydb?sslmode=disable" up
```

- If you need run rollback
```bash
migrate -path=database/migrations -database "postgres://user:password@localhost:5432/mydb?sslmode=disable" down
```

- Case, for some reason, need run migration from the beginning
```bash
migrate create -ext=sql -dir=database/migrations -seq init
```
