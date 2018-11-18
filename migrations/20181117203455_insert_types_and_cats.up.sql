INSERT INTO types (name, created_at, updated_at) 
VALUES 
    ('expenses', current_timestamp(), current_timestamp()), 
    ('bills', current_timestamp(), current_timestamp()),
    ('splurge', current_timestamp(), current_timestamp()),
    ('investments', current_timestamp(), current_timestamp());

INSERT INTO categories (name, created_at, updated_at) 
VALUES 
    ('food', current_timestamp(), current_timestamp()), 
    ('transportation', current_timestamp(), current_timestamp()),
    ('utilities', current_timestamp(), current_timestamp()),
    ('rent', current_timestamp(), current_timestamp()),
    ('entertainment', current_timestamp(), current_timestamp()),
    ('savings', current_timestamp(), current_timestamp()),
    ('insurance', current_timestamp(), current_timestamp()),
    ('charity', current_timestamp(), current_timestamp());

