//
//  AdminManagementView.swift
//  Prestige Wheels
//
//  Created by Michael Luegmayer on 22.04.25.
//

import SwiftUI
import Foundation

/// Protocol for any management ViewModel
protocol ManagementViewModelProtocol: ObservableObject {
    /// Items must be Hashable to use \self as identifier in ForEach
    associatedtype Item: Hashable
    var items: [Item] { get }
    func fetchItems()
    func delete(item: Item)
}

/// A generic SwiftUI view for managing items
struct AdminManagementView<VM: ManagementViewModelProtocol, RowContent: View, CreateContent: View>: View {
    @StateObject private var viewModel: VM
    private let title: String
    private let rowContent: (VM.Item) -> RowContent
    private let createContent: (VM) -> CreateContent
    @State private var showCreate = false

    init(
        viewModel: @autoclosure @escaping () -> VM,
        title: String,
        @ViewBuilder rowContent: @escaping (VM.Item) -> RowContent,
        @ViewBuilder createContent: @escaping (VM) -> CreateContent
    ) {
        self._viewModel = StateObject(wrappedValue: viewModel())
        self.title = title
        self.rowContent = rowContent
        self.createContent = createContent
    }

    var body: some View {
        List {
            ForEach(viewModel.items, id: \.self) { item in
                rowContent(item)
                    .swipeActions { deleteAction(item) }
            }
        }
        .navigationTitle(title)
        .toolbar { addButton }
        .sheet(isPresented: $showCreate) { createContent(viewModel) }
        .onAppear { viewModel.fetchItems() }
    }

    private func deleteAction(_ item: VM.Item) -> some View {
        Button(role: .destructive) { viewModel.delete(item: item) } label: {
            Label("Delete", systemImage: "trash")
        }
    }

    private var addButton: some ToolbarContent {
        ToolbarItem(placement: .primaryAction) {
            Button { showCreate = true } label: { Image(systemName: "plus") }
        }
    }
}
