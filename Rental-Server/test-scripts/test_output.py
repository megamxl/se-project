# =======================
# 📤 Output Helper Module
# =======================

import json

def print_test_header(title: str):
    print("")
    print("\n" + "=" * 60)
    print(f"🔹 {title}")
    print("=" * 60)

def print_test_footer(status_code: int, extra: str = ""):
    print(f"✅ Status Code: {status_code} — {extra}")
    print("-" * 60)

def warn_if_500(response):
    if response.status_code == 500:
        print("⚠️  Got 500 Internal Server Error — check backend error handling.")
    return response.status_code

def print_verbose_json(response, verbose: bool):
    if verbose:
        try:
            print("📝 Response:", json.dumps(response.json(), indent=2))
        except Exception:
            print("❌ Failed to print JSON response.")

def print_verbose_text(response, verbose: bool):
    if verbose:
        try:
            print("📝 Response:", response.text)
        except Exception:
            print("❌ Failed to print text response.")

def safe_json_parse(response):
    try:
        return response.json()
    except Exception:
        return None