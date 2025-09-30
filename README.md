# Guilliman

Guilliman is a personal finance backend built in Go.  
It helps track and analyze monthly expenses, with a focus on the **50/30/20 rule** and future plans for **AI/ML-driven recommendations** on spending, saving, and investing.

---

## ✨ Features
- Expense tracking (backend API, paired with a React Native frontend).
- Budget analysis based on the **50% needs / 30% wants / 20% savings** model.
- PostgreSQL database for reliable storage.
- Container-ready with Docker/Podman.
- Roadmap: personalized AI/ML recommendations and investment feedback.

---

## 🛠️ Tech Stack
- **Backend**: Go + Gin
- **Database**: PostgreSQL
- **Containerization**: Docker / Podman
- **Frontend**: React Native (separate repository)

---

## 🚀 Getting Started

### Requirements
- Docker or Podman
- Docker Compose or Podman Compose

### Quickstart
Clone the repository and run:

```bash
docker-compose up -d
```

This will start the backend service and a PostgreSQL database.
The entrypoint script entrypoint.sh is included for convenience in container runs.

---

## 📱 Frontend

The React Native app that consumes this backend lives in a separate repository.

---

## 📖 Roadmap
	•	Expense tracking API
	•	AI/ML-powered recommendations
	•	Investment insights and feedback

## 📄 License

This project is licensed under the MIT License. See LICENSE for details.
