import requests

BASE_URL = "http://localhost:80"

# Step 1: Register a new user
register_data = {
    "username": "john_doe",
    "email": "john@example.com",
    "password": "securePass123"
}

register_response = requests.post(f"{BASE_URL}/users", json=register_data)

if register_response.status_code == 201:
    print("✅ User registered successfully.")
else:
    print("❌ Failed to register user:", register_response.text)
    exit(1)

# Step 2: Login
login_data = {
    "email": register_data["email"],
    "password": register_data["password"]
}

session = requests.Session()  # use session to keep cookies

login_response = session.post(f"{BASE_URL}/login", json=login_data)

if login_response.status_code == 200 and 'jwt' in session.cookies:
    jwt_token = session.cookies.get('jwt')
    print("✅ Logged in successfully. JWT received via cookie.")
else:
    print("❌ Login failed or missing JWT cookie:", login_response.text)
    exit(1)

# Step 3: Get user info
user_info_response = session.get(f"{BASE_URL}/users")

if user_info_response.status_code == 200:
    user_info = user_info_response.json()
    print("👤 User Info:", user_info)
else:
    print("❌ Failed to fetch user info:", user_info_response.text)
    exit(1)

# Step 4: Update user email
updated_username = "john webber"
update_data = {
    "username": updated_username,
    "email": register_data["email"],
    "password": register_data["password"]
}

update_response = session.put(f"{BASE_URL}/users/all", json=update_data)

if update_response.status_code == 200:
    print(f"✅ Email updated ")
else:
    print("❌ Failed to update user:", update_response.text)
    exit(1)

exit(0)