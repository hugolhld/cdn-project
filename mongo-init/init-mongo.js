// init-mongo.js
print("🚀 Initialisation de la base de données...");

db = db.getSiblingDB('users_db'); // Création de la DB

db.createCollection('users'); // Création d'une collection "users"
db.createCollection('files')

db.users.insertMany([
    { name: "Admin", email: "admin@example.com", role: "admin" },
    { name: "User", email: "user@example.com", role: "user" }
]);

print("✅ Base de données et collection 'users' initialisées !");