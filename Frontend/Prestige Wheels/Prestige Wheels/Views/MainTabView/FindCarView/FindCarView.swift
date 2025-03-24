//
//  FindCarView.swift
//  Prestige Wheels
//
//  Created by Michael Luegmayer on 24.03.25.
//

import SwiftUI

struct FindCarView: View {
    var body: some View {
        NavigationStack{
            VStack {
                HStack(spacing: 8){
                    DatePicker(
                        selection: .constant(Date()),
                        displayedComponents: .date
                    ) {
                        Text("From:")
                    }
                    DatePicker(
                        selection: .constant(Date()),
                        displayedComponents: .date
                    ) {
                        Text("To:")
                    }
                }
                .padding()
                
                List(1..<10) { _ in
                    Text("Test\nTest")
                }
            }
            .navigationTitle("Prestige Wheels")
        }
    }
}

#Preview {
    FindCarView()
}
