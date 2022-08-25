# Dwarf - URL Shortener made with Go

High performance url shortener 

## Tech Stack

**Client:** Typescript, Svelte

**Service:** Go, Gofiber, Gorm

**Database:** PostgreSQL inside Docker


## Create Action

<img src="https://github.com/irahuldogra/dwarf/blob/main/screenshots/add-action.gif" alt="drawing" width="auto"/>

## Upadte Action

<img src="https://github.com/irahuldogra/dwarf/blob/main/screenshots/update-action.gif" alt="drawing" width="auto"/>

## Delete Action

<img src="https://github.com/irahuldogra/dwarf/blob/main/screenshots/delete-action.gif" alt="drawing" width="auto"/>


## Run Frontend Locally

Clone the project

```bash
  git clone https://github.com/irahuldogra/dwarf.git
```

Go to the project directory

To run client

```bash
  cd dwarf/app
```

Install dependencies

```bash
  npm install
```


Start the frontend

```bash
  npm run dev
```

# Run Locally Backend


To run Server

```bash
  cd dwarf/service
```

Install dependencies

```bash
  go mod download
```


Start the backend

```bash
  make run
```

# Environment Variables

checkout env.example files

## Authors

- [@irahuldogra](https://www.github.com/irahuldogra)
