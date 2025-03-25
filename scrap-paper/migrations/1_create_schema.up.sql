CREATE TABLE users (
	id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
	email text NOT NULL UNIQUE,
	password text NOT NULL,
	token text NOT NULL,
	created_at timestamp with time zone DEFAULT now(),
	updated_at timestamp with time zone DEFAULT now()
);
CREATE TABLE scrap_papers (
	id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
	content text NOT NULL,
	is_private boolean NOT NULL DEFAULT false,
	user_id uuid NOT NULL REFERENCES users(id),
	created_at timestamp with time zone DEFAULT now(),
	updated_at timestamp with time zone DEFAULT now()
);