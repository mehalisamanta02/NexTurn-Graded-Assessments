from book_management import books

class Transaction:
    def __init__(self, customer_name, book_title, quantity_sold):
        self.customer_name = customer_name
        self.book_title = book_title
        self.quantity_sold = quantity_sold

    def display_details(self):
        return f"Customer: {self.customer_name}, Book: {self.book_title}, Quantity: {self.quantity_sold}"

sales = []

def sell_book(customer_name, book_title, quantity):
    try:
        book = next((book for book in books if book.title.lower() == book_title.lower()), None)
        if not book:
            raise ValueError("Book not found.")
        if book.quantity < quantity:
            raise ValueError(f"Only {book.quantity} copies available. Sale cannot be completed.")
        book.quantity -= quantity
        sale = Transaction(customer_name, book_title, quantity)
        sales.append(sale)
        print("Sale successful!")
    except ValueError as e:
        print(f"Error: {e}")

def view_sales():
    if sales:
        for sale in sales:
            print(sale.display_details())
    else:
        print("No sales records found.")
