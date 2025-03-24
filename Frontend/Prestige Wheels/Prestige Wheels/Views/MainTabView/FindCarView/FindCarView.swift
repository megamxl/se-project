//
//  FindCarView.swift
//  Prestige Wheels
//
//  Created by Michael Luegmayer on 24.03.25.
//

import SwiftUI
import OpenAPIClient

struct FindCarView: View {
    @StateObject private var viewModel = FindCarViewModel()
    
    var body: some View {
        NavigationStack {
            VStack {
                // MARK: Header
                HStack(spacing: 8) {
                    DatePicker("From:", selection: $viewModel.fromDate, displayedComponents: .date)
                    DatePicker("To:", selection: $viewModel.toDate, displayedComponents: .date)
                }
                .padding()
                
                // MARK: Content
                if viewModel.isLoading {
                    ProgressView("Loading Cars...")
                        .padding()
                } else if let error = viewModel.errorMessage {
                    ContentUnavailableView("API Error", systemImage: "exclamationmark.triangle.fill", description: Text(error))
                } else {
                    List(viewModel.cars, id: \.VIN) { car in
                        VStack(alignment: .leading) {
                            Text(car.brand ?? "") // Annahme: Car hat die Properties 'brand' und 'model'
                                .font(.headline)
                            Text(car.model ?? "")
                                .font(.subheadline)
                        }
                    }
                }
            }
            .navigationTitle("Prestige Wheels")
            .onAppear {
                viewModel.loadCars()
            }
        }
    }
}

#Preview {
    FindCarView()
}
