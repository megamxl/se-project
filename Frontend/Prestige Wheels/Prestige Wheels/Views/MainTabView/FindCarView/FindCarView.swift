//
//  FindCarView.swift
//  Prestige Wheels
//
//  Created by Michael Luegmayer on 24.03.25.
//

import SwiftUI
import OpenAPIClient
import OSLog

struct FindCarView: View {
    @StateObject private var viewModel = FindCarViewModel()
    @EnvironmentObject var authenticationViewModel: AuthenticationViewModel
    @EnvironmentObject var route: RouteObject
    
    var body: some View {
        NavigationStack {
            VStack {
                // MARK: Header
                HStack(spacing: 8) {
                    DatePicker("From:", selection: $viewModel.fromDate, displayedComponents: .date)
                    DatePicker("To:", selection: $viewModel.toDate, displayedComponents: .date)
                }
                .padding()
                .background(Color.secondary.opacity(0.25))
                
                // MARK: Content
                if viewModel.isLoading {
                    ProgressView("Loading Cars...")
                        .vAlign(.center)
                }
                else if viewModel.errorMessage != nil || viewModel.cars.isEmpty {
                    ContentUnavailableView("No Result", systemImage: "car.2.fill", description: Text("No cars are available!"))
                } else {
                    List(viewModel.cars, id: \.VIN) { car in
                        Button {
                            route.pathFindCar.append(.findCarDetailView(car: car,
                                                                        currency: viewModel.selectedCurrency,
                                                                        from: viewModel.fromDate,
                                                                        to: viewModel.toDate))
                        } label: {
                            CarRow(car: car, currency: viewModel.selectedCurrency)
                                .listRowSeparator(.hidden)
                                .listRowInsets(.init())
                                .padding(.vertical, .spacingXS)
                        }
                        .listRowSeparator(.hidden)
                    }
                    .listStyle(.plain)
                    .contentMargins(.top, 0, for: .scrollContent)
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
                viewModel.listCars()
            }
            .onChange(of: authenticationViewModel.isLoggedIn) { isLoggedIn in
                if isLoggedIn {
                    Logger.backgroundProcessing.log("✅ Login success, loading cars")
                    viewModel.listCars()
                }
            }
            .onChange(of: viewModel.selectedCurrency) {
                Logger.backgroundProcessing.log("🔄 refresh")
                viewModel.listCars()
            }
            .onChange(of: viewModel.toDate) {
                Logger.backgroundProcessing.log("🔄 refresh")
                viewModel.listCars()
            }
            .onChange(of: viewModel.fromDate) {
                Logger.backgroundProcessing.log("🔄 refresh")
                viewModel.listCars()
            }
            .refreshable {
                Logger.backgroundProcessing.log("🔄 loading cars")
                viewModel.listCars()
            }
            .alert("Invalid Booking Dates", isPresented: $viewModel.showDateValidationError) {
                Button("OK", role: .cancel) { }
            } message: {
                Text("Please ensure the start date is today or later, before the end date, and at least 1 day difference.")
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
        .environmentObject(AuthenticationViewModel())
}
