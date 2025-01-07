from book_management import add_book, view_books, search_book
from customer_management import add_customer, view_customers
from sales_management import sell_book, view_sales

def main_menu():
    while True:
        print("\nWelcome to BookMart!")
        print("1. Book Management")
        print("2. Customer Management")
        print("3. Sales Management")
        print("4. Exit")
        choice = input("Enter your choice: ").strip()

        if choice == '1':
            book_menu()
        elif choice == '2':
            customer_menu()
        elif choice == '3':
            sales_menu()
        elif choice == '4':
            print("Thank you for using BookMart!")
            break
        else:
            print("Invalid choice. Please try again.")

def book_menu():
    while True:
        print("\nBook Management")
        print("1. Add Book")
        print("2. View Books")
        print("3. Search Book")
        print("4. Back to Main Menu")
        choice = input("Enter your choice: ").strip()

        if choice == '1':
            title = input("Enter title: ")
            author = input("Enter author: ")
            try:
                price = float(input("Enter price: "))
                quantity = int(input("Enter quantity: "))
                add_book(title, author, price, quantity)
            except ValueError:
                print("Invalid input! Price must be a number and quantity must be an integer.")
        elif choice == '2':
            view_books()
        elif choice == '3':
            keyword = input("Enter title or author to search: ")
            search_book(keyword)
        elif choice == '4':
            break
        else:
            print("Invalid choice. Please try again.")

def customer_menu():
    while True:
        print("\nCustomer Management")
        print("1. Add Customer")
        print("2. View Customers")
        print("3. Back to Main Menu")
        choice = input("Enter your choice: ").strip()

        if choice == '1':
            name = input("Enter name: ")
            email = input("Enter email: ")
            phone = input("Enter phone number: ")
            add_customer(name, email, phone)
        elif choice == '2':
            view_customers()
        elif choice == '3':
            break
        else:
            print("Invalid choice. Please try again.")

def sales_menu():
    while True:
        print("\nSales Management")
        print("1. Sell Book")
        print("2. View Sales Records")
        print("3. Back to Main Menu")
        choice = input("Enter your choice: ").strip()

        if choice == '1':
            customer_name = input("Enter customer name: ")
            book_title = input("Enter book title: ")
            try:
                quantity = int(input("Enter quantity: "))
                sell_book(customer_name, book_title, quantity)
            except ValueError:
                print("Invalid input! Quantity must be an integer.")
        elif choice == '2':
            view_sales()
        elif choice == '3':
            break
        else:
            print("Invalid choice. Please try again.")

if __name__ == "__main__":
    main_menu()
