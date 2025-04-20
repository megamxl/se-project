# =========================
# 📊 Booking Service API Tests
# =========================

import unittest
import random
import string
import time
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

BOOKING_BASE_URL = "http://localhost:8098"
CAR_BASE_URL = "http://localhost:8098"
VERBOSE = "--no-output" not in sys.argv

def random_vin():
    return ''.join(random.choices(string.ascii_uppercase + string.digits, k=17))

def random_currency():
    return random.choice([
        "USD", "JPY", "BGN", "CZK", "DKK", "GBP", "HUF", "PLN", "RON", "SEK",
        "CHF", "ISK", "NOK", "TRY", "AUD", "BRL", "CAD", "CNY", "HKD", "IDR",
        "ILS", "INR", "KRW", "MXN", "MYR", "NZD", "PHP", "SGD", "THB", "ZAR", "EUR"
    ])

# ============================
# 📦 Booking Service Test Suite
# ============================

class BookingAPITest(unittest.TestCase):
    @classmethod
    def setUpClass(cls):
        print_test_header("🔧 Class Setup: Register + Login + Create Car")
        cls.currency = random_currency()
        cls.vin = random_vin()
        cls.session = login_and_get_session("john@example.com", "securePass123")
        cls.test_car = {
            "VIN": cls.vin,
            "model": "Model S",
            "brand": "Tesla",
            "imageURL": "http://example.com/car.jpg",
            "pricePerDay": 99.99
        }
        response = cls.session.post(f"{CAR_BASE_URL}/cars", json=cls.test_car)
        warn_if_500(response)
        print_test_footer(response.status_code, f"Pre-created VIN: {cls.vin}")
        print("Waiting for Pulsar to sync car to booking service...")
        time.sleep(1)  # Wait for Pulsar to sync car to booking service

    

    def setUp(self):
        self.session = self.__class__.session
        self.vin = self.__class__.vin
        self.currency = self.__class__.currency
        self.start_date = datetime.today().date().isoformat()
        self.end_date = (datetime.today() + timedelta(days=3)).date().isoformat()

    def tearDown(self):
        response = self.session.get(f"{BOOKING_BASE_URL}/booking")
        if response.status_code == 200:
            for booking in response.json():
                booking_id = booking.get("bookingId")
                if booking_id:
                    self.session.delete(f"{BOOKING_BASE_URL}/booking", params={"bookingId": booking_id})

    def test_01_book_car(self):
        print_test_header("🗕️ Book a Car")
        response = self.session.post(f"{BOOKING_BASE_URL}/booking", json={
            "VIN": self.vin,
            "currency": self.currency,
            "startTime": self.start_date,
            "endTime": self.end_date
        })
        status = warn_if_500(response)
        print_test_footer(status, f"Booking car {self.vin} from {self.start_date} to {self.end_date}")
        print_verbose_json(response, VERBOSE)
        self.assertEqual(status, 200)

    def test_02_get_bookings(self):
        print_test_header("📋 Get User Bookings")
        response = self.session.get(f"{BOOKING_BASE_URL}/booking")
        status = warn_if_500(response)
        print_test_footer(status, "Fetched all bookings for user")
        print_verbose_json(response, VERBOSE)
        self.assertEqual(status, 200)

    def test_03_update_booking_status(self):
        print_test_header("🔄 Update Booking Status")
        book_resp = self.session.post(f"{BOOKING_BASE_URL}/booking", json={
            "VIN": self.vin,
            "currency": self.currency,
            "startTime": self.start_date,
            "endTime": self.end_date
        })
        booking_id = book_resp.json().get("bookingId")
        response = self.session.put(f"{BOOKING_BASE_URL}/booking", json={
            "bookingId": booking_id,
            "status": "confirmed"
        })
        status = warn_if_500(response)
        print_test_footer(status, f"Updated booking {booking_id} to confirmed")
        print_verbose_text(response, VERBOSE)
        self.assertEqual(status, 200)

    def test_04_get_booking_by_id(self):
        print_test_header("🔍 Get Booking by ID")
        book_resp = self.session.post(f"{BOOKING_BASE_URL}/booking", json={
            "VIN": self.vin,
            "currency": self.currency,
            "startTime": self.start_date,
            "endTime": self.end_date
        })
        booking_id = book_resp.json().get("bookingId")
        response = self.session.get(f"{BOOKING_BASE_URL}/booking/{booking_id}")
        status = warn_if_500(response)
        print_test_footer(status, f"Fetched booking ID: {booking_id}")
        print_verbose_json(response, VERBOSE)
        self.assertEqual(status, 200)

    def test_05_delete_booking(self):
        print_test_header("❌ Delete Booking")
        book_resp = self.session.post(f"{BOOKING_BASE_URL}/booking", json={
            "VIN": self.vin,
            "currency": self.currency,
            "startTime": self.start_date,
            "endTime": self.end_date
        })
        booking_id = book_resp.json().get("bookingId")
        response = self.session.delete(f"{BOOKING_BASE_URL}/booking", params={"bookingId": booking_id})
        status = warn_if_500(response)
        print_test_footer(status, f"Deleted booking ID: {booking_id}")
        print_verbose_text(response, VERBOSE)
        self.assertEqual(status, 204)

    def test_06_book_with_invalid_date_range(self):
        print_test_header("🚫 Book with Invalid Date Range")
        response = self.session.post(f"{BOOKING_BASE_URL}/booking", json={
            "VIN": self.vin,
            "currency": self.currency,
            "startTime": self.end_date,
            "endTime": self.start_date
        })
        status = warn_if_500(response)
        print_test_footer(status, "Invalid date range booking attempt")
        print_verbose_text(response, VERBOSE)
        self.assertIn(status, [400, 422])

    def test_07_get_invalid_booking_id(self):
        print_test_header("❓ Get Booking with Invalid ID")
        response = self.session.get(f"{BOOKING_BASE_URL}/booking/invalid-id")
        status = warn_if_500(response)
        print_test_footer(status, "Invalid booking ID lookup")
        print_verbose_text(response, VERBOSE)
        self.assertIn(status, [400, 404])

    def test_08_book_car_duplicate(self):
        print_test_header("🗕️ Book a Car")
        response = self.session.post(f"{BOOKING_BASE_URL}/booking", json={
            "VIN": self.vin,
            "currency": self.currency,
            "startTime": self.start_date,
            "endTime": self.end_date
        })
        status = warn_if_500(response)
        print_test_footer(status, f"Booking car {self.vin} from {self.start_date} to {self.end_date}")
        print_verbose_json(response, VERBOSE)
        self.assertEqual(status, 500)


if __name__ == "__main__":
    unittest.main(argv=[arg for arg in sys.argv if arg != "--no-output"])
