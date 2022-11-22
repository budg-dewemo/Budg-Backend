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
                      avatar VARCHAR(255) NULL,
                      PRIMARY KEY (id)
);

ALTER TABLE User ADD Column avatar VARCHAR(255) NULL;

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
VALUES ('admin', 'admin', 'admin@admin', 'admin@admin.com', '$2a$09$0BxHCT2cE/V3JurhuJQKM.vN4FrFKExYldvmvBWLpJGSGTULJO2iS', '2022-01-01 00:00:00', 1);
INSERT INTO User (name, last_name,username, email, password, last_login, active)
VALUES ('John', 'Smith', 'johnsmith', 'johnsmith@gmail.com', '$2a$09$0BxHCT2cE/V3JurhuJQKM.vN4FrFKExYldvmvBWLpJGSGTULJO2iS', '2022-01-01 00:00:00', 1);

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




-- Obtener todos los gastos asociados al presupuesto activo del usuario 2
SELECT * FROM Expense WHERE budget_id = (select id from Budget where current_budget = 1 and user_id=2);

-- Obtener todas las categorías asociadas al usuario 1 y 2 (El usuario 1 es el administrador)
SELECT * FROM Category WHERE user_id = 1 OR user_id = 2;

-- Obtener el monto total de gastos asociados al presupuesto activo del usuario 2
SELECT SUM(amount) FROM Expense WHERE budget_id = (select id from Budget where current_budget = 1 and user_id=2);
-- get rest of budget for user 2
SELECT (amount - (SELECT SUM(amount) FROM Expense WHERE budget_id = (select id from Budget where current_budget = 1 and user_id=2))) FROM Budget WHERE current_budget = 1 and user_id=2;

-- Obtner el monto total de gastos asociados al presupuesto activo del usuario 2
SELECT SUM(amount) FROM Expense WHERE budget_id = (select id from Budget where current_budget = 1 and user_id=2);

-- Dashboards


-- obtener los 5 gastos más altos del usuario 2
SELECT * FROM Expense WHERE budget_id = (select id from Budget where current_budget = 1 and user_id=2) ORDER BY amount DESC LIMIT 5;

-- Obtener el monto total de gastos asociados al presupuesto activo del usuario 2 agrupados por categoría
SELECT category_id, SUM(amount) FROM Expense WHERE budget_id = (select id from Budget where current_budget = 1 and user_id=2) GROUP BY category_id;

-- get amount and top 5 categories for user 2 in active budget
-- obtener el monto total de gastos asociados al presupuesto activo del usuario 2 agrupados por categoría
SELECT category_id, SUM(amount)
    FROM Expense
    WHERE budget_id = (select id from Budget where current_budget = 1 and user_id=2)
    GROUP BY category_id
    ORDER BY SUM(amount) DESC LIMIT 5;

-- get amount name and top 5 categories for user 2 in active budget
SELECT c.name, SUM(e.amount)
    FROM Expense e INNER JOIN Category c
        ON e.category_id = c.id WHERE
                                    e.budget_id = (select id from Budget where current_budget = 1 and user_id=2)
                                GROUP BY e.category_id ORDER BY SUM(e.amount) DESC LIMIT 5;

-- get amount for last 5 budgets for user 2
SELECT b.id, SUM(e.amount)
    FROM Expense e
        INNER JOIN Budget b ON e.budget_id = b.id
    WHERE b.user_id = 2 GROUP BY b.id ORDER BY b.id DESC LIMIT 5;


-- get sum of transactions type expense gruped by monthname for user 2
SELECT MONTHNAME(t.date), SUM(t.amount)
FROM User_transaction t
         INNER JOIN Budget b ON t.budget_id = b.id
WHERE b.user_id = 2 and t.type='expense' GROUP BY MONTHNAME(t.date) ORDER BY MONTH(t.date) DESC LIMIT 12;



-- get the last week expenses for user 2 grouped by category
SELECT c.name, SUM(e.amount)
    FROM User_transaction e
        INNER JOIN Category c ON e.category_id = c.id
    WHERE e.date BETWEEN DATE_SUB(NOW(), INTERVAL 7 DAY) AND NOW()
      AND e.type='expense' AND e.user_id=2
    GROUP BY c.name;

SELECT c.name, SUM(e.amount) FROM User_transaction e INNER JOIN Category c ON e.category_id = c.id WHERE e.date BETWEEN DATE_SUB(NOW(), INTERVAL 7 DAY) AND NOW() AND e.type='expense' AND e.user_id=2 GROUP BY c.name order by SUM(e.amount);