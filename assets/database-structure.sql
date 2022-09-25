-- Create Table User
CREATE TABLE User (
                      id INT NOT NULL AUTO_INCREMENT,
                      name VARCHAR(50) NOT NULL,
                      last_name VARCHAR(50) NOT NULL,
                      username VARCHAR(50) NOT NULL,
                      email VARCHAR(50) NOT NULL,
                      password VARCHAR(255) NOT NULL,
                      last_login DATETIME NULL,
                      active BOOLEAN NOT NULL,
                      PRIMARY KEY (id)
);

-- Create Table Budget
CREATE TABLE Budget (
                        id INT NOT NULL AUTO_INCREMENT,
                        name VARCHAR(50) NOT NULL,
                        user_id INT NOT NULL,
                        amount DECIMAL(10,2) NOT NULL,
                        start_date DATETIME NOT NULL,
                        end_date DATETIME NOT NULL,
                        current_budget Boolean NOT NULL,
                        PRIMARY KEY (id),
                        FOREIGN KEY (user_id) REFERENCES User(id)
);

-- Create Table Category
CREATE TABLE Category (
                          id INT NOT NULL AUTO_INCREMENT,
                          name VARCHAR(50) NOT NULL,
                          user_id INT NOT NULL,
                          PRIMARY KEY (id),
                          FOREIGN KEY (user_id) REFERENCES User(id)
);

-- Create Table Transaction
CREATE TABLE Expense (
                             id INT NOT NULL AUTO_INCREMENT,
                             user_id INT NOT NULL,
                             budget_id INT NOT NULL,
                             amount DECIMAL(10,2) NOT NULL,
                             description VARCHAR(255) NOT NULL,
                             category_id INT NOT NULL,
                             date DATETIME NOT NULL,
                             PRIMARY KEY (id),
                             FOREIGN KEY (user_id) REFERENCES User(id),
                             FOREIGN KEY (budget_id) REFERENCES Budget(id),
                             FOREIGN KEY (category_id) REFERENCES Category(id)
);
-- ----------------------------------------------------------------------------------------------------
-- EXAMPLE DATA
-- Create User
INSERT INTO User (name, last_name,username, email, password, last_login, active)
VALUES ('admin', 'admin', 'admin@admin', 'admin@admin.com', '123', '2022-01-01 00:00:00', 1);
INSERT INTO User (name, last_name,username, email, password, last_login, active)
VALUES ('John', 'Smith', 'johnsmith', 'johnsmith@gmail.com', '123', '2022-01-01 00:00:00', 1);

-- Create Budget
INSERT INTO Budget (user_id,name, amount, start_date, end_date, current_budget)
VALUES (2, 'Enero',1000, '2018-01-01 00:00:00', '2022-02-1 00:00:00', 1);
-- Get active budget
SELECT * FROM Budget WHERE current_budget = 1;
-- Create Category
INSERT INTO Category (name, user_id)
VALUES ('Food', 1);

Insert into Category (name, user_id)
VALUES ('drinks', 1);

INSERT INTO Category (name, user_id)
VALUES ('clothes', 1);

-- Create expenses
INSERT INTO Expense (user_id, budget_id, amount, description, category_id, date)
VALUES (2, 1, 105, 'Groceries', 1, '2022-01-01 00:00:00');
INSERT INTO Expense (user_id, budget_id, amount, description, category_id, date)
VALUES (2, 1, 1, 'Alfajor', 1, '2022-01-01 00:00:00');
INSERT INTO Expense (user_id, budget_id, amount, description, category_id, date)
VALUES (2, 1, 5, 'Cerveza', 1, '2022-01-01 00:00:00');

INSERT INTO Expense (user_id, budget_id, amount, description, category_id, date)
VALUES (2, 1, 50, 'Vodka', 2, '2022-01-01 00:00:00');
INSERT INTO Expense (user_id, budget_id, amount, description, category_id, date)
VALUES (2, 1, 100, 'Tshirt', 3, '2022-01-01 00:00:00');
insert into Expense (user_id, budget_id, amount, description, category_id, date)
values (2, 1, 150, 'Jeans', 3, '2022-01-01 00:00:00');



-- Example queries
-- Get all expenses for active budget for user 2
SELECT * FROM Expense WHERE budget_id = (select id from Budget where current_budget = 1 and user_id=2);

-- Get all categories for user 1 and user 2
SELECT * FROM Category WHERE user_id = 1 OR user_id = 2;

-- Get Ammount of expenses for active budget for user 2
SELECT SUM(amount) FROM Expense WHERE budget_id = (select id from Budget where current_budget = 1 and user_id=2);

-- get rest of budget for user 2
SELECT (amount - (SELECT SUM(amount) FROM Expense WHERE budget_id = (select id from Budget where current_budget = 1 and user_id=2))) FROM Budget WHERE current_budget = 1 and user_id=2;

-- Get sum of all expenses for user 2 in active budget
SELECT SUM(amount) FROM Expense WHERE budget_id = (select id from Budget where current_budget = 1 and user_id=2);

-- Dashboards

-- get top 5 expenses for user 2
SELECT * FROM Expense WHERE budget_id = (select id from Budget where current_budget = 1 and user_id=2) ORDER BY amount DESC LIMIT 5;

-- get total amount of expenses for user 2 by category in active budget
SELECT category_id, SUM(amount) FROM Expense WHERE budget_id = (select id from Budget where current_budget = 1 and user_id=2) GROUP BY category_id;

-- get amount and top 5 categories for user 2 in active budget
SELECT category_id, SUM(amount) FROM Expense WHERE budget_id = (select id from Budget where current_budget = 1 and user_id=2) GROUP BY category_id ORDER BY SUM(amount) DESC LIMIT 5;

-- get amount name and top 5 categories for user 2 in active budget
SELECT c.name, SUM(e.amount) FROM Expense e INNER JOIN Category c ON e.category_id = c.id WHERE e.budget_id = (select id from Budget where current_budget = 1 and user_id=2) GROUP BY e.category_id ORDER BY SUM(e.amount) DESC LIMIT 5;

-- get amount for last 5 budgets for user 2
SELECT b.id, SUM(e.amount) FROM Expense e INNER JOIN Budget b ON e.budget_id = b.id WHERE b.user_id = 2 GROUP BY b.id ORDER BY b.id DESC LIMIT 5;