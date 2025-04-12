CREATE TABLE IF NOT EXISTS products (
                                        id SERIAL PRIMARY KEY,
                                        name TEXT NOT NULL,
                                        description TEXT,
                                        price NUMERIC(10, 2) NOT NULL
    );
