print('🔧 [Mongo Init] Switching to user_service DB...')
db = db.getSiblingDB('user_service')

print("🧨 Dropping existing 'users' collection (if any)...")
db.users.drop()

print("📐 Creating 'users' collection with schema validation...")
db.createCollection('users', {
  validator: {
    $jsonSchema: {
      bsonType: 'object',
      required: ['_id', 'email', 'name', 'password'],
      properties: {
        _id: { bsonType: 'string' },
        email: { bsonType: 'string' },
        name: { bsonType: 'string' },
        password: { bsonType: 'string' },
        admin: { bsonType: 'bool' }
      }
    }
  },
  validationLevel: 'strict',
  validationAction: 'error'
})

print('🔐 Creating unique index on email...')
db.users.createIndex({ email: 1 }, { unique: true })

print("✅ Finished setting up 'users' collection.")
