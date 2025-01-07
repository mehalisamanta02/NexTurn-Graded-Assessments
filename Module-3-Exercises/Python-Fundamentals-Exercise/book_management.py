class Book:
    def __init__(self, title, author, price, quantity):
        self.title = title
        self.author = author
        self.price = price
        self.quantity = quantity

    def display_details(self):
        return f"Title: {self.title}, Author: {self.author}, Price: {self.price}, Quantity: {self.quantity}"

books = []

def add_book(title, author, price, quantity):
    try:
        if price <= 0 or quantity <= 0:
            raise ValueError("Price and quantity must be positive numbers.")
        book = Book(title, author, price, quantity)
        books.append(book)
        print("Book added successfully!")
    except ValueError as e:
        print(f"Error: {e}")

def view_books():
    if books:
        for book in books:
            print(book.display_details())
    else:
        print("No books available.")

def search_book(keyword):
    results = [book for book in books if keyword.lower() in book.title.lower() or keyword.lower() in book.author.lower()]
    if results:
        for book in results:
            print(book.display_details())
    else:
        print("No matching books found.")
