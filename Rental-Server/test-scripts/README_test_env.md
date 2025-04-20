# 🧪 Car Rental Test Environment – Setup & Usage

This guide explains how to set up and run the Python-based test scripts for the Car Rental backend services (e.g., `carService`, `bookingService`).

---

## 📁 Directory Overview

```
test-scripts/
├── test_cars.py          # Car-related API tests
├── test_bookings.py      # Booking-related API tests
├── register_user.py      # Helper to create/register test users
├── test_util.py          # Shared login/session helper
├── test_output.py        # Shared output formatting
├── setup_test_env.sh     # Auto-setup script for Python env
└── venv/                 # Virtual environment (auto-created)
```

---

## 🚀 Getting Started

### 1. Navigate to the test script directory:

```bash
cd Rental-Server/test-scripts
```

### 2. Run the setup script (once per machine):

```bash
./setup_test_env.sh
```

This will:
- Create a `venv/`
- Install dependencies (e.g., `requests`)

### 3. Activate the environment:

```bash
source venv/bin/activate
```

> ⚠️ You must do this in every new terminal session.

---

## 🧪 Running Tests

```bash
# Run Car Service tests
python test_cars.py

# Run Booking Service tests
python test_bookings.py
```

### Optional: Suppress Extra Output

```bash
python test_cars.py --no-output
```

---

## 👤 Registering a User (Manual)

Use this script to register a test user:

```bash
python register_user.py
```

You can change the user credentials directly in the script before running it.

---

## 🔁 Re-Activating the Environment

Each time you open a new terminal:

```bash
cd Rental-Server/test-scripts
source venv/bin/activate
```

---

## 🛠️ Manual Dependency Management

To install dependencies manually:

```bash
pip install requests
```

To export installed packages:

```bash
pip freeze > requirements.txt
```

---

## 🧯 Common Issues

| Problem | Fix |
|--------|-----|
| `ModuleNotFoundError: No module named 'requests'` | Activate the environment: `source venv/bin/activate` |
| `zsh: command not found: python` | Use `python3` instead or re-run the setup script |

---