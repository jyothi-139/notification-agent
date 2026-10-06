CREATE TABLE delivery_status (

    id SERIAL PRIMARY KEY,

    notification_id VARCHAR(100),

    status VARCHAR(50),

    remarks TEXT,

    updated_at TIMESTAMP DEFAULT NOW()
);