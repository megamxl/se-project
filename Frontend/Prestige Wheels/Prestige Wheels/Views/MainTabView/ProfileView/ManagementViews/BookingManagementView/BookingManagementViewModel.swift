//
//  BookingManagementViewModel.swift
//  Prestige Wheels
//
//  Created by Michael Luegmayer on 22.04.25.
//

import Foundation
import Combine
import OpenAPIClient

class BookingManagementViewModel: ObservableObject, ManagementViewModelProtocol {
    @Published var items: [OpenAPIClientAPI.Booking] = []
    @Published var isLoading = false
    @Published var errorMessage: String? = nil

    func fetchItems() {
        isLoading = true
        OpenAPIClientAPI.BookingAPI.getBookings(
            apiResponseQueue: DispatchQueue.main
        ) { [weak self] bookings, error in
            guard let self = self else { return }
            self.isLoading = false
            if let error = error {
                self.errorMessage = error.localizedDescription
            } else {
                self.items = bookings?.compactMap { $0 } ?? []
            }
        }
    }

    func delete(item: OpenAPIClientAPI.Booking) {
        guard let id = item.bookingId else { return }
        isLoading = true
        OpenAPIClientAPI.BookingAPI.deleteBooking(
            bookingId: id,
            apiResponseQueue: DispatchQueue.main
        ) { [weak self] _, _ in
            guard let self = self else { return }
            self.isLoading = false
            self.fetchItems()
        }
    }

    func add(request: OpenAPIClientAPI.BookCarRequest, completion: @escaping () -> Void) {
        isLoading = true
        OpenAPIClientAPI.BookingAPI.bookCar(
            bookCarRequest: request,
            apiResponseQueue: DispatchQueue.main
        ) { [weak self] _, _ in
            guard let self = self else { return }
            self.isLoading = false
            self.fetchItems()
            completion()
        }
    }
}
