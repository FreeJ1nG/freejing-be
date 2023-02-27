# Endpoints

## Blogs

- `POST` /v1/blogs { title: string, content: string }
- `DELETE` /v1/blogs/:id
- `PATCH` /v1/blogs/:id { title: string, content: string }
- `GET` /v1/blogs
- `GET` /v1/blogs/:id

## Auth

- `GET` /v1/auth/:username
- `POST` /v1/auth { username: string, email: string, password: string }
- `PATCH` /v1/auth { username: string, email: string, password: string }
- `DELETE` /v1/auth/:username
