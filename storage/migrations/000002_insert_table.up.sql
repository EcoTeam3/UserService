-- Insert statements for users table
INSERT INTO users (username, email, password_hash)
VALUES
    ('john_doe', 'john.doe@example.com', 'hashed_password_1'),
    ('jane_smith', 'jane.smith@example.com', 'hashed_password_2'),
    ('alice_jones', 'alice.jones@example.com', 'hashed_password_3'),
    ('bob_brown', 'bob.brown@example.com', 'hashed_password_4'),
    ('charlie_davis', 'charlie.davis@example.com', 'hashed_password_5');

-- Insert statements for user_profiles table
INSERT INTO user_profiles (user_id, full_name, bio, location, avatar_url)
VALUES
    ((SELECT id FROM users WHERE username = 'john_doe'), 'John Doe', 'Software developer and tech enthusiast.', 'San Francisco, CA', 'https://example.com/avatar/john_doe.jpg'),
    ((SELECT id FROM users WHERE username = 'jane_smith'), 'Jane Smith', 'Digital marketer with a love for analytics.', 'New York, NY', 'https://example.com/avatar/jane_smith.jpg'),
    ((SELECT id FROM users WHERE username = 'alice_jones'), 'Alice Jones', 'Graphic designer and illustrator.', 'Los Angeles, CA', 'https://example.com/avatar/alice_jones.jpg'),
    ((SELECT id FROM users WHERE username = 'bob_brown'), 'Bob Brown', 'Content writer and SEO expert.', 'Chicago, IL', 'https://example.com/avatar/bob_brown.jpg'),
    ((SELECT id FROM users WHERE username = 'charlie_davis'), 'Charlie Davis', 'Project manager with a knack for efficiency.', 'Austin, TX', 'https://example.com/avatar/charlie_davis.jpg');
