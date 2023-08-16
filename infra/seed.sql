-- Create the users table
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) NOT NULL,
    name VARCHAR(100) NOT NULL,
    gender VARCHAR(10),
    dob DATE,
    email VARCHAR(100),
    phone VARCHAR(20)
);

-- Insert sample data
INSERT INTO users (username, name, gender, dob, email, phone) VALUES
    ('johdo', 'John Doe', 'Male', '1990-01-15', 'john.doe@example.com', '123-456-7890'),
    ('Jansm', 'Jane Smith', 'Female', '1985-03-22', 'jane.smith@example.com', '987-654-3210'),
    ('Micjo', 'Michael Johnson', 'Male', '1992-08-05', 'michael.johnson@example.com', '555-123-4567'),
    ('Emibr', 'Emily Brown', 'Female', '1988-12-10', 'emily.brown@example.com', '111-222-3333'),
    ('Davwi', 'David Wilson', 'Male', '1995-06-28', 'david.wilson@example.com', '444-555-6666'),
    ('Olida', 'Olivia Davis', 'Female', '1994-09-18', 'olivia.davis@example.com', '777-888-9999'),
    ('Wilma', 'William Martinez', 'Male', '1987-11-03', 'william.martinez@example.com', '222-333-4444'),
    ('Sopga', 'Sophia Garcia', 'Female', '1991-04-12', 'sophia.garcia@example.com', '666-777-8888'),
    ('Jamro', 'James Rodriguez', 'Male', '1993-07-20', 'james.rodriguez@example.com', '999-000-1111'),
    ('Isata', 'Isabella Taylor', 'Female', '1998-02-25', 'isabella.taylor@example.com', '123-456-7890');
