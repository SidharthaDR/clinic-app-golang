# 🏥 Clinic Management App (Go + Gin + PostgreSQL)

A simple yet robust web application for managing clinic patient data, built with **Go (Gin)** and **PostgreSQL**, featuring JWT-based authentication and role-based access control (Doctor & Receptionist).

---

## 🎯 Key Features

- **User Auth**: Signup & Login with JWT tokens  
- **Role-based Access**:
  - **Receptionist**: Create,View,Update & Delete patients  
  - **Doctor**: View & Update patient data  
- **Full Patient CRUD** (RESTful API)  
- **MVC Architecture**: Controllers, Models, Middleware cleanly organized  
- **DB Setup**: PostgreSQL with environment-based configuration  
- **Unit Tests**: For authentication and patient workflows  
- **API Protection**: Auth & Role-based middleware
- **API Testing**: Postman

---

## 🧱 Project Structure


```md
clinic-app-golang/
│   .env        ---> Make your own .env file, described in coming steps
│   .gitignore
│   go.mod
│   go.sum
│   README.md
│
├───cmd
│       main.go
│
├───internal
│   ├───controllers
│   │       auth.go
│   │       patient.go
│   │
│   ├───middleware
│   │       auth.go
│   │       role_middleware.go
│   │
│   ├───models
│   │       patient.go
│   │       user.go
│   │
│   ├───repository
│   ├───services
│   └───utils
│           db.go
│
└───tests
        .env
        example_test.go
```

---



---

## 🚀 Running Locally

1. Set up PostgreSQL and create a DB with name "clinic"
        
        1.1 create DB with name "clinic_test" (Dummy db for testing)

2. Clone this repo
3. Create a `.env` file at root:

   ```.env
    DB_HOST= localhost
    DB_PORT= 5432
    DB_USER= postgres
    DB_PASSWORD= YourPassowrd123
    DB_NAME= clinic

   ```
4. Run the server from root:
    ```cmd
        go run cmd/main.go
    ```

5. Postman for testing Api, and doing CRUD operations based on roles:
  ```.env
   1. Open Postman App or Web.
   2. In the sidebar:
     Click "Collections".
   3. Click the "Import" button (top-left).
   4. In the popup:
        Choose "Upload Files".
        Select your .json file (e.g., clinic-api.postman_collection.json).
        Click "Import".
             
```
 - After you import this into Postman, after Auth phase, use generated login token for role based( Doctor/Receptionist) operations. we made 2 users -> doctor01, reception1 
 - Refer this Postman doc of this project for further clarity: https://documenter.getpostman.com/view/43825571/2sB2x5GsE7
 - After this, now you can Create/read/update/delete based on your role.
