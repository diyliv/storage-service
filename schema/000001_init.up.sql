CREATE TABLE IF NOT EXISTS metrics(
    assetId UUID, 
    value float, 
    quality INTEGER NOT NULL, 
    tstamp TIMESTAMP NOT NULL 
);