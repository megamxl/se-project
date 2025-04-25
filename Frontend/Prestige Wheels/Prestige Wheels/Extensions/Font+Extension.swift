//
//  Font+Extension.swift
//  Prestige Wheels
//
//  Created by Michael Luegmayer on 08.04.25.
//

import SwiftUI

import SwiftUI

extension View {
    func faDuotoneLight(size: CGFloat, color: Color? = nil) -> some View {
        applyFont(named: "FontAwesome6Duotone-Light", size: size, color: color)
    }
    func faDuotoneRegular(size: CGFloat, color: Color? = nil) -> some View {
        applyFont(named: "FontAwesome6Duotone-Regular", size: size, color: color)
    }
    func faDuotoneSolid(size: CGFloat, color: Color? = nil) -> some View {
        applyFont(named: "FontAwesome6Duotone-Solid", size: size, color: color)
    }
    func faDuotoneThin(size: CGFloat, color: Color? = nil) -> some View {
        applyFont(named: "FontAwesome6Duotone-Thin", size: size, color: color)
    }

    // shared helper
    private func applyFont(named name: String, size: CGFloat, color: Color?) -> some View {
        // @ViewBuilder lets us conditionally apply modifiers
        @ViewBuilder
        var built: some View {
            if let c = color {
                self
                    .font(.custom(name, size: size))
                    .foregroundColor(c)
            } else {
                self
                    .font(.custom(name, size: size))
            }
        }
        return built
    }
}
