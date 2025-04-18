//
//  FindCarViewModel.swift
//  Prestige Wheels
//
//  Created by Michael Luegmayer on 25.03.25.
//

import Foundation
import OpenAPIClient

class FindCarViewModel: ObservableObject {
    @Published var cars: [OpenAPIClientAPI.CarListInner] = []
    @Published var isLoading: Bool = false
    @Published var errorMessage: String?
    
    @Published var fromDate: Date = Date()
    @Published var toDate: Date = Date()
    
    @Published var selectedCurrency: OpenAPIClientAPI.Currency = .eur
    
    func listCars() {
        isLoading = true
        
        OpenAPIClientAPI.CarsAPI.listCars(
            currency: OpenAPIClientAPI.Currency.eur,
            startTime: fromDate,
            endTime: toDate,
            apiResponseQueue: DispatchQueue.main
        ) { [weak self] (cars, error) in
            
            guard let self = self else { return }
            
            self.isLoading = false
            
            if let error = error {
                self.errorMessage = error.localizedDescription
            } else if let cars = cars {
                self.cars = cars
            }
        }
    }
    
    func addCar() {
        
    }
    
    func deleteCar() {
        
    }
    
    func updateCar() {
        
    }
}
