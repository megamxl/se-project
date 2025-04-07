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
                        .vAlign(.center)
                } else if let error = viewModel.errorMessage {
                    ContentUnavailableView("API Error", systemImage: "exclamationmark.triangle.fill", description: Text(error))
                } else {
                    List(viewModel.cars, id: \.VIN) { car in
                        CarRow(car: car, currency: viewModel.selectedCurrency)
                            .listRowSeparator(.hidden)
                            .listRowInsets(.init())
                            .padding(.vertical, .spacingXS)
                            .padding(.horizontal)
                    }
                    .listStyle(.plain)
                }
            }
            .toolbar {
                ToolbarItem(placement: .topBarTrailing) {
                    Menu {
                        ForEach(OpenAPIClientAPI.Currency.allCases, id: \.self) { currency in
                            Button {
                                viewModel.selectedCurrency = currency
                            } label: {
                                Text(currency.rawValue)
                            }
                        }
                    } label: {
                        Text(currencySymbol(for: viewModel.selectedCurrency.rawValue))
                            .font(.title2)
                            .fontWeight(.semibold)
                            .fontDesign(.rounded)
                    }
                }
            }
            .navigationTitle("Prestige Wheels")
            .onAppear {
                if LoginViewModel.shared.isLoggedIn {
                    viewModel.loadCars()
                }
            }
            .refreshable {
                if LoginViewModel.shared.isLoggedIn {
                    print("loading cars")
                    viewModel.loadCars()
                } else {
                    print("not logged in")
                }
            }
        }
    }
    
    func currencySymbol(for currencyCode: String) -> String {
        let formatter = NumberFormatter()
        formatter.numberStyle = .currency
        formatter.currencyCode = currencyCode
        return formatter.currencySymbol ?? currencyCode
    }
}

#Preview {
    FindCarView()
}
