-- Active: 1702542279972@@127.0.0.1@3306@library
-- Create Table User ADD
CREATE TABLE user (
    user_id INT AUTO_INCREMENT PRIMARY KEY NOT NULL,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    password VARCHAR(2000) NOT NULL,
    level ENUM('superadmin','admin','member') NOT NULL,
    is_enabled BOOLEAN NOT NULL,
    gender ENUM('male','female') NOT NULL,
    telp VARCHAR(15) NOT NULL,
    birthdate DATE NOT NULL,
    address VARCHAR(500) NOT NULL,
    foto VARCHAR(2000) NOT NULL,
    batas INT DEFAULT 3,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL
) engine=InnoDB;

CREATE TABLE category (
    category_id int PRIMARY KEY AUTO_INCREMENT,
    category VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) engine=InnoDB;

CREATE TABLE author (
    author_id int PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) engine=InnoDB;

CREATE TABLE publisher (
    publisher_id int PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) engine=InnoDB;

CREATE TABLE rak (
    rak_id int PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    rows_rak INT,
    col INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) engine=InnoDB;

CREATE TABLE book (
    book_id INT PRIMARY KEY AUTO_INCREMENT,
    title VARCHAR(255) NOT NULL,
    category_id INT NOT NULL,
    author_id INT NOT NULL,
    publisher_id INT NOT NULL,
    isbn VARCHAR(255),
    page_count INT NOT NULL,
    stock INT NOT NULL,
    publication_year INT NOT NULL,
    foto VARCHAR(2000),
    rak_id INT NOT NULL,
    price INT NOT NULL,
    admin_id int NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (category_id) REFERENCES category(category_id),
    FOREIGN KEY (author_id) REFERENCES author(author_id),
    FOREIGN KEY (publisher_id) REFERENCES publisher(publisher_id),
    FOREIGN KEY (rak_id) REFERENCES rak(rak_id),
    FOREIGN KEY (admin_id) REFERENCES user(user_id)
) engine=InnoDB;

CREATE TABLE book_loan (
    loan_id int PRIMARY KEY AUTO_INCREMENT,
    checkout_date TIMESTAMP NOT NULL,
    due_date TIMESTAMP NOT NULL,
    return_date TIMESTAMP DEFAULT NULL,
    status ENUM('onloan','returned','overdue') NOT NULL,
    book_id INT NOT NULL,
    user_id INT NOT NULL,
    admin_id INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (book_id) REFERENCES book(book_id),
    FOREIGN KEY (user_id) REFERENCES user(user_id),
    FOREIGN KEY (admin_id) REFERENCES user(user_id)
) engine=InnoDB;

CREATE TABLE penalties (
    penalty_id int PRIMARY KEY AUTO_INCREMENT,
    loan_id INT NOT NULL,
    penalty_amount INT NOT NULL,
    reason VARCHAR(1000),
    payment_status ENUM('paid','unpaid'),
    due_date TIMESTAMP NOT NULL,
    admin_id INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (loan_id) REFERENCES book_loan(loan_id),
    FOREIGN KEY (admin_id) REFERENCES user(user_id)
) engine=InnoDB;

ALTER Table book
ADD COLUMN deleted_at TIMESTAMP DEFAULT NULL

insert into user(name,
            email,
            password,
            level,
            is_enabled,
            gender,
            telp,
            birthdate,
            address,
            foto,
            batas) 
values('superadmin','superadmin@gmail.com','$2a$10$efAjNoRgH8Xa7hBB5B5y5Oi8SgMw.osO1iJUK6dLPZoKHf2ep.9FW',
        'superadmin',1,'male','08213254','2006-01-02','solo','https://res.cloudinary.com/dhtypvjsk/image/upload/v1703409186/uddbqkxn3xsytzrmo5ek.jpg',3)

insert into user(name,
            email,
            password,
            level,
            is_enabled,
            gender,
            telp,
            birthdate,
            address,
            foto,
            batas) 
values('admin','admin@gmail.com','$2a$10$efAjNoRgH8Xa7hBB5B5y5Oi8SgMw.osO1iJUK6dLPZoKHf2ep.9FW',
        'superadmin',1,'male','08213254','2006-01-02','solo','https://res.cloudinary.com/dhtypvjsk/image/upload/v1703409186/uddbqkxn3xsytzrmo5ek.jpg',3)

insert into user(name,
            email,
            password,
            level,
            is_enabled,
            gender,
            telp,
            birthdate,
            address,
            foto,
            batas) 
values('member','member@gmail.com','$2a$10$efAjNoRgH8Xa7hBB5B5y5Oi8SgMw.osO1iJUK6dLPZoKHf2ep.9FW',
        'superadmin',1,'male','08213254','2006-01-02','solo','https://res.cloudinary.com/dhtypvjsk/image/upload/v1703409186/uddbqkxn3xsytzrmo5ek.jpg',3)

