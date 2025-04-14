import unittest
import requests
import random
import string
from datetime import datetime, timedelta

BASE_URL = "http://localhost:80"  # change as needed

def random_vin():
    return ''.join(random.choices(string.ascii_uppercase + string.digits, k=17))

def random_currency():
    return random.choice([
        "USD", "JPY", "EUR", "GBP", "AUD", "CAD", "CHF", "CNY", "SEK", "NZD"
    ])

class CarBookingAPITest(unittest.TestCase):

    def setUp(self):
        self.vin = random_vin()
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
        self.headers = {
            "Authorization": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NDQ1Njg5MTYsIm5hbWUiOiJtQG0uYXQiLCJyb2xlcyI6ImN1c3RvbWVyIiwic3ViIjoiZDA2ZGY3ZTktZDEwNS00M2I3LTkxZjUtNjRjZWE1YTBiYTBlIn0.v-_d3Zsh5wJePcte-zA2_pJwnShSSbMOoevOmwwjP4I",
            "Content-Type": "application/json"
        }

    def test_01_add_car(self):
        response = requests.post(f"{BASE_URL}/cars", json=self.test_car, headers=self.headers)
        self.assertEqual(response.status_code, 201)

    def test_02_get_cars_with_currency(self):
        response = requests.get(f"{BASE_URL}/cars", params={"currency": self.currency}, headers=self.headers)
        self.assertEqual(response.status_code, 200)

    def test_03_get_cars_with_all_params(self):
        response = requests.get(f"{BASE_URL}/cars", params={
            "currency": self.currency,
            "startTime": self.start_date,
            "endTime": self.end_date
        }, headers=self.headers)
        self.assertEqual(response.status_code, 200)

    def test_04_update_car(self):
        self.test_car["pricePerDay"] = 120.0
        response = requests.put(f"{BASE_URL}/cars", json=self.test_car, headers=self.headers)
        self.assertEqual(response.status_code, 200)

    def test_05_book_car(self):
        response = requests.post(f"{BASE_URL}/booking", json={
            "VIN": self.vin,
            "currency": self.currency,
            "startTime": self.start_date,
            "endTime": self.end_date
        }, headers=self.headers)
        self.assertEqual(response.status_code, 201)
        self.booking = response.json()
        self.booking_id = self.booking.get("bookingId")

    def test_06_get_user_bookings(self):
        response = requests.get(f"{BASE_URL}/booking", headers=self.headers)
        self.assertEqual(response.status_code, 200)

    def test_07_update_booking(self):
        response = requests.put(f"{BASE_URL}/booking", json={
            "bookingId": self.booking_id,
            "status": "confirmed"
        })
        self.assertEqual(response.status_code, 200)

    def test_08_get_booking_by_id(self):
        response = requests.get(f"{BASE_URL}/booking/{self.booking_id}", headers=self.headers)
        self.assertEqual(response.status_code, 200)

    def test_09_get_all_bookings_admin(self):
        response = requests.get(f"{BASE_URL}/bookings/all/", headers=self.headers)
        self.assertIn(response.status_code, [200, 400])  # 400 for edge case

    def test_10_delete_booking(self):
        response = requests.delete(f"{BASE_URL}/booking", params={"bookingId": self.booking_id}, headers=self.headers)
        self.assertEqual(response.status_code, 204)

    def test_11_delete_car(self):
        response = requests.delete(f"{BASE_URL}/cars", params={"VIN": self.vin}, headers=self.headers)
        self.assertEqual(response.status_code, 204)

    # Edge Cases
    def test_12_add_invalid_car(self):
        response = requests.post(f"{BASE_URL}/cars", json={"model": "X"}, headers=self.headers)
        self.assertIn(response.status_code, [400, 422])

    def test_13_book_nonexistent_car(self):
        response = requests.post(f"{BASE_URL}/booking", json={
            "VIN": "INVALIDVIN0000000",
            "currency": self.currency,
            "startTime": self.start_date,
            "endTime": self.end_date
        })
        self.assertIn(response.status_code, [400, 404])

    def test_14_get_booking_invalid_id(self):
        response = requests.get(f"{BASE_URL}/booking/invalid-id", headers=self.headers)
        self.assertIn(response.status_code, [400, 404])

    def test_15_delete_nonexistent_booking(self):
        response = requests.delete(f"{BASE_URL}/booking", params={"bookingId": "nonexistent"}, headers=self.headers)
        self.assertIn(response.status_code, [400, 500])

if __name__ == "__main__":
    unittest.main()