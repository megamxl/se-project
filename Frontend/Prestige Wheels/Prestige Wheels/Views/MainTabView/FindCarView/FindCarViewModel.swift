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
    @Published var showDateValidationError: Bool = false
    
    @Published var fromDate: Date = Date()
    @Published var toDate: Date = Calendar.current.date(byAdding: .day, value: 1, to: Date()) ?? Date()
    
    @Published var selectedCurrency: OpenAPIClientAPI.Currency = .eur
    
    var isValidDateSelection: Bool {
        let today = Calendar.current.startOfDay(for: Date())
        let from = Calendar.current.startOfDay(for: fromDate)
        let to = Calendar.current.startOfDay(for: toDate)
        
        let minimumEndDate = Calendar.current.date(byAdding: .day, value: 1, to: from) ?? from
        
        return from >= today && from < to && to >= minimumEndDate
    }
    
    func listCars() {
        if !isValidDateSelection {
            showDateValidationError = true
            self.cars = [] 
            return
        }
        
        isLoading = true
        errorMessage = nil
        
        OpenAPIClientAPI.CarsAPI.listCars(
            currency: selectedCurrency,
            startTime: fromDate,
            endTime: toDate,
            apiResponseQueue: DispatchQueue.main
        ) { [weak self] (cars, error) in
            guard let self = self else { return }
            
            self.isLoading = false
            
            if let error {
                self.errorMessage = error.localizedDescription
                self.cars = []
            } else if let cars {
                self.cars = cars
                self.errorMessage = nil
            } else {
                self.cars = [] 
                self.errorMessage = "Unknown error occurred."
            }
        }
    }
}
