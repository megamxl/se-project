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
    let from: Date
    let to: Date
    
    public init(car: OpenAPIClientAPI.CarListInner,
                currency: OpenAPIClientAPI.Currency,
                from: Date,
                to: Date) {
        self.car = car
        self.currency = currency
        self.from = from
        self.to = to
    }
}
