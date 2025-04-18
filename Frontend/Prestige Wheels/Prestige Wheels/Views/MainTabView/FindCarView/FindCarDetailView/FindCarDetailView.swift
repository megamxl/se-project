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
        Text(viewModel.car.brand ?? "")
    }
}

#Preview {
    FindCarDetailView(viewModel: .init(car: .init()))
}
