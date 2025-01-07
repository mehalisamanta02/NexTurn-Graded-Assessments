class Customer:
    def __init__(self, name, email, phone):
        self.name = name
        self.email = email
        self.phone = phone

    def display_details(self):
        return f"Name: {self.name}, Email: {self.email}, Phone: {self.phone}"

customers = []

def add_customer(name, email, phone):
    if not email or not phone:
        print("Error: Email and phone number are required.")
        return
    customer = Customer(name, email, phone)
    customers.append(customer)
    print("Customer added successfully!")

def view_customers():
    if customers:
        for customer in customers:
            print(customer.display_details())
    else:
        print("No customers found.")
