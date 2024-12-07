package database

var schema = `
    CREATE TABLE IF NOT EXISTS url (
        id SERIAL PRIMARY KEY,
        
        short VARCHAR(10) UNIQUE NOT NULL,
        long VARCHAR(255) NOT NULL,
        views INT DEFAULT 0,

        createdAt TIMESTAMP DEFAULT NOW(),  
        updaatedAt TIMESTAMP DEFAULT NOW()
    );

    CREATE INDEX IF NOT EXISTS index_short on url(short);
`
