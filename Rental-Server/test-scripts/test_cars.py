# =========================
# 🧪 Car Service API Tests
# =========================

import unittest
import random
import string
from datetime import datetime, timedelta
import sys
from test_util import login_and_get_session
from test_output import (
    print_test_header,
    print_test_footer,
    warn_if_500,
    print_verbose_json,
    print_verbose_text
)

# =====================
# 🔧 Global Test Config
# =====================

def parse_base_url(default="http://localhost:9098"):
    for i, arg in enumerate(sys.argv):
        if arg == "--base-url" and i + 1 < len(sys.argv):
            return sys.argv[i + 1]
    return default

BASE_URL = parse_base_url()
VERBOSE = "--no-output" not in sys.argv

def random_vin():
    return ''.join(random.choices(string.ascii_uppercase + string.digits, k=17))

VIN = random_vin()

def random_currency():
    return random.choice([
        "USD", "JPY", "BGN", "CZK", "DKK", "GBP", "HUF", "PLN", "RON", "SEK",
        "CHF", "ISK", "NOK", "TRY", "AUD", "BRL", "CAD", "CNY", "HKD", "IDR",
        "ILS", "INR", "KRW", "MXN", "MYR", "NZD", "PHP", "SGD", "THB", "ZAR", "EUR"
    ])

# ============================
# 📦 Car Service Test Suite
# ============================

class CarAPITest(unittest.TestCase):

    # ---------------------
    # 🔧 Setup before tests
    # ---------------------
    def setUp(self):
        self.vin = VIN
        self.currency = random_currency()
        self.start_date = datetime.today().date().isoformat()
        self.end_date = (datetime.today() + timedelta(days=3)).date().isoformat()
        self.test_car = {
            "VIN": self.vin,
            "model": "Model S",
            "brand": "Tesla",
            "imageURL": "http://example.com/car.jpg",
            "pricePerDay": 99.99
        }
        self.session = login_and_get_session("john@example.com", "securePass123" , login_url=BASE_URL+"/login")

    # -----------------
    # ✅ Main Test Cases
    # -----------------

    def test_01_add_car(self):
        print_test_header("Add New Car")
        response = self.session.post(f"{BASE_URL}/cars", json=self.test_car)
        status = warn_if_500(response)
        print_test_footer(status, f"VIN: {self.vin}")
        print_verbose_json(response, VERBOSE)
        self.assertEqual(status, 201)

    def test_02_get_cars_with_currency(self):
        print_test_header("Get Cars by Currency")
        response = self.session.get(f"{BASE_URL}/cars", params={"currency": self.currency})
        status = warn_if_500(response)
        print_test_footer(status, f"Currency: {self.currency}")
        print_verbose_json(response, VERBOSE)
        self.assertEqual(status, 200)

    def test_03_get_cars_with_all_params(self):
        print_test_header("Get Cars by Time Range and Currency")
        response = self.session.get(f"{BASE_URL}/cars", params={
            "currency": self.currency,
            "startTime": self.start_date,
            "endTime": self.end_date
        })
        status = warn_if_500(response)
        print_test_footer(status, f"{self.start_date} → {self.end_date} in {self.currency}")
        print_verbose_json(response, VERBOSE)
        self.assertEqual(status, 200)

    def test_04_update_car(self):
        print_test_header("Update Car Price")
        self.test_car["pricePerDay"] = 120.0
        print(self.test_car)
        response = self.session.put(f"{BASE_URL}/cars", json=self.test_car)
        status = warn_if_500(response)
        print_test_footer(status, f"Updated price: {self.test_car['pricePerDay']}")
        print_verbose_json(response, VERBOSE)
        self.assertEqual(status, 200)

    def test_05_delete_car(self):
        print_test_header("Delete Car")
        self.session.post(f"{BASE_URL}/cars", json=self.test_car)
        response = self.session.delete(f"{BASE_URL}/cars", params={"VIN": self.vin})
        status = warn_if_500(response)
        print_test_footer(status, f"Deleted VIN: {self.vin}")
        print_verbose_text(response, VERBOSE)
        self.assertEqual(status, 204)

    def test_06_add_invalid_car(self):
        print_test_header("Add Invalid Car (Missing Fields)")
        response = self.session.post(f"{BASE_URL}/cars", json={"model": "X"})
        status = warn_if_500(response)
        print_test_footer(status, "Attempted to add incomplete car object")
        print_verbose_text(response, VERBOSE)
        self.assertIn(status, [400, 422])

    def test_07_get_nonexistent_car(self):
        print_test_header("Get Nonexistent Car by VIN")
        response = self.session.get(f"{BASE_URL}/cars", params={"VIN": "NONEXISTENTVIN"})
        status = warn_if_500(response)
        print_test_footer(status, "Tried to get nonexistent VIN")
        print_verbose_json(response, VERBOSE)
        self.assertIn(status, [404, 400])

    def test_08_update_nonexistent_car(self):
        print_test_header("Update Nonexistent Car")
        fake_car = self.test_car.copy()
        fake_car["VIN"] = "NONEXISTENTVIN"
        response = self.session.put(f"{BASE_URL}/cars", json=fake_car)
        status = warn_if_500(response)
        print_test_footer(status, "Update attempt on nonexistent VIN")
        print_verbose_text(response, VERBOSE)
        self.assertEqual(status, 500)

    def test_09_delete_nonexistent_car(self):
        print_test_header("Delete Nonexistent Car")
        response = self.session.delete(f"{BASE_URL}/cars", params={"VIN": "NONEXISTENTVIN"})
        status = warn_if_500(response)
        print_test_footer(status, "Delete attempt on nonexistent VIN")
        print_verbose_text(response, VERBOSE)
        self.assertEqual(status, 500)

    def test_10_add_duplicate_vin(self):
        print_test_header("Add Duplicate VIN")
        self.session.post(f"{BASE_URL}/cars", json=self.test_car)
        response = self.session.post(f"{BASE_URL}/cars", json=self.test_car)
        status = warn_if_500(response)
        print_test_footer(status, f"Tried adding VIN twice: {self.vin}")
        print_verbose_text(response, VERBOSE)
        self.assertEqual(status, 500)

# ==================
# 🚀 Run the test suite
# ==================

if __name__ == "__main__":
    filtered_args = []
    skip_next = False
    for i, arg in enumerate(sys.argv):
        if skip_next:
            skip_next = False
            continue
        if arg == "--base-url":
            skip_next = True  # Skip the next arg (URL)
            continue
        filtered_args.append(arg)

    unittest.main(argv=filtered_args)