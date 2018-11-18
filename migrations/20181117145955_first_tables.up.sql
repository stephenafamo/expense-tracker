-- Expenses/Splurge/Bills/Investments
CREATE TABLE types (
	id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	"name" STRING UNIQUE NOT NULL,
	created_at TIMESTAMP WITH TIME ZONE NULL DEFAULT current_timestamp(),
	updated_at TIMESTAMP WITH TIME ZONE NULL,
	FAMILY "primary" (id, "name", created_at, updated_at)
);

-- Food, Trnasportation, Rent, Utilities, Charity, Entertainment, Savings, Insurance
CREATE TABLE categories (
	id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	"name" STRING UNIQUE NOT NULL,
	created_at TIMESTAMP WITH TIME ZONE NULL DEFAULT current_timestamp(),
	updated_at TIMESTAMP WITH TIME ZONE NULL,
	FAMILY "primary" (id, "name", created_at, updated_at)
);

CREATE TABLE transactions (
	id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	type_id UUID NULL DEFAULT NULL REFERENCES types (id) ON DELETE SET NULL ON UPDATE CASCADE,
	category_id UUID NULL DEFAULT NULL REFERENCES categories (id) ON DELETE SET NULL ON UPDATE CASCADE,
	narration STRING NOT NULL,
	amount NUMERIC NOT NULL,
	currency STRING NOT NULL DEFAULT 'NGN',
	exchange_rate NUMERIC NOT NULL DEFAULT 1,
	amount_local NUMERIC AS (amount * exchange_rate) STORED,
	created_at TIMESTAMP WITH TIME ZONE NULL DEFAULT current_timestamp(),
	updated_at TIMESTAMP WITH TIME ZONE NULL,
	FAMILY "primary" (id, type_id, category_id, narration, amount, currency, exchange_rate, amount_local, created_at, updated_at)
);
