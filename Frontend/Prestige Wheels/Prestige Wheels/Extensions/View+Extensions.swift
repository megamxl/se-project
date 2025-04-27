//
//  View+Extensions.swift
//  Prestige Wheels
//
//  Created by Michael Luegmayer on 25.03.25.
//

import SwiftUI

extension View {
    /// Horizontally aligns the view by expanding its width to fill the available space
    /// and positioning the content according to the given alignment.
    ///
    /// - Parameter alignment: The horizontal alignment to use (e.g. `.leading`, `.center`, or `.trailing`).
    /// - Returns: A view that is stretched horizontally with the specified alignment.
    ///
    /// ### Usage:
    /// ```swift
    /// Text("Hello, SwiftUI!")
    ///     .hAlign(.center)
    /// ```
    func hAlign(_ alignment: Alignment) -> some View {
        frame(maxWidth: .infinity, alignment: alignment)
    }
    
    /// Vertically aligns the view by expanding its height to fill the available space
    /// and positioning the content according to the given alignment.
    ///
    /// - Parameter alignment: The vertical alignment to use (e.g. `.top`, `.center`, or `.bottom`).
    /// - Returns: A view that is stretched vertically with the specified alignment.
    ///
    /// ### Usage:
    /// ```swift
    /// Text("Hello, SwiftUI!")
    ///     .vAlign(.center)
    /// ```
    func vAlign(_ alignment: Alignment) -> some View {
        frame(maxHeight: .infinity, alignment: alignment)
    }
    
    // MARK: - Conditional Modifier
    
    /// Conditionally applies a transformation to the view based on the provided Boolean condition.
    ///
    /// Use this extension to avoid writing explicit if–else blocks in your view body.
    ///
    /// - Parameters:
    ///   - condition: A Boolean value determining whether to apply the transformation.
    ///   - transform: A closure that takes the view as input and returns a modified view.
    /// - Returns: Either the transformed view (if `condition` is true) or the original view.
    ///
    /// ### Usage:
    /// ```swift
    /// Text("Strike Through Example")
    ///     .if(true) { view in
    ///         view.strikethrough()
    ///     }
    /// ```
    @ViewBuilder
    func `if`<Content: View>(_ condition: @autoclosure () -> Bool,
                             transform: (Self) -> Content) -> some View
    {
        if condition() {
            transform(self)
        } else {
            self
        }
    }
}
