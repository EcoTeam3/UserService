-- Delete statements for user_profiles table
DELETE FROM user_profiles WHERE user_id IN (
    SELECT id FROM users WHERE username IN ('john_doe', 'jane_smith', 'alice_jones', 'bob_brown', 'charlie_davis')
);

-- Delete statements for users table
DELETE FROM users WHERE username IN ('john_doe', 'jane_smith', 'alice_jones', 'bob_brown', 'charlie_davis');
