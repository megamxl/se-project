//
//  FindCarDetailViewModel.swift
//  Prestige Wheels
//
//  Created by Heinz Schweitzer on 18.04.25.
//

import Foundation
import OpenAPIClient

class FindCarDetailViewModel: ObservableObject {
    
    let car: OpenAPIClientAPI.CarListInner
    let currency: OpenAPIClientAPI.Currency
    
    public init(car: OpenAPIClientAPI.CarListInner, currency: OpenAPIClientAPI.Currency) {
        self.car = car
        self.currency = currency
    }
}
