CREATE TABLE retry_logs (

    id SERIAL PRIMARY KEY,

    notification_id VARCHAR(100),

    error_message TEXT,

    retry_count INT,

    created_at TIMESTAMP DEFAULT NOW()
);