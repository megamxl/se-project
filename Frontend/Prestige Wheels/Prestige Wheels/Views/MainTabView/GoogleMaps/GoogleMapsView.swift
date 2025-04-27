//
//  GoogleMapsView.swift
//  Prestige Wheels
//
//  Created by Heinz Schweitzer on 26.04.25.
//

import SwiftUI
import GoogleMaps

struct GoogleMapsView: UIViewRepresentable {
    func makeUIView(context: Context) -> GMSMapView {
        let camera = GMSCameraPosition.camera(withLatitude: 48.157975, longitude: 16.381778, zoom: 15)
        let mapView = GMSMapView(frame: .zero, camera: camera)
        
        // Marker hinzufügen
        let marker = GMSMarker()
        marker.position = CLLocationCoordinate2D(latitude: 48.157975, longitude: 16.381778)
        marker.title = "Pick up"
        marker.snippet = "Wien"
        marker.map = mapView
        
        return mapView
    }
    
    func updateUIView(_ uiView: GMSMapView, context: Context) {
        
    }
}
