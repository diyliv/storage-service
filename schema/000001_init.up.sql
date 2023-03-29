CREATE TABLE IF NOT EXISTS metrics(
    aasetId UUID, 
    value float, 
    quality INTEGER NOT NULL, 
    tstamp TIMESTAMP NOT NULL 
);