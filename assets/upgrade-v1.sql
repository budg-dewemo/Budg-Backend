-- actualizacion, se renombra la tabla expenses a user_transactions
-- y se agrega la columna type, que puede ser 'expense' o 'income'

-- se renombra la tabla expenses a user_transactions

ALTER TABLE Expense RENAME TO User_transaction;


-- se agrega la columna type
ALTER TABLE User_transaction ADD COLUMN type VARCHAR(10) NOT NULL DEFAULT 'expense';
ALTER TABLE User_transaction ADD CONSTRAINT user_transactions_type_check CHECK (type IN ('expense', 'income'));