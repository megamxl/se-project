import requests

def login_and_get_session(email, password, login_url="http://localhost:8091/login"):
    session = requests.Session()
    response = session.post(login_url, json={"email": email, "password": password})
    if response.status_code == 200 and 'jwt' in session.cookies:
        return session
    raise Exception(f"❌ Login failed: {response.status_code} - {response.text}")