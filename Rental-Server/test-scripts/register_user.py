import requests

BASE_URL = "http://localhost:80"

register_data = {
    "username": "john_doe",
    "email": "john@example.com",
    "password": "securePass123"
}

response = requests.post(f"{BASE_URL}/users", json=register_data)

if response.status_code == 201:
    print("✅ User registered successfully.")
elif response.status_code == 409:
    print("ℹ️  User already exists.")
else:
    print("❌ Failed to register user:", response.status_code, response.text)