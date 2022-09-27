db = db.getSiblingDB(process.env.DB_NAME)

db.createCollection(process.env.HATS_COLLECTION_NAME);

for (var i = 0; i < process.env.TOTAL_HATS; i++) {
  db[process.env.HATS_COLLECTION_NAME].insertOne({ lastUsage: null, usedInPartyId: '' });
}

db.createCollection(process.env.LOCK_FLAG_COLLECTION_NAME);
db.lockFlag.insertOne({ isLocked: false });
 