print('🔧 [Mongo Init] Switching to car_service DB...')
db = db.getSiblingDB('car_service')

print("🧨 Dropping existing 'cars' collection (if any)...")
db.cars.drop()

print("📐 Creating 'cars' collection with schema validation...")
db.createCollection('cars', {
  validator: {
    $jsonSchema: {
      bsonType: 'object',
      required: ['_id', 'model', 'brand', 'price_per_day'],
      properties: {
        _id: {
          bsonType: 'string',
          description: 'VIN (must be unique)'
        },
        model: {
          bsonType: 'string',
          description: 'Model is required and must be a string'
        },
        brand: {
          bsonType: 'string',
          description: 'Brand is required and must be a string'
        },
        price_per_day: {
          bsonType: 'double',
          minimum: 0,
          description: 'Price per day must be >= 0'
        },
        image_url: {
          bsonType: 'string',
          description: 'Optional image URL'
        }
      }
    }
  },
  validationLevel: 'strict',
  validationAction: 'error'
})

print("✅ Finished setting up 'cars' collection.")
