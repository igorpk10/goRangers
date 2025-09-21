# 🦸‍♂️ Power Rangers API

A **RESTful API** built with **Go (Golang)** that provides rich and structured information about the **Power Rangers franchise**, including seasons, Rangers, Zords, villains, movies, and actors.

## 🔥 About the Project

The Power Rangers franchise spans decades of seasons, teams, and characters. This API aims to organize that data in a consistent and accessible way for use in apps, websites, or educational projects.

## 🧩 Planned Endpoints

### 🏢 Franchise Details (`/api/v1/franchise`)
- Genre
- List of seasons
- Image URL for each season
- Icon/image for genre (e.g., “eternally red”)
- Official YouTube channel

### 📅 Season Details (`/api/v1/seasons/:id`)
- Title
- Subtitle
- Rangers
- Zords
- Villains
- Description / story of the season
- Key episodes
- Image URL of the season

### 🎥 Movies (`/api/v1/movies`)
- List of movies (titles, release years, synopsis, poster URLs, etc.)

### 🧑 Ranger / Actor (`/api/v1/rangers/:id`)
- Name
- Photo gallery
- Biography / background
- Fun facts

### 🎭 Actors (`/api/v1/actors/:id`)
- Name
- Profile photo
- Biography
- Curiosities and trivia

---

## 📦 Tech Stack

- **Go (Golang)**
- **Gin** – HTTP Web Framework
- **GORM** – ORM for interacting with the database
- **PostgreSQL** or **SQLite**
- **Swagger/OpenAPI** (optional) for docs

## 🚀 Getting Started

```bash
# Clone the repository
git clone https://github.com/your-username/power-rangers-api.git

# Enter the directory
cd power-rangers-api

# Install dependencies
go mod tidy

# Run the API
go run cmd/main.go


## Gogo Power Rangers