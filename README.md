# Healthcare CRM

Monorepo hub for the Healthcare CRM system — an internal platform for managing patients, appointments, and exams. This project was developed as part of the Software Engineering course at Universidad del Valle de Guatemala.

**Note:** The code find inside of this repository is a re-write / port of a previous version of the project found [here](https://github.com/AngelEsquit/CRM). This re-write was done
in order to clean up infrastructure, migrate to utilizing a component library & improve code quality overall.

## Project Structure

This project uses Github Submodules to manage Frontend, Backend & Database. In the future, Go-Migration will be used to manage database versions instead of individual SQL scripts.

## How to Run (Development)

### 1. Clone the repository (SSH)

Clone via SSH
```bash
git clone git@github.com:TonitoMC/healthcare-crm.git
cd healthcare-crm
```

Update Submodules
```bash
git submodule update --init --recursive
```

Run with Docker
```bash
docker compose -f dockerfile.dev.yml up --build
```

Open in browser
```bash
http://localhost:5173
```
