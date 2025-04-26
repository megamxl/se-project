//
//  FindCarDetailView.swift
//  Prestige Wheels
//
//  Created by Heinz Schweitzer on 18.04.25.
//

import SwiftUI
import MapKit

struct FindCarDetailView: View {
    
    @EnvironmentObject var route: RouteObject
    @ObservedObject var viewModel: FindCarDetailViewModel
    
    var body: some View {
        ScrollView {
            VStack(spacing: 0) {
                AsyncImage(
                    url: URL(string: viewModel.car.imageURL ?? ""),
                    transaction: Transaction(animation: .default)
                ) { phase in
                    switch phase {
                    case .success(let image):
                        image
                            .resizable()
                            .aspectRatio(contentMode: .fit)
                            .cornerRadius(20)
                    case .failure(_):
                        Image(systemName: "car.fill")
                            .resizable()
                            .scaledToFit()
                            .frame(height: 250)
                            .frame(maxWidth: .infinity)
                            .foregroundColor(.gray)
                            .background(Color(.systemGray5))
                            .clipped()
                    default:
                        ProgressView()
                            .frame(height: 250)
                            .frame(maxWidth: .infinity)
                    }
                }
                .ignoresSafeArea(.container, edges: .top)
                
                VStack(alignment: .leading, spacing: 16) {
                    
                    VStack(alignment: .leading, spacing: 0) {
                        HStack {
                            Text(viewModel.car.model ?? "")
                                .foregroundStyle(.primary)
                                .font(.title)
                                .fontWeight(.bold)
                                .fontDesign(.rounded)
                                .frame(maxWidth: .infinity, alignment: .leading)
                            
                            if let priceOverAll = viewModel.car.priceOverAll {
                                Text("\(priceOverAll.formatted(.currency(code: viewModel.currency.rawValue)))")
                                    .foregroundStyle(.primary)
                                    .font(.title2)
                                    .fontWeight(.bold)
                                    .fontDesign(.rounded)
                            }
                        }
                        
                        Text(viewModel.car.brand ?? "")
                            .foregroundStyle(.secondary)
                            .font(.headline)
                            .fontDesign(.rounded)
                    }
                    
                    HStack {
                        Label("From:", systemImage: "calendar")
                            .fontWeight(.semibold)
                        Text(viewModel.from.formatted(date: .abbreviated, time: .omitted))
                            .frame(maxWidth: .infinity, alignment: .trailing)
                    }
                    
                    HStack {
                        Label("To:", systemImage: "calendar")
                            .fontWeight(.semibold)
                        Text(viewModel.to.formatted(date: .abbreviated, time: .omitted))
                            .frame(maxWidth: .infinity, alignment: .trailing)
                    }
                    
                    Button {
                        viewModel.bookCar()
                    } label: {
                        Text("Book this Car")
                            .font(.headline)
                            .fontDesign(.rounded)
                            .foregroundColor(.white)
                            .frame(maxWidth: .infinity)
                            .padding()
                            .background(Color.black)
                            .cornerRadius(12)
                            .shadow(color: .black.opacity(0.2), radius: 8, x: 0, y: 4)
                    }
                    
                    Divider()
                    
                    Text("Car Details")
                        .font(.title3)
                        .fontWeight(.semibold)
                        .fontDesign(.rounded)
                    
                    VStack(alignment: .leading, spacing: 8) {
                        if let vin = viewModel.car.VIN {
                            HStack {
                                Text("VIN:")
                                    .fontWeight(.semibold)
                                Text(vin)
                                    .frame(maxWidth: .infinity, alignment: .trailing)
                            }
                        }
                        
                        if let brand = viewModel.car.brand {
                            HStack {
                                Text("Brand:")
                                    .fontWeight(.semibold)
                                Text(brand)
                                    .frame(maxWidth: .infinity, alignment: .trailing)
                            }
                        }
                        
                        if let model = viewModel.car.model {
                            HStack {
                                Text("Model:")
                                    .fontWeight(.semibold)
                                Text(model)
                                    .frame(maxWidth: .infinity, alignment: .trailing)
                            }
                        }
                        
                        if let pricePerDay = viewModel.car.pricePerDay {
                            HStack {
                                Text("Price per Day:")
                                    .fontWeight(.semibold)
                                Text("\(pricePerDay.formatted(.currency(code: viewModel.currency.rawValue)))")
                                    .frame(maxWidth: .infinity, alignment: .trailing)
                            }
                        }
                        
                        if let priceOverAll = viewModel.car.priceOverAll {
                            HStack {
                                Text("Total Price:")
                                    .fontWeight(.semibold)
                                Text("\(priceOverAll.formatted(.currency(code: viewModel.currency.rawValue)))")
                                    .frame(maxWidth: .infinity, alignment: .trailing)
                            }
                        }
                    }
                    .font(.body)
                    .fontDesign(.rounded)
                    .padding()
                    .background(Color(.systemGray6))
                    .cornerRadius(12)
                    
                    Divider()
                    
                    Text("Car Pickup")
                        .font(.title3)
                        .fontWeight(.semibold)
                        .fontDesign(.rounded)
                    
                    if viewModel.selectedMapProvider == .apple {
                        Map(position: $viewModel.position) {
                            Marker("FH Campus Wien", coordinate: CLLocationCoordinate2D(latitude: 48.157975, longitude: 16.381778))
                        }
                        .frame(height: 200)
                        .cornerRadius(12)
                    } else {
                        Text("Google Maps")
                    }
                    
                    Spacer()
                }
                .padding()
            }
        }
        .background(Color(.systemBackground))
        .alert("Prestige Wheels", isPresented: $viewModel.showAlert) {
            Button("OK", role: .cancel) {
                route.pathFindCar.removeLast()
            }
        } message: {
            Text(viewModel.alertMessage)
        }
        .edgesIgnoringSafeArea(.top)
        .toolbarBackground(.hidden, for: .navigationBar)
        .toolbar {
            ToolbarItem(placement: .navigationBarLeading) {
                Button(action: {
                    route.pathFindCar.removeLast()
                }) {
                    ZStack {
                        Color.gray
                            .frame(width: 30, height: 30, alignment: .center)
                            .cornerRadius(5)
                        
                        Image(systemName: "chevron.left")
                            .foregroundColor(.black)
                    }
                }
            }
        }
        .navigationBarBackButtonHidden(true)
    }
}

#Preview {
    FindCarDetailView(viewModel: .init(car: .init(), currency: .eur, from: .now, to: .now))
}

