//
//  CarCreateView.swift
//  Prestige Wheels
//
//  Created by Michael Luegmayer on 22.04.25.
//

import SwiftUI
import OpenAPIClient

struct CarCreateView: View {
    @Environment(\.dismiss) private var dismiss
    @ObservedObject var viewModel: CarManagementViewModel
    @State private var VIN = ""
    @State private var model = ""
    @State private var brand = ""
    @State private var imageURL = ""
    @State private var pricePerDay = ""

    var body: some View {
        NavigationStack {
            Form {
                TextField("VIN", text: $VIN)
                TextField("Model", text: $model)
                TextField("Brand", text: $brand)
                TextField("Image URL", text: $imageURL)
                TextField("Price per day", text: $pricePerDay)
                    .keyboardType(.decimalPad)
            }
            .navigationTitle("Add Car")
            .toolbar {
                ToolbarItem(placement: .cancellationAction) {
                    Button("Cancel") { dismiss() }
                }
                ToolbarItem(placement: .confirmationAction) {
                    Button("Save") {
                        guard let price = Double(pricePerDay) else { return }
                        let car = OpenAPIClientAPI.Car(
                            VIN: VIN,
                            model: model,
                            brand: brand,
                            imageURL: imageURL,
                            pricePerDay: price
                        )
                        viewModel.add(car: car) { dismiss() }
                    }
                }
            }
        }
    }
}

struct CarCreateView_Previews: PreviewProvider {
    static var previews: some View {
        CarCreateView(viewModel: CarManagementViewModel())
    }
}
