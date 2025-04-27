//
//  FindCarDetailView.swift
//  Prestige Wheels
//
//  Created by Heinz Schweitzer on 18.04.25.
//

import SwiftUI

struct FindCarDetailView: View {
    
    @EnvironmentObject var route: RouteObject
    @ObservedObject var viewModel: FindCarDetailViewModel
    
    var body: some View {
        VStack(alignment: .leading, spacing: 16) {
            AsyncImage(
                url: URL(string: viewModel.car.imageURL ?? ""),
                transaction: Transaction(animation: .default)
            ) { phase in
                switch phase {
                case .success(let image):
                    image
                        .resizable()
                        .scaledToFit()
                        .frame(maxWidth: .infinity)
                        .clipped()
                case .failure(_):
                    Image(systemName: "car.fill")
                        .resizable()
                        .scaledToFit()
                        .frame(maxWidth: .infinity)
                        .foregroundColor(.gray)
                        .padding()
                default:
                    ProgressView()
                        .frame(maxWidth: .infinity)
                }
            }
            
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
                .foregroundStyle(.primary)
                .font(.title2)
                .fontDesign(.rounded)
            
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
                .padding(.vertical)

            Text("Car Details")
                .font(.title3)
                .fontWeight(.semibold)
                .fontDesign(.rounded)
                .padding(.bottom, 4)

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

            Spacer()
        }
        .padding()
        .ignoresSafeArea(edges: .top)
        .alert("Prestige Wheels", isPresented: $viewModel.showAlert) {
            Button("OK", role: .cancel) {
                route.path.removeLast()
            }
        } message: {
            Text(viewModel.alertMessage)
        }
    }
}

#Preview {
    FindCarDetailView(viewModel: .init(car: .init(), currency: .eur, from: .now, to: .now))
}
