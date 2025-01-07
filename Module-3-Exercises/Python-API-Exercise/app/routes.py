from flask import Blueprint, jsonify, request
from .models import Book
from app import db

main_bp = Blueprint('main', __name__)

# POST /books - Add a new book
@main_bp.route('/books', methods=['POST'])
def add_book():
    data = request.get_json()

    if not all(key in data for key in ['title', 'author', 'published_year', 'genre']):
        return jsonify({"error": "Invalid data", "message": "All fields are required"}), 400

    try:
        book = Book(
            title=data['title'],
            author=data['author'],
            published_year=data['published_year'],
            genre=data['genre']
        )
        db.session.add(book)
        db.session.commit()
        return jsonify({"message": "Book added successfully", "book_id": book.id}), 201
    except Exception as e:
        return jsonify({"error": "Database error", "message": str(e)}), 500

# GET /books - Retrieve all books
@main_bp.route('/books', methods=['GET'])
def get_books():
    books = Book.query.all()
    return jsonify([
        {
            "id": book.id,
            "title": book.title,
            "author": book.author,
            "published_year": book.published_year,
            "genre": book.genre
        } for book in books
    ]), 200

# GET /books/<id> - Retrieve a specific book
@main_bp.route('/books/<int:book_id>', methods=['GET'])
def get_book(book_id):
    book = Book.query.get(book_id)
    if not book:
        return jsonify({"error": "Book not found", "message": "No book exists with the provided ID"}), 404

    return jsonify({
        "id": book.id,
        "title": book.title,
        "author": book.author,
        "published_year": book.published_year,
        "genre": book.genre
    }), 200

# PUT /books/<id> - Update a book
@main_bp.route('/books/<int:book_id>', methods=['PUT'])
def update_book(book_id):
    data = request.get_json()
    book = Book.query.get(book_id)

    if not book:
        return jsonify({"error": "Book not found", "message": "No book exists with the provided ID"}), 404

    try:
        book.title = data.get('title', book.title)
        book.author = data.get('author', book.author)
        book.published_year = data.get('published_year', book.published_year)
        book.genre = data.get('genre', book.genre)
        db.session.commit()
        return jsonify({"message": "Book updated successfully"}), 200
    except Exception as e:
        return jsonify({"error": "Database error", "message": str(e)}), 500

# DELETE /books/<id> - Delete a book
@main_bp.route('/books/<int:book_id>', methods=['DELETE'])
def delete_book(book_id):
    book = Book.query.get(book_id)
    if not book:
        return jsonify({"error": "Book not found", "message": "No book exists with the provided ID"}), 404

    try:
        db.session.delete(book)
        db.session.commit()
        return jsonify({"message": "Book deleted successfully"}), 200
    except Exception as e:
        return jsonify({"error": "Database error", "message": str(e)}), 500
