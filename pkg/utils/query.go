package utils

const (
	Create_User_Table = `
			CREATE TABLE IF NOT EXISTS user (
			id INT PRIMARY KEY AUTO_INCREMENT , 
			username VARCHAR(255) UNIQUE,
			email VARCHAR(255) UNIQUE, 
			password VARCHAR(255),
			isAdmin BOOLEAN DEFAULT FALSE,
			adminRequest BOOLEAN DEFAULT FALSE
			);`
	Create_Book_Table = `
	    	CREATE TABLE IF NOT EXISTS book (
			id INT PRIMARY KEY AUTO_INCREMENT,
			title VARCHAR(255) UNIQUE,
			author VARCHAR(255),
			publication_date DATE,
			quantity INT,
			genre VARCHAR(255),
			description LONGTEXT,
			rating FLOAT,
			address VARCHAR(255)
			);`
	Create_Issue_Table = `
			CREATE TABLE IF NOT EXISTS issue( 
			id INT PRIMARY KEY AUTO_INCREMENT,
			user_id INT NOT NULL, 
			book_id INT NOT NULL,
			issue_date DATE,
			expected_return_date DATE,
			returned_date DATE,
			isReturned BOOLEAN DEFAULT FALSE,
			returnRequested BOOLEAN DEFAULT FALSE,
			issueRequested BOOLEAN DEFAULT FALSE,
			fine FLOAT DEFAULT 0, 
			FOREIGN KEY (user_id) REFERENCES user(id), 
			FOREIGN KEY (book_id) REFERENCES book(id)
			);`
			
)
