# Bookstore
A RESTful API for managing a bookstore, written in Go.

## Disclaimer
> [!CAUTION]
> **This project is not intended for use with production data or in a production environment.**
> 
> This is a personal project created for my own learning purposes.

## Overview
A RESTful API for managing a bookstore, featuring endpoints on managing books. Includes a token-based authentication system to secure user access.

**Tech Stack:** Go, Gin framework, SQLite

## API Endpoints
#### 1. Login to Generate Token
**`POST`** /login

Generate a JSON Web Tokens (JWT) for authenticated users. This is needed to access the other API endpoints. 

### These API Endpoints require JWT.
#### 2. Find All Books
**`GET`** /books

Retrieve all the books.

#### 3. Find Book by ID
**`GET`** /books/*{id}*

Retrieve a book based on the Book ID.

#### 4. Create Book
**`POST`** /books

Create a new book with specific details, especially with its ISBN.

#### 5. Update Book
**`PUT`** /books/*{id}*

Update the details of a book.

#### 6. Delete Book
**`DELETE`** /books/*{id}*

Delete a book based on its ID. Only user with admin role can delete a book.

#### 7. Register Account for Manager role
**`POST`** /register

This is to create an account with a manager role.

### These API Endpoints do not require JWT.
#### 8. Changing Password for an Account (Not Reset Password)
**`POST`** /changePassword

Change/update password for an account.
