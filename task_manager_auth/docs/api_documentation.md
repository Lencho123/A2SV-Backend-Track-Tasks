# Task Management API Documentation

## User Registration
- `POST /register`
- Body: `{ "username": "example", "password": "123456" }`

## Login
- `POST /login`
- Body: `{ "username": "example", "password": "123456" }`
- Returns: `{ "token": "<JWT Token>" }`

## Authenticated Routes
- Pass JWT in header: `Authorization: Bearer <token>`
- Admin-only: Create, Update, Delete tasks
- Users: View tasks by ID and all tasks
