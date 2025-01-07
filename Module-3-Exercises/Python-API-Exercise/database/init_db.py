import os
from app import db, create_app

app = create_app()

with app.app_context():
    db.drop_all()  # Optional: Clear the database if running again
    db.create_all()

    # Add sample data
    sample_books = [
        {"title": "The Great Gatsby", "author": "F. Scott Fitzgerald", "published_year": 1925, "genre": "Fiction"},
        {"title": "1984", "author": "George Orwell", "published_year": 1949, "genre": "Sci-Fi"},
        {"title": "To Kill a Mockingbird", "author": "Harper Lee", "published_year": 1960, "genre": "Fiction"}
    ]

    for book in sample_books:
        new_book = db.Model(**book)
        db.session.add(new_book)

    db.session.commit()
    print("Database initialized with sample data!")
