CREATE TABLE notification_history (

    id SERIAL PRIMARY KEY,

    notification_id VARCHAR(100),

    channel VARCHAR(50),

    recipient VARCHAR(255),

    status VARCHAR(50),

    created_at TIMESTAMP DEFAULT NOW()
);