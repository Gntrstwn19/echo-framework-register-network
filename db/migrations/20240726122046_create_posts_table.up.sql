CREATE TABLE customer (
                          customer_id int NOT NULL PRIMARY KEY,
                          name VARCHAR(255) NOT NULL,
                          address VARCHAR(255) NOT NULL,
                          username VARCHAR(255) NOT NULL,
                          password VARCHAR(255) NOT NULL,
                          phone_number VARCHAR(255) NOT NULL,
                          email VARCHAR(255) NOT NULL
);


CREATE TABLE payment (
                         payment_id SERIAL NOT NULL PRIMARY KEY,
                         customer_id INT NOT NULL,
                         payment_date DATE NOT NULL,
                         payment_amount DECIMAL(8, 2) NOT NULL,
                         payment_status VARCHAR(50) NOT NULL,
                         payment_method VARCHAR(255) NOT NULL,
                         CONSTRAINT fk_payment_customer FOREIGN KEY (customer_id) REFERENCES customer(customer_id)
);
