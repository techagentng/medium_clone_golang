CREATE TABLE IF NOT EXISTS posts (
    id  VARCHAR (50) NOT NULL PRIMARY KEY ,
    title VARCHAR (50) NOT NULL ,
    details TEXT NOT NULL ,
    access INT DEFAULT 0,
    views INT DEFAULT 0,
    user_id VARCHAR (50) NOT NULL ,
    created_at INT,
    updated_at INT,
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
    );