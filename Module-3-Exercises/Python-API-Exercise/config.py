import os

class Config:
	SECRET_KEY = os.environ.get('SECRET KEY')
	SQLALCHEMY_DATABASE_URI = r'C:\Users\Mehali\db\bookbuddy.db'
	SQLALCHEMY_TRACK_MODIFICATIONS = False