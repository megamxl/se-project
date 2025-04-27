//
//  CarManagementViewModel.swift
//  Prestige Wheels
//
//  Created by Michael Luegmayer on 22.04.25.
//

import Foundation
import Combine
import OpenAPIClient

class CarManagementViewModel: ObservableObject, ManagementViewModelProtocol {
    @Published var items: [OpenAPIClientAPI.CarListInner] = []
    @Published var isLoading = false
    @Published var errorMessage: String? = nil

    func fetchItems() {
        isLoading = true
        OpenAPIClientAPI.CarsAPI.listCars(
            currency: .eur,
            startTime: nil,
            endTime: nil,
            apiResponseQueue: DispatchQueue.main
        ) { [weak self] cars, error in
            guard let self = self else { return }
            self.isLoading = false
            if let error = error {
                self.errorMessage = error.localizedDescription
            } else {
                self.items = cars?.compactMap { $0 } ?? []
            }
        }
    }

    func delete(item: OpenAPIClientAPI.CarListInner) {
        guard let vin = item.VIN else { return }
        isLoading = true
        OpenAPIClientAPI.CarsAPI.deleteCar(
            VIN: vin,
            apiResponseQueue: DispatchQueue.main
        ) { [weak self] _, _ in
            guard let self = self else { return }
            self.isLoading = false
            self.fetchItems()
        }
    }

    func add(car: OpenAPIClientAPI.Car, completion: @escaping () -> Void) {
        isLoading = true
        OpenAPIClientAPI.CarsAPI.addCar(
            car: car,
            apiResponseQueue: DispatchQueue.main
        ) { [weak self] _, _ in
            guard let self = self else { return }
            self.isLoading = false
            self.fetchItems()
            completion()
        }
    }
}
