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
    
    public init(car: OpenAPIClientAPI.CarListInner) {
        self.car = car
    }
}
