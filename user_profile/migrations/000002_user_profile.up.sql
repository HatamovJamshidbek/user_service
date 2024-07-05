
CREATE TYPE user_role AS ENUM ('musician', 'listener', 'producer');

CREATE TABLE user_profiles (
                              id uuid default  gen_random_uuid(),
                               user_id uuid   references users(id),
                               full_name VARCHAR(100),
                               bio TEXT,
                               role user_role,
                               location VARCHAR(100),
                               avatar_url VARCHAR(255),
                               website VARCHAR(255),
                                created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
                               updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
                               deleted_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
